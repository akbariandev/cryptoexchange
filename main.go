package main

import (
	v1 "github.com/akbarian.dev/cryptoexchange/internal/api/v1"
	"github.com/akbarian.dev/cryptoexchange/pkg/logger"
	"github.com/akbarian.dev/cryptoexchange/pkg/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	logger.InitializeLogger()

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
