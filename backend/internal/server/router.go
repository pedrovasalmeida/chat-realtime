package server

import (
	"fmt"
	"net/http"

	"github.com/pedrovasalmeida/website-realtime-chat/backend/internal/chat"
	"github.com/pedrovasalmeida/website-realtime-chat/backend/internal/ws"
)

func NewRouter(hub *chat.Hub) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintln(w, "ok")
	})
	mux.Handle("GET /ws", ws.NewHandler(hub))
	return mux
}
