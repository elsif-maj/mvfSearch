package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/elsif-maj/umbraSearch/app"
	"github.com/elsif-maj/umbraSearch/handler"
)

func main() {
	// ENV and DB
	server, err := app.Setup()
	if err != nil {
		log.Printf("Failed to setup search service: %v", err)
		os.Exit(1)
	}
	defer server.DBConn.Close(context.Background())
	defer server.KVConn.Close()
	// Routes
	handler.SetupRoutes(server)
	// Server Loop
	http.ListenAndServe(":3000", nil)
}
