package main

import (
	"log"
	"net/http"
	"os"

	"github.com/PeloDev/go-server-template/homepage"
	"github.com/PeloDev/go-server-template/server"
)

var (
	// TODO: serve via TLS (securely)
	// CertFile    = os.Getenv("GO_CERT_FILE")
	// KeyFile     = os.Getenv("GO_KEY_FILE")
	ServiceAddr = os.Getenv("GO_SERVICE_ADDR")
)

func main() {
	if len(ServiceAddr) < 1 {
		ServiceAddr = ":8080"
	}

	logger := log.New(os.Stdout, "goServerT ", log.LstdFlags|log.Lshortfile)

	// define features with logger
	h := homepage.NewHandlers(logger)

	mux := http.NewServeMux()

	// feature routes
	h.SetupRoutes(mux)

	srv := server.New(mux, ServiceAddr)

	logger.Println("server starting")
	// err := srv.ListenAndServeTLS(CertFile, KeyFile) // TODO: serve via TLS (securely)
	err := srv.ListenAndServe()
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}
