package main

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"flag"
)

func init() {
	flag.BoolVar(&version, "version", false, "Show the built commit hash and exit.")
	flag.BoolVar(&debug, "debug", false, "Whether to enable debugging output.")
	flag.StringVar(&confPath, "conf", "/etc/wlcg-site-snmp/conf.json", "Path to JSON configuration definition.")
}

var (
	// Global flag indicating whether debug messages should be logged.
	debug bool

	// Global flag containing the path of the configuration file.
	confPath string

	// Global flag indicating whether to display the built commit and exit.
	version bool

	// Git commit hash this program was built from.
	commit string
)

func main() {
	// Grab the command line arguments.
	flag.Parse()

	// If the caller requests the version just print it and quit.
	if version {
		fmt.Printf("Built commit hash: %s\n", commit)
		os.Exit(0)
	}

	// Enable verbosy output based on the input arguments.
	if debug {
		log.SetLevel(log.DebugLevel)
	}

	/*
	 * Try to load the configuration, quitting on error. Bear in mind this
	 * process includes some transformations on the 'raw' configuration to
	 * generate, for instance, unique interface IDs within the context of the
	 * program.
	 */
	log.Printf("trying to read configuration from %s\n", confPath)
	conf, err := readConf(confPath)
	if err != nil {
		log.Fatalf("couldn't load the configuration from %s: %v\n", confPath, err)
	}

	// Signal we're beginning polling
	log.Infof("kicking off at %s", time.Now().UTC().Format(iso8601Format))

	/*
	 * Bootstrap a builtin HTTP server if configured to do so. Note the server
	 * runs in its own goroutine (i.e. concurrently). This implies whe need
	 * to synchronise access to shared variables...
	 */
	var httpServer Server
	if conf.Outputs.Server.HTTP.Enabled {
		httpServer = NewServer()
		go Serve(httpServer, conf.Outputs.Server, false)
	}

	/*
	 * Bootstrap a builtin HTTPS server if configured to do so. Note the server
	 * runs in its own goroutine (i.e. concurrently). This implies whe need
	 * to synchronise access to shared variables...
	 */
	var httpsServer Server
	if conf.Outputs.Server.HTTPS.Enabled {
		httpsServer = NewServer()
		go Serve(httpsServer, conf.Outputs.Server, true)
	}

	/*
	 * These two local variables hold the input and output octet counts at two
	 * different points in time: they can be thought of as 'snapshots'. Input
	 * and output traffic rates are computed based on these counts.
	 */
	var (
		currentIfaceCounts  map[string]ifaceCount
		previousIfaceCounts map[string]ifaceCount
	)

	// Time to begin polling indefinitely...
	for {
		// Get the current I/O octet counts. If we fail to do so, restart the loop.
		currentIfaceCounts, err = gatherOctetCounters(conf.BorderSwitches)
		if err != nil {
			log.Warnf("couldn't get interface counts: %v\n", err)
			continue
		}

		/*
		 * If we had some previous values, begin computing rates. In the first iteration
		 * this won't be the case, so we'll just a single timeout to get enough data to
		 * begin crunching the numbers.
		 */
		if previousIfaceCounts != nil {
			// Initialise the rate accumulators.
			var (
				inRate  float64 = 0
				outRate float64 = 0
			)

			/*
			 * We'll compute data in the usual fashion: (currRate - prevRate) / interval. However,
			 * instead of resorting to the configured interval we'll use timestamps generated at
			 * the time of the octet count acquisition (i.e. at the time of sending the SNMP Gets).
			 * These timestamps are stored alongside the counts on the `currentIfaceCounts` and
			 * `previousIfaceCounts` maps so that it's all very convenient.
			 */
			for iFaceId, currentIfaceCounts := range currentIfaceCounts {
				inRate += float64(currentIfaceCounts.inOctets.OctetCount-previousIfaceCounts[iFaceId].inOctets.OctetCount) /
					currentIfaceCounts.inOctets.TimeStamp.Sub(previousIfaceCounts[iFaceId].inOctets.TimeStamp).Seconds()

				outRate += float64(currentIfaceCounts.outOctets.OctetCount-previousIfaceCounts[iFaceId].outOctets.OctetCount) /
					currentIfaceCounts.outOctets.TimeStamp.Sub(previousIfaceCounts[iFaceId].outOctets.TimeStamp).Seconds()

				log.WithFields(log.Fields{
					"inOctets":  inRate,
					"outOctets": outRate,
				}).Debug("current rate accumulators")
			}

			// Generate the output report (i.e. embed the previous accumulators into the output JSON).
			report, err := generateReport(conf, inRate, outRate)
			if err != nil {
				log.Errorf("couldn't generate the report: %v\n", err)
				continue
			}

			/*
			 * If configured to do so, dump the report to a file on disk. Note this
			 * file will be truncated!
			 */
			if conf.Outputs.File.Enabled {
				log.WithFields(log.Fields{"path": conf.Outputs.File.Path}).Debug("generating output report")
				if err := outputToFile(conf.Outputs.File.Path, report); err != nil {
					log.Errorf("couldn't store the report: %v\n", err)
				}
			}

			/*
			 * If the builtin HTTP server has been configured, update the report it serves.
			 * Be sure to do that leveraging mutual exclusion (i.e. a mutex) to avoid race
			 * conditions!
			 */
			if conf.Outputs.Server.HTTP.Enabled {
				httpReport := make([]byte, len(report))
				copy(httpReport, report)
				httpServer.lock.Lock()
				*httpServer.currentReport = httpReport
				httpServer.lock.Unlock()
			}

			/*
			 * If the builtin HTTPS server has been configured, update the report it serves.
			 * Be sure to do that leveraging mutual exclusion (i.e. a mutex) to avoid race
			 * conditions!
			 */
			if conf.Outputs.Server.HTTPS.Enabled {
				httpsReport := make([]byte, len(report))
				copy(httpsReport, report)
				httpServer.lock.Lock()
				*httpsServer.currentReport = httpsReport
				httpServer.lock.Unlock()
			}

			// If enabled, copy the report to a given destination machine.
			if conf.Outputs.Scp.Enabled {
				if err := copyOverScp(conf.Outputs.Scp, report); err != nil {
					log.Errorf("couldn't copy the report over SCP: %v\n", err)
				}
			}
		} else {
			log.Debug("skipping rate computation: no previous data found...")
		}

		// Just refresh the previous data to get ready for the next iteration...
		previousIfaceCounts = currentIfaceCounts

		// ... and wait the configured interval to repeat the process all over again.
		log.Debugf("waiting %d seconds till the next polling process...", conf.Interval)
		time.Sleep(time.Duration(conf.Interval) * time.Second)
	}
}
