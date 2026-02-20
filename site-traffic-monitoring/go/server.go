package main

import (
	"fmt"
	"net/http"
	"sync"

	log "github.com/sirupsen/logrus"
)

// Type Server bundles a references to a mutex and a slice so as to
// serve data in a thread-safe fashion.
type Server struct {
	// The reference of a mutex shared with main to prevent race
	// conditions when accessing the currentReport member.
	lock *sync.Mutex

	// A reference to a slice containing the current report.
	currentReport *[]byte
}

// NewServer instantiates a new mutex and a slice with which to initialise a
// new Server.
func NewServer() Server {
	var (
		m             sync.Mutex
		currentReport []byte
	)

	return Server{&m, &currentReport}
}

// Serve will bootstrap the server based both on configured parameters made available on
// serverConf and tls. If the latter is false, the server will begin listening on HTTP, whilst
// if it is true, the server will instead support TLS at the transport layer. Note this function
// will not return: calls to http.ListenAndServe and http.ListenAndServeTLS run indefinitely.
// This implies Serve is supposed to be executed on its own goroutine.
func Serve(server Server, serverConf ServerOutputConf, tls bool) {
	if !tls {
		log.Debugf("begin listening on %s:%d over HTTP\n", serverConf.HTTP.BindAddress, serverConf.HTTP.BindPort)
		log.Fatal(http.ListenAndServe(
			fmt.Sprintf("%s:%d", serverConf.HTTP.BindAddress, serverConf.HTTP.BindPort), server))
	}

	log.Debugf("loading certificate from: %s\n", serverConf.HTTPS.CertPath)
	log.Debugf("loading key from: %s\n", serverConf.HTTPS.KeyPath)
	log.Debugf("begin listening on %s:%d over HTTPs\n", serverConf.HTTPS.BindAddress, serverConf.HTTPS.BindPort)
	log.Fatal(http.ListenAndServeTLS(
		fmt.Sprintf("%s:%d", serverConf.HTTPS.BindAddress, serverConf.HTTPS.BindPort),
		serverConf.HTTPS.CertPath, serverConf.HTTPS.KeyPath, server))
}

// Method ServeHTTP makes type Server implement the http.Handler interface.
// It defines how requests are served.
func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Requests will always be JSON documents
	w.Header().Set("Content-Type", "application/json")

	// Decide what to do based on the requested path.
	switch r.URL.Path {
	// By default we will serve the current report, accessed in mutual exclusion with main.
	case "/":
		w.WriteHeader(http.StatusOK)
		log.Debugf("serving report!\n")
		s.lock.Lock()
		fmt.Fprint(w, string(*s.currentReport))
		// Remember to release the lock when the function returns!
		defer s.lock.Unlock()
	// If anything else is requested, return a 404 with an error message.
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{"msg": "wrong path!"}`+"\n")
	}
}
