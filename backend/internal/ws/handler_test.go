package ws

import (
	"context"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/pedrovasalmeida/website-realtime-chat/backend/internal/chat"
	"github.com/pedrovasalmeida/website-realtime-chat/backend/internal/protocol"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func TestHandlerJoinMessagePresenceAndDisconnect(t *testing.T) {
	hub := chat.NewHub(chat.WithIDGenerator(func() string { return "m1" }))
	handler := NewHandler(hub, WithUserIDGenerator(sequence("u1", "u2")))
	server := httptest.NewServer(handler)
	defer server.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	alice := dialClient(t, ctx, server.URL)
	defer alice.Close(websocket.StatusNormalClosure, "test complete")
	writeClientEvent(t, ctx, alice, protocol.ClientEvent{Type: protocol.EventJoin, Name: "Alice"})
	assertPresenceEvent(t, readServerEvent(t, ctx, alice), []protocol.User{{ID: "u1", Name: "Alice"}})

	bob := dialClient(t, ctx, server.URL)
	writeClientEvent(t, ctx, bob, protocol.ClientEvent{Type: protocol.EventJoin, Name: "Bob"})
	assertPresenceEvent(t, readServerEvent(t, ctx, alice), []protocol.User{{ID: "u1", Name: "Alice"}, {ID: "u2", Name: "Bob"}})
	assertPresenceEvent(t, readServerEvent(t, ctx, bob), []protocol.User{{ID: "u1", Name: "Alice"}, {ID: "u2", Name: "Bob"}})

	writeClientEvent(t, ctx, alice, protocol.ClientEvent{Type: protocol.EventMessage, Content: "hello"})
	assertMessageEvent(t, readServerEvent(t, ctx, alice), "m1", "u1", "Alice", "hello")
	assertMessageEvent(t, readServerEvent(t, ctx, bob), "m1", "u1", "Alice", "hello")

	if err := bob.Close(websocket.StatusNormalClosure, "leaving"); err != nil {
		t.Fatalf("close bob: %v", err)
	}
	assertPresenceEvent(t, readServerEvent(t, ctx, alice), []protocol.User{{ID: "u1", Name: "Alice"}})
}

func TestHandlerRejectsMessageBeforeJoin(t *testing.T) {
	hub := chat.NewHub()
	server := httptest.NewServer(NewHandler(hub, WithUserIDGenerator(sequence("u1"))))
	defer server.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	conn := dialClient(t, ctx, server.URL)
	defer conn.Close(websocket.StatusNormalClosure, "test complete")

	writeClientEvent(t, ctx, conn, protocol.ClientEvent{Type: protocol.EventMessage, Content: "hello"})
	event := readServerEvent(t, ctx, conn)
	if event.Type != protocol.EventError || event.Error != "first event must be join" {
		t.Fatalf("unexpected event: %#v", event)
	}
}

func TestHandlerReportsInvalidMessage(t *testing.T) {
	hub := chat.NewHub()
	server := httptest.NewServer(NewHandler(hub, WithUserIDGenerator(sequence("u1"))))
	defer server.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	conn := dialClient(t, ctx, server.URL)
	defer conn.Close(websocket.StatusNormalClosure, "test complete")

	writeClientEvent(t, ctx, conn, protocol.ClientEvent{Type: protocol.EventJoin, Name: "Alice"})
	readServerEvent(t, ctx, conn)
	writeClientEvent(t, ctx, conn, protocol.ClientEvent{Type: protocol.EventMessage, Content: "   "})

	event := readServerEvent(t, ctx, conn)
	if event.Type != protocol.EventError || event.Error != protocol.ErrMessageRequired.Error() {
		t.Fatalf("unexpected event: %#v", event)
	}
}

func dialClient(t *testing.T, ctx context.Context, serverURL string) *websocket.Conn {
	t.Helper()

	conn, _, err := websocket.Dial(ctx, "ws"+serverURL[len("http"):], nil)
	if err != nil {
		t.Fatalf("dial websocket: %v", err)
	}
	return conn
}

func writeClientEvent(t *testing.T, ctx context.Context, conn *websocket.Conn, event protocol.ClientEvent) {
	t.Helper()

	if err := wsjson.Write(ctx, conn, event); err != nil {
		t.Fatalf("write event: %v", err)
	}
}

func readServerEvent(t *testing.T, ctx context.Context, conn *websocket.Conn) protocol.ServerEvent {
	t.Helper()

	var event protocol.ServerEvent
	if err := wsjson.Read(ctx, conn, &event); err != nil {
		t.Fatalf("read event: %v", err)
	}
	return event
}

func assertPresenceEvent(t *testing.T, event protocol.ServerEvent, want []protocol.User) {
	t.Helper()

	if event.Type != protocol.EventPresence {
		t.Fatalf("expected presence, got %#v", event)
	}
	if len(event.Users) != len(want) {
		t.Fatalf("expected users %#v, got %#v", want, event.Users)
	}
	for i := range want {
		if event.Users[i] != want[i] {
			t.Fatalf("expected users %#v, got %#v", want, event.Users)
		}
	}
}

func assertMessageEvent(t *testing.T, event protocol.ServerEvent, id string, userID string, userName string, content string) {
	t.Helper()

	if event.Type != protocol.EventMessage || event.Message == nil {
		t.Fatalf("expected message, got %#v", event)
	}
	if event.Message.ID != id || event.Message.UserID != userID || event.Message.UserName != userName || event.Message.Content != content {
		t.Fatalf("unexpected message: %#v", event.Message)
	}
}

func sequence(values ...string) func() string {
	index := 0
	return func() string {
		value := values[index]
		index++
		return value
	}
}
