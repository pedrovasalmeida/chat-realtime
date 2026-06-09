package ws

import (
	"context"
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/pedrovasalmeida/website-realtime-chat/backend/internal/chat"
	"github.com/pedrovasalmeida/website-realtime-chat/backend/internal/protocol"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type Handler struct {
	hub        *chat.Hub
	nextUserID func() string
}

type Option func(*Handler)

func WithUserIDGenerator(generator func() string) Option {
	return func(handler *Handler) {
		handler.nextUserID = generator
	}
}

func NewHandler(hub *chat.Hub, options ...Option) *Handler {
	var sequence uint64
	handler := &Handler{
		hub: hub,
		nextUserID: func() string {
			next := atomic.AddUint64(&sequence, 1)
			return fmt.Sprintf("u%d", next)
		},
	}
	for _, option := range options {
		option(handler)
	}
	return handler
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{InsecureSkipVerify: true})
	if err != nil {
		return
	}
	defer conn.Close(websocket.StatusNormalClosure, "connection closed")

	ctx := r.Context()
	firstEvent, err := readClientEvent(ctx, conn)
	if err != nil {
		return
	}
	if firstEvent.Type != protocol.EventJoin {
		writeServerEvent(ctx, conn, protocol.ServerEvent{Type: protocol.EventError, Error: "first event must be join"})
		return
	}

	userID := h.nextUserID()
	client, err := h.hub.Join(userID, firstEvent.Name)
	if err != nil {
		writeServerEvent(ctx, conn, protocol.ServerEvent{Type: protocol.EventError, Error: err.Error()})
		return
	}
	defer h.hub.Leave(userID)

	go func() {
		for event := range client.Send {
			if err := writeServerEvent(ctx, conn, event); err != nil {
				return
			}
		}
	}()

	for {
		event, err := readClientEvent(ctx, conn)
		if err != nil {
			return
		}

		switch event.Type {
		case protocol.EventMessage:
			if _, err := h.hub.SendMessage(userID, event.Content); err != nil {
				h.hub.SendTo(userID, protocol.ServerEvent{Type: protocol.EventError, Error: err.Error()})
			}
		default:
			h.hub.SendTo(userID, protocol.ServerEvent{Type: protocol.EventError, Error: "unsupported event type"})
		}
	}
}

func readClientEvent(ctx context.Context, conn *websocket.Conn) (protocol.ClientEvent, error) {
	var event protocol.ClientEvent
	if err := wsjson.Read(ctx, conn, &event); err != nil {
		return protocol.ClientEvent{}, err
	}
	return event, nil
}

func writeServerEvent(ctx context.Context, conn *websocket.Conn, event protocol.ServerEvent) error {
	return wsjson.Write(ctx, conn, event)
}
