package handler

import (
	"net/http"
	"strconv"
	"time"
)

const queryDelay = "delay"

type Handler struct {
	Port string
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	delay := r.URL.Query().Get(queryDelay)
	if delay != "" {
		delayInt, _ := strconv.ParseInt(delay, 10, 64)
		time.Sleep(time.Duration(delayInt) * time.Second)
	}

	h.PrintRequestInfo(w, r)
	h.PrintServerInfo(w, r)
}

// New returns a new server
func New(port string) http.Handler {
	return &Handler{Port: port}
}
