package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/M00nF1sh/echoserver/pkg/handler"
)

func main() {
	argPorts := flag.String("ports", "8080", "comma separated list of ports to listen on")
	flag.Parse()

	ports := strings.Split(*argPorts, ",")
	serverList := make([]*http.Server, 0, len(ports))
	for _, port := range ports {

		server := &http.Server{Addr: ":" + port, Handler: handler.New(port)}
		serverList = append(serverList, server)
		fmt.Printf("listen on %v \n", port)
		go server.ListenAndServe()
	}


	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	<-gracefulStop

	fmt.Println("start shutdown servers")
	for _, server := range serverList {
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		_ = server.Shutdown(ctx)
	}
	fmt.Println("complete shutdown servers")
}
