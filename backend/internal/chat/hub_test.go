package chat

import (
	"testing"
	"time"

	"github.com/pedrovasalmeida/website-realtime-chat/backend/internal/protocol"
)

func TestHubJoinLeaveBroadcastsPresence(t *testing.T) {
	t.Parallel()

	hub := NewHub()

	alice, err := hub.Join("u1", "Alice")
	if err != nil {
		t.Fatalf("join alice: %v", err)
	}
	assertPresence(t, readEvent(t, alice.Send), []protocol.User{{ID: "u1", Name: "Alice"}})

	bob, err := hub.Join("u2", "Bob")
	if err != nil {
		t.Fatalf("join bob: %v", err)
	}
	wantBoth := []protocol.User{{ID: "u1", Name: "Alice"}, {ID: "u2", Name: "Bob"}}
	assertPresence(t, readEvent(t, alice.Send), wantBoth)
	assertPresence(t, readEvent(t, bob.Send), wantBoth)

	hub.Leave("u2")
	assertPresence(t, readEvent(t, alice.Send), []protocol.User{{ID: "u1", Name: "Alice"}})
}

func TestHubSendMessageBroadcastsServerMessage(t *testing.T) {
	t.Parallel()

	fixedTime := time.Date(2026, 6, 9, 15, 30, 0, 0, time.UTC)
	hub := NewHub(
		WithClock(func() time.Time { return fixedTime }),
		WithIDGenerator(func() string { return "m1" }),
	)

	alice, err := hub.Join("u1", "Alice")
	if err != nil {
		t.Fatalf("join alice: %v", err)
	}
	readEvent(t, alice.Send)

	bob, err := hub.Join("u2", "Bob")
	if err != nil {
		t.Fatalf("join bob: %v", err)
	}
	readEvent(t, alice.Send)
	readEvent(t, bob.Send)

	msg, err := hub.SendMessage("u1", "  hello everyone  ")
	if err != nil {
		t.Fatalf("send message: %v", err)
	}

	if msg.ID != "m1" || msg.UserID != "u1" || msg.UserName != "Alice" || msg.Content != "hello everyone" || !msg.SentAt.Equal(fixedTime) {
		t.Fatalf("unexpected message: %#v", msg)
	}

	assertMessage(t, readEvent(t, alice.Send), msg)
	assertMessage(t, readEvent(t, bob.Send), msg)
}

func TestHubRejectsInvalidMessage(t *testing.T) {
	t.Parallel()

	hub := NewHub()

	if _, err := hub.SendMessage("missing", "hello"); err == nil {
		t.Fatal("expected error for missing user")
	}

	alice, err := hub.Join("u1", "Alice")
	if err != nil {
		t.Fatalf("join alice: %v", err)
	}
	readEvent(t, alice.Send)

	if _, err := hub.SendMessage("u1", "   "); err == nil {
		t.Fatal("expected error for blank message")
	}
}

func readEvent(t *testing.T, ch <-chan protocol.ServerEvent) protocol.ServerEvent {
	t.Helper()

	select {
	case event := <-ch:
		return event
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for event")
		return protocol.ServerEvent{}
	}
}

func assertPresence(t *testing.T, got protocol.ServerEvent, want []protocol.User) {
	t.Helper()

	if got.Type != protocol.EventPresence {
		t.Fatalf("expected presence event, got %#v", got)
	}
	if len(got.Users) != len(want) {
		t.Fatalf("expected %d users, got %d: %#v", len(want), len(got.Users), got.Users)
	}
	for i := range want {
		if got.Users[i] != want[i] {
			t.Fatalf("user %d: expected %#v, got %#v", i, want[i], got.Users[i])
		}
	}
}

func assertMessage(t *testing.T, got protocol.ServerEvent, want protocol.Message) {
	t.Helper()

	if got.Type != protocol.EventMessage {
		t.Fatalf("expected message event, got %#v", got)
	}
	if got.Message == nil {
		t.Fatal("expected message payload")
	}
	if *got.Message != want {
		t.Fatalf("expected %#v, got %#v", want, *got.Message)
	}
}
