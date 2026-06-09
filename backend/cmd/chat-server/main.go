package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/pedrovasalmeida/website-realtime-chat/backend/internal/chat"
	"github.com/pedrovasalmeida/website-realtime-chat/backend/internal/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	httpServer := &http.Server{
		Addr:              ":" + port,
		Handler:           server.NewRouter(chat.NewHub()),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("chat server listening on :%s", port)
	if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("server stopped: %v", err)
	}
}
