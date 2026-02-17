package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/bramvdbogaerde/go-scp/auth"
	"golang.org/x/crypto/ssh"

	scp "github.com/bramvdbogaerde/go-scp"
	snmp "github.com/gosnmp/gosnmp"
	log "github.com/sirupsen/logrus"
)

// This map makes it easier to parse the configured SNMP version into
// a version identifier 'understandable' by gosnmp.
var snmpVersionMap map[string]snmp.SnmpVersion = map[string]snmp.SnmpVersion{
	"1":  snmp.Version1,
	"2c": snmp.Version2c,
	"3":  snmp.Version3,
}

// This map makes it easier to parse the configured SNMPv3 authentication
// protocol into the internal representation leveraged by gosnmp.
var snmpAuthProtMap = map[string]snmp.SnmpV3AuthProtocol{
	// "NoAuth": snmp.NoAuth,
	"MD5":    snmp.MD5,
	"SHA":    snmp.SHA,
	"SHA224": snmp.SHA224,
	"SHA256": snmp.SHA256,
	"SHA384": snmp.SHA384,
	"SHA512": snmp.SHA512,
}

// This map makes it easier to parse the configured SNMPv3 privacy
// protocol into the internal representation leveraged by gosnmp.
var snmpPrivProtMap = map[string]snmp.SnmpV3PrivProtocol{
	// "NoPriv":  snmp.NoPriv,
	"DES":     snmp.DES,
	"AES":     snmp.AES,
	"AES192":  snmp.AES192,
	"AES256":  snmp.AES256,
	"AES192C": snmp.AES192C,
	"AES256C": snmp.AES256C,
}

// outputToFile truncates the provided path and then writes the second argument 'as-is'
// into said path. It returns any errors thrown by os.Create and/or os.Write.
func outputToFile(path string, report []byte) error {
	fd, err := os.Create(path)
	if err != nil {
		return err
	}
	defer fd.Close()

	_, err = fd.Write(report)
	return err
}

// generateReport embeds inRate and outRate into the appropriate JSON fields to create
// a report as specified by the definition of the outputReport type. Fields of conf are
// also embedded into the final report. The generate report is provided as a byte slice
// for convenience and any marshalling errors are returned.
func generateReport(conf Conf, inRate float64, outRate float64) ([]byte, error) {
	report, err := json.MarshalIndent(outputReport{
		Description:         conf.ReportDescription,
		UpdatedLast:         time.Now().UTC().Format(iso8601Format),
		InBytesPerSec:       inRate,
		OutBytesPerSec:      outRate,
		UpdateInterval:      fmt.Sprintf("%d seconds", conf.Interval),
		MonitoredInterfaces: conf.monitoredInterfaces,
	}, "", "\t")
	if err != nil {
		return nil, err
	}
	return report, nil
}

// readConf will attempt to read a JSON configuration file from path and then unmarshal
// it into a Conf-type structure. It will also process defined switches by calling
// processBorderSwitch on each of them. Any errors are returned to the caller.
func readConf(path string) (Conf, error) {
	rawConf, err := os.ReadFile(path)
	if err != nil {
		return Conf{}, err
	}

	var conf Conf

	if err := json.Unmarshal(rawConf, &conf); err != nil {
		return Conf{}, err
	}

	for i := range conf.BorderSwitches {
		monIfacesChunk, err := processBorderSwitch(&conf.BorderSwitches[i])
		if err != nil {
			return Conf{}, err
		}
		conf.monitoredInterfaces = append(conf.monitoredInterfaces, monIfacesChunk...)
	}

	return conf, nil
}

// processBorder switch will update a switch's (of type BorderSwitch) SNMP Version,
// monitored input OIDs, monitored output OIDs and monitored interface IDs. The OIDs
// are either the regular ifHC[In|Out]OctetsOID or their HC counterparts based on the
// switch's definition in the configuration file. Interface IDs are constructed by
// concatenating the switch's name and the interface description as specified on the
// configuration file. This ensures each ID will be associated with a unique string
// throughout the entire codebase. As the input switch is passed by reference, it will
// be modified in-place. The slice of monitored interface IDs will be returned, together
// with any errors caused by the process.
func processBorderSwitch(bSwitch *BorderSwitch) ([]string, error) {
	parsedVer, ok := snmpVersionMap[strings.ToLower(bSwitch.SNMPVersion)]
	if !ok {
		return nil, fmt.Errorf("couldn't parse configured SNMP Version: %s", bSwitch.SNMPVersion)
	}
	bSwitch.SNMPVersionParsed = parsedVer

	if parsedVer == snmp.Version3 {
		parsedAuthProt, ok := snmpAuthProtMap[strings.ToUpper(bSwitch.SNMPAuthProt)]
		if !ok {
			return nil, fmt.Errorf("couldn't parse configured SNMPv3 authentication protocol: %s", bSwitch.SNMPAuthProt)
		}
		bSwitch.SNMPAuthProtParsed = parsedAuthProt

		parsedPrivProt, ok := snmpPrivProtMap[strings.ToUpper(bSwitch.SNMPPrivProt)]
		if !ok {
			return nil, fmt.Errorf("couldn't parse configured SNMPv3 privacy protocol: %s", bSwitch.SNMPPrivProt)
		}
		bSwitch.SNMPPrivProtParsed = parsedPrivProt
	}

	monIfaces := []string{}
	for _, iFace := range bSwitch.Interfaces {
		if bSwitch.HCSupport {
			(*bSwitch).inMonitOIDs = append(bSwitch.inMonitOIDs, fmt.Sprintf(
				"%s.%d", ifHCInOctetsOID, iFace.SNMPIndex,
			))
		} else {
			(*bSwitch).inMonitOIDs = append(bSwitch.inMonitOIDs, fmt.Sprintf(
				"%s.%d", ifInOctetsOID, iFace.SNMPIndex,
			))
		}

		if bSwitch.HCSupport {
			bSwitch.outMonitOIDs = append(bSwitch.outMonitOIDs, fmt.Sprintf(
				"%s.%d", ifHCOutOctetsOID, iFace.SNMPIndex,
			))
		} else {
			bSwitch.outMonitOIDs = append(bSwitch.outMonitOIDs, fmt.Sprintf(
				"%s.%d", ifOutOctetsOID, iFace.SNMPIndex,
			))
		}

		monIfaces = append(monIfaces, fmt.Sprintf("%s_%s", bSwitch.HostName, iFace.Descr))
	}

	return monIfaces, nil
}

// getIndex from OID will split an OID by the dots and return the last element. Given how
// interface tables are traversed, this amounts to extracting the interface ID from an OID.
func getIndexFromOID(oid string) string {
	elms := strings.Split(oid, ".")
	return elms[len(elms)-1]
}

// gatherOctetCounters queries switches to find the input and output octet counts for
// each interface at a particular point in time. These counts are gathered by querying
// switches through SNMP Get requests to the appropriate OIDs. These counts are returned
// as a map where the keys are the interface's unique ID as generated on processBorderSwitch
// and the values are both the input and output octet counts together with the timestamps
// of when these counts where gathered. This allows the program to compute input and output
// octet rates based solely on this map: nothing else is needed. Any errors encountered
// during the process are returned back to the caller.
func gatherOctetCounters(BorderSwitches []BorderSwitch) (map[string]ifaceCount, error) {
	interfaceCounts := map[string]ifaceCount{}

	for _, borderSwitch := range BorderSwitches {
		var snmpCli *snmp.GoSNMP
		if borderSwitch.SNMPVersionParsed == snmp.Version3 {
			snmpCli = &snmp.GoSNMP{
				Target:        borderSwitch.HostName,
				Port:          161,
				Version:       borderSwitch.SNMPVersionParsed,
				Timeout:       time.Duration(5) * time.Second,
				SecurityModel: snmp.UserSecurityModel,
				MsgFlags:      snmp.AuthPriv,
				SecurityParameters: &snmp.UsmSecurityParameters{
					UserName:                 borderSwitch.SNMPSecName,
					AuthenticationProtocol:   borderSwitch.SNMPAuthProtParsed,
					AuthenticationPassphrase: borderSwitch.SNMPAuthPass,
					PrivacyProtocol:          borderSwitch.SNMPPrivProtParsed,
					PrivacyPassphrase:        borderSwitch.SNMPPrivPass,
				},
			}
			log.Debugf("crafted snmp client for switch %s: %s, %s, %s\n", snmpCli.Target, snmpCli.Version, snmpCli.MsgFlags, snmpCli.SecurityParameters.Description())
		} else {
			snmpCli = &snmp.GoSNMP{
				Target:    borderSwitch.HostName,
				Port:      161,
				Version:   borderSwitch.SNMPVersionParsed,
				Community: borderSwitch.SNMPCommunity,
				Timeout:   time.Duration(2) * time.Second,
			}
			log.Debugf("crafted snmp client for switch %s: %s, %s\n", snmpCli.Target, snmpCli.Version, snmpCli.Community)
		}

		if err := snmpCli.Connect(); err != nil {
			log.Warnf("couldn't connect to router %s: %v. Skipping it...\n", borderSwitch.HostName, err)
			continue
		}

		reply, err := snmpCli.Get(borderSwitch.inMonitOIDs)
		if err != nil {
			log.Warnf("error getting input statistics for switch %s: %v\n", borderSwitch.HostName, err)
			snmpCli.Conn.Close()
			continue
		}

		for _, snmpPDU := range reply.Variables {
			log.WithFields(log.Fields{
				"ifaceId":     fmt.Sprintf("%s-%s", borderSwitch.HostName, getIndexFromOID(snmpPDU.Name)),
				"octectCount": snmp.ToBigInt(snmpPDU.Value).Uint64(),
			}).Debug("input octet count")

			interfaceCounts[fmt.Sprintf("%s-%s", borderSwitch.HostName, getIndexFromOID(snmpPDU.Name))] = ifaceCount{
				inOctets: dataPoint{
					TimeStamp:  time.Now(),
					OctetCount: snmp.ToBigInt(snmpPDU.Value).Uint64(),
				},
			}
		}

		reply, err = snmpCli.Get(borderSwitch.outMonitOIDs)
		if err != nil {
			log.Warnf("error getting output statistics for switch %s: %v\n", borderSwitch.HostName, err)
			snmpCli.Conn.Close()
			continue
		}

		for _, snmpPDU := range reply.Variables {
			currIfaceCount := interfaceCounts[fmt.Sprintf("%s-%s", borderSwitch.HostName, getIndexFromOID(snmpPDU.Name))]

			log.WithFields(log.Fields{
				"ifaceId":     fmt.Sprintf("%s-%s", borderSwitch.HostName, getIndexFromOID(snmpPDU.Name)),
				"octectCount": snmp.ToBigInt(snmpPDU.Value).Uint64(),
			}).Debug("output octet count")

			currIfaceCount.outOctets = dataPoint{
				TimeStamp:  time.Now(),
				OctetCount: snmp.ToBigInt(snmpPDU.Value).Uint64(),
			}

			interfaceCounts[fmt.Sprintf("%s-%s", borderSwitch.HostName, getIndexFromOID(snmpPDU.Name))] = currIfaceCount
		}

		snmpCli.Conn.Close()
	}

	return interfaceCounts, nil
}

// copyOverScp leverages the SCP protocol to copy the report to an external machine. Bear in mind SSH key-based
// authentication must be set up before this function can be leveraged. The function returns any errors encountered
// along the process. It also implements host key checking so as to prevent MITM attacks. Note that the implementation
// of the ssh.HostKeyCallback type is largely based on the ideas set forth in [0].
// 0: https://stackoverflow.com/questions/44269142/golang-ssh-getting-must-specify-hoskeycallback-error-despite-setting-it-to-n
func copyOverScp(conf ScpOutputConf, report []byte) error {
	clientConfig, err := auth.PrivateKey(conf.User, conf.PrivateKeyPath, func(_ string, _ net.Addr, serverKey ssh.PublicKey) error {
		if conf.ServerPublicKey == "" {
			log.Warn("proceeding with no SSH key verification whatsoever")
			return nil
		}
		mServerKey := serverKey.Type() + " " + base64.StdEncoding.EncodeToString(serverKey.Marshal())
		if mServerKey != conf.ServerPublicKey {
			return fmt.Errorf("the server's SSH key (%q) is not the expected one: %q", mServerKey, conf.ServerPublicKey)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("couldn't configure authentication: %v", err)
	}

	client := scp.NewClient(fmt.Sprintf("%s:%d", conf.Hostname, conf.Port), &clientConfig)
	if err := client.Connect(); err != nil {
		return err
	}
	defer client.Close()

	reportReader := bytes.NewReader(report)

	if err := client.Copy(context.Background(), reportReader, conf.RemotePath,
		conf.Permissions, int64(reportReader.Len())); err != nil {
		return fmt.Errorf("error while copying the file: %v", err)
	}

	return nil
}
