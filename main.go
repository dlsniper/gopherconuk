package main

import (
	"log"
	"net/http"
	"os"

	"github.com/PeloDev/go-server-template/examples"
	"github.com/PeloDev/go-server-template/homepage"
	"github.com/rs/cors"
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
	egs := examples.NewHandlers(logger)

	mux := http.NewServeMux()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "PUT", "OPTIONS", "POST", "DELETE"},
		AllowCredentials: true,
	})

	// decorate existing handler with cors functionality set in c
	handler := c.Handler(mux)

	// feature routes
	h.SetupRoutes(mux)
	egs.SetupRoutes(mux)

	logger.Println("server starting")

	// srv := server.New(mux, ServiceAddr)
	// err := srv.ListenAndServeTLS(CertFile, KeyFile) // TODO: serve via TLS (securely)
	err := http.ListenAndServe(ServiceAddr, handler)
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}
