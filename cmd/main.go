package main

import (
	"flag"
	"net/http"
	"strings"

	"github.com/M00nF1sh/echoserver/pkg/server"
)

func main() {
	argPorts := flag.String("ports", "8080", "comma separated list of ports to listen on")
	flag.Parse()

	ports := strings.Split(*argPorts, ",")
	for _, port := range ports {
		server := server.New(port)
		go http.ListenAndServe(":"+port, server)
	}

	finish := make(chan bool)
	<-finish
}
