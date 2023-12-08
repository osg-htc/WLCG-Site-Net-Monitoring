package main

import (
	"time"

	snmp "github.com/gosnmp/gosnmp"
)

const (
	// ifHCInOctetsOID is the OID of High Capacity (64 bit) interface input octet counts.
	// Check http://oid-info.com/cgi-bin/display?oid=1.3.6.1.2.1.31.1.1.1.6.
	ifHCInOctetsOID string = ".1.3.6.1.2.1.31.1.1.1.6"

	// ifHCOutOctetsOID is the OID of High Capacity (64 bit) interface output octet counts.
	// Check http://oid-info.com/cgi-bin/display?oid=1.3.6.1.2.1.31.1.1.1.10.
	ifHCOutOctetsOID string = ".1.3.6.1.2.1.31.1.1.1.10"

	// ifInOctetsOID is the OID of 32 bit interface input octet counts.
	// Check http://oid-info.com/cgi-bin/display?oid=1.3.6.1.2.1.2.2.1.10.
	ifInOctetsOID string = ".1.3.6.1.2.1.2.2.1.10"

	// ifOutOctetsOID is the OID of 32 bit interface output octet counts.
	// Check http://oid-info.com/cgi-bin/display?oid=1.3.6.1.2.1.2.2.1.16.
	ifOutOctetsOID string = ".1.3.6.1.2.1.2.2.1.16"

	// The template used to format time.Time values as strings. Check https://pkg.go.dev/time#Time.Format.
	iso8601Format string = "2006-01-02T15:04:05.000000-07:00"
)

// Type OutputConf defines the outputs to serve generated reports on.
type OutputConf struct {
	// File defines a regular file to dump reports onto.
	File FileOutputConf `json:"file"`

	// Server defines the builtin HTTP(S) server configuration.
	Server ServerOutputConf `json:"server"`

	// SCP-based remote copying configuration
	Scp ScpOutputConf `json:"scp"`
}

// Type FileOutput defines the regular file to dump reports onto.
type FileOutputConf struct {
	// Whether to generate the file or not.
	Enabled bool `json:"enabled"`

	// The path to write generated reports to. This will will be truncated.
	Path string `json:"path"`
}

// Type ServerOutputConf defines the configuration of the builtin HTTP(S) servers.
type ServerOutputConf struct {
	// The builtin HTTP server configuration.
	HTTP HTTPConf `json:"http"`

	// The builtin HTTPS server configuration.
	HTTPS HTTPSConf `json:"https"`
}

// Type ScpOutputConf defines the options controlling where to copy the report over SCP.
// Bear in mind in order for this output to work the machine this program is running on
// must be capable of running scp on a shell and have it work 'as-is'. That is, key-based
// authentication should already be set up.
type ScpOutputConf struct {
	// Whether to copy the report over SSH or nor.
	Enabled bool `json:"enabled"`

	// The user to establish the SSH session as.
	User string `json:"user"`

	// The hostname to establish the SSH session with.
	Hostname string `json:"hostname"`

	// The port to establish the SSH session on.
	Port int `json:"port"`

	// The path of the local private key grating access to the URI's user at the URI's hostname.
	PrivateKeyPath string `json:"PrivateKeyPath"`

	// The Base64-encoded server's public key for key verification.
	ServerPublicKey string `json:"ServerPublicKey"`

	// The remote path to copy the report into.
	RemotePath string `json:"remotePath"`

	// The permissions of the target file as if given to chmod(1).
	Permissions string `json:"permissions"`
}

// Type HTTPConf defines the configuration of the builtin HTTP Server.
type HTTPConf struct {
	// Whether to bootstrap the server or not.
	Enabled bool `json:"enabled"`

	// The address to bind the server to. If set to 0.0.0.0, the server will
	// listen on all available interfaces.
	BindAddress string `json:"bindAddress"`

	// The port to bind the server to.
	BindPort int `json:"bindPort"`
}

// Type HTTPSConf defines the configuration of the builtin HTTPS Server.
type HTTPSConf struct {
	// Whether to bootstrap the server or not.
	Enabled bool `json:"enabled"`

	// The address to bind the server to. If set to 0.0.0.0, the server will
	// listen on all available interfaces.
	BindAddress string `json:"bindAddress"`

	// The port to bind the server to.
	BindPort int `json:"bindPort"`

	// The path to the server's certificate.
	CertPath string `json:"certPath"`

	// The path to the server's private key.
	KeyPath string `json:"keyPath"`
}

// Type Conf defines all the configuration parameters.
type Conf struct {
	// The description to embed into the report 'as-is'.
	ReportDescription string `json:"description"`

	// The time (in seconds) to wait between data acquisitions.
	Interval uint32 `json:"interval"`

	// The configuration on where to expose the generated report.
	Outputs OutputConf `json:"outputs"`

	// The border switches to monitor.
	BorderSwitches []BorderSwitch `json:"borderSwitches"`

	// The interfaces being monitored to be embedded into the report.
	monitoredInterfaces []string
}

// Type BorderSwitch defines a switch whose interfaces we are to be monitored.
type BorderSwitch struct {
	// The switches hostname.
	HostName string `json:"hostname"`

	// Whether the switch's SNMP implementation offers High Capacity counters.
	HCSupport bool `json:"hcSupport"`

	// The switch's SNMP implementation's version.
	SNMPVersion string `json:"snmpVersion"`

	// The parsed SNMP implementation meant for internal program use.
	SNMPVersionParsed snmp.SnmpVersion

	// The SNMP community to include on SNMP Get requests.
	SNMPCommunity string `json:"snmpCommunity"`

	// The interfaces to be monitored.
	Interfaces []Interface `json:"interfaces"`

	// The OIDs to be monitored when gathering input data. These will be
	// populated by processBorderSwitch.
	inMonitOIDs []string

	// The OIDs to be monitored when gathering input data. These will be
	// populated by processBorderSwitch.
	outMonitOIDs []string
}

// Type Interface defines an interface to monitor.
type Interface struct {
	// The interface's description.
	Descr string `json:"descr"`

	// The associated SNMP Index as obtained by 'walking' over
	// IF-MIB::ifDescr (i.e. OID 1.3.6.1.2.1.2.2.1.2).
	SNMPIndex uint `json:"snmpIndex"`
}

// Type ifaceCount stores input and output octet counts at a point in time.
type ifaceCount struct {
	// The input octet count together with the acquisition timestamp.
	inOctets dataPoint

	// The output octet count together with the acquisition timestamp.
	outOctets dataPoint
}

// Type dataPoint bundles an octet count together with the time of acquisition.
type dataPoint struct {
	// The time of the acquisition.
	TimeStamp time.Time

	// The value of the count.
	OctetCount uint64
}

// Type outputReport defines the structure of generated reports
type outputReport struct {
	// The site description as defined on Conf.ReportDescription
	Description string `json:"Description"`

	// The ISO 8601-formatted time the report was generated.
	UpdatedLast string `json:"UpdatedLast"`

	// The global input rate, in bytes/s at UpdatedLast.
	InBytesPerSec float64 `json:"InBytesPerSec"`

	// The global output rate, in bytes/s at UpdatedLast.
	OutBytesPerSec float64 `json:"OutBytesPerSec"`

	// The interval at which this report is updated.
	UpdateInterval string `json:"UpdateInterval"`

	// The list monitored interfaces identified by their unique IDs.
	MonitoredInterfaces []string `json:"MonitoredInterfaces"`
}
