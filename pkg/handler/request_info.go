package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type requestInfo struct {
	Host       string
	URL        string
	RemoteAddr string
	Headers    map[string][]string
}

// PrintRequestInfo prints request information
func (h *Handler) PrintRequestInfo(w http.ResponseWriter, r *http.Request) {
	info := requestInfo{
		Host:       r.Host,
		URL:        r.URL.Path,
		RemoteAddr: r.RemoteAddr,
		Headers:    r.Header,
	}

	payload, _ := json.MarshalIndent(info, "", "    ")
	fmt.Fprintln(w, string(payload))
}
