package server

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
)

type serverInfo struct {
	Hostname string
	Port     string
	NetAddrs []string
}

// PrintServerInfo prints server information
func (svr *Server) PrintServerInfo(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	var netAddrs []string
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		netAddrs = append(netAddrs, addr.String())
	}

	info := serverInfo{
		Hostname: hostname,
		Port:     svr.Port,
		NetAddrs: netAddrs,
	}
	payload, _ := json.MarshalIndent(info, "", "    ")
	fmt.Fprintln(w, string(payload))
}
