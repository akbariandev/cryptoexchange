package main

import (
	v1 "gitlab.com/hotelian-company/challenge/internal/api/v1"
	"gitlab.com/hotelian-company/challenge/pkg/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	handler := http.NewServeMux()
	v1.NewRouter(handler)
	httpServer := server.New(handler)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case <-interrupt:
	}

	err := httpServer.Shutdown()
	if err != nil {
		log.Fatalf("error shutdown server : %v", err)
	}

}
