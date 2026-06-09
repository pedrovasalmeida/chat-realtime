# Realtime Chat Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Build the first working realtime chat with Docker Compose, Go/WebSocket backend, React/Vite/Tailwind frontend, backend tests, frontend unit tests, and wiki ingest.

**Architecture:** Use a simple monorepo with `backend/`, `frontend/`, and Docker files at the root. The Go backend owns the in-memory chat hub, presence, validation, and WebSocket endpoint; the React frontend renders the approved classic chat UI and talks to `/ws`.

**Tech Stack:** Go 1.22, `net/http`, `nhooyr.io/websocket`, React, Vite, TypeScript, Tailwind CSS, Vitest, Testing Library, Docker Compose, Air.

---

## Execution Rules

- This directory is not a Git repository. Do not run `git init`, do not create commits, and do not add AI/Copilot/coauthor attribution. Each task ends with a checkpoint instead of a commit.
- Use TDD for backend and frontend unit behavior: write the failing test, run it, implement the minimal code, rerun it.
- Before implementing Task 9 UI styling, invoke the `frontend-design` skill.
- After code implementation, update the wiki and append to `docs/wiki/log.md`.

## File Structure

### Root

- Create `.gitignore`: ignore local dependencies, build outputs, environment files, and `.superpowers/`.
- Create `.env.example`: document local ports and frontend WebSocket URL.
- Create `docker-compose.yml`: dev orchestration for backend and frontend.
- Modify `README.md`: add local run/test commands after implementation.

### Backend

- Create `backend/go.mod`: Go module and WebSocket dependency.
- Create `backend/Dockerfile`: dev container with Air and production build target.
- Create `backend/.air.toml`: hot reload configuration.
- Create `backend/cmd/chat-server/main.go`: server entrypoint.
- Create `backend/internal/protocol/events.go`: event constants, payload types, validation.
- Create `backend/internal/protocol/events_test.go`: protocol validation tests.
- Create `backend/internal/chat/hub.go`: in-memory clients, presence, and broadcast.
- Create `backend/internal/chat/hub_test.go`: hub behavior tests.
- Create `backend/internal/ws/handler.go`: WebSocket lifecycle and JSON event handling.
- Create `backend/internal/ws/handler_test.go`: WebSocket handler tests.
- Create `backend/internal/server/router.go`: HTTP router and health endpoint.
- Create `backend/internal/server/router_test.go`: health endpoint test.

### Frontend

- Create `frontend/package.json`: npm scripts and dependencies.
- Create `frontend/index.html`: Vite entry.
- Create `frontend/tsconfig.json`, `frontend/tsconfig.node.json`: TypeScript config.
- Create `frontend/vite.config.ts`: Vite and Vitest config.
- Create `frontend/vitest.setup.ts`: Testing Library setup.
- Create `frontend/tailwind.config.js`, `frontend/postcss.config.js`: Tailwind setup.
- Create `frontend/src/main.tsx`: React entrypoint.
- Create `frontend/src/index.css`: Tailwind and page styles.
- Create `frontend/src/types/chat.ts`: shared frontend protocol types.
- Create `frontend/src/chat/chatReducer.ts`: deterministic chat state reducer.
- Create `frontend/src/chat/chatReducer.test.ts`: reducer unit tests.
- Create `frontend/src/hooks/useChatConnection.ts`: WebSocket hook with reconnect.
- Create `frontend/src/components/ConnectionStatus.tsx` and test.
- Create `frontend/src/components/JoinForm.tsx` and test.
- Create `frontend/src/components/MessageComposer.tsx` and test.
- Create `frontend/src/components/MessageList.tsx` and test.
- Create `frontend/src/components/PeopleList.tsx` and test.
- Create `frontend/src/App.tsx`: app composition.

### Documentation

- Modify `docs/wiki/project-overview.md`: implementation status and validation commands.
- Modify `docs/wiki/architecture.md`: final file layout and runtime flow.
- Modify `docs/wiki/backend.md`: implemented backend modules and tests.
- Modify `docs/wiki/frontend.md`: implemented UI, hook, and tests.
- Modify `docs/wiki/realtime-protocol.md`: final JSON event examples.
- Modify `docs/wiki/docker.md`: Docker Compose commands and ports.
- Modify `docs/wiki/testing.md`: actual test commands.
- Modify `docs/wiki/index.md`: refresh summaries if material wording changes.
- Modify `docs/wiki/log.md`: append implementation completion entry.

---

## Task 1: Docker and repository development foundation

**Files:**
- Create: `.gitignore`
- Create: `.env.example`
- Create: `docker-compose.yml`
- Create: `backend/Dockerfile`
- Create: `backend/.air.toml`
- Create: `frontend/Dockerfile`

- [ ] **Step 1: Create root local-development files**

Create `.gitignore`:

```gitignore
.DS_Store
.env
.superpowers/

backend/tmp/
backend/bin/

frontend/node_modules/
frontend/dist/
frontend/coverage/

*.log
```

Create `.env.example`:

```dotenv
BACKEND_PORT=8080
FRONTEND_PORT=5173
VITE_WS_URL=ws://localhost:8080/ws
```

- [ ] **Step 2: Create Docker Compose config**

Create `docker-compose.yml`:

```yaml
services:
  backend:
    build:
      context: ./backend
      target: dev
    environment:
      PORT: "8080"
    ports:
      - "8080:8080"
    volumes:
      - ./backend:/app
      - go-mod-cache:/go/pkg/mod

  frontend:
    build:
      context: ./frontend
      target: dev
    environment:
      VITE_WS_URL: ws://localhost:8080/ws
    ports:
      - "5173:5173"
    volumes:
      - ./frontend:/app
      - frontend-node-modules:/app/node_modules
    depends_on:
      - backend

volumes:
  go-mod-cache:
  frontend-node-modules:
```

- [ ] **Step 3: Create backend dev Docker files**

Create `backend/Dockerfile`:

```dockerfile
FROM golang:1.22-alpine AS dev

WORKDIR /app
RUN apk add --no-cache git && go install github.com/air-verse/air@v1.52.3
EXPOSE 8080
CMD ["air", "-c", ".air.toml"]

FROM golang:1.22-alpine AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /chat-server ./cmd/chat-server

FROM alpine:3.20 AS runtime

WORKDIR /app
COPY --from=build /chat-server /app/chat-server
EXPOSE 8080
CMD ["/app/chat-server"]
```

Create `backend/.air.toml`:

```toml
root = "."
tmp_dir = "tmp"

[build]
cmd = "go build -o ./tmp/chat-server ./cmd/chat-server"
bin = "./tmp/chat-server"
include_ext = ["go"]
exclude_dir = ["tmp"]
delay = 1000
stop_on_error = true

[log]
time = true
```

- [ ] **Step 4: Create frontend dev Docker file**

Create `frontend/Dockerfile`:

```dockerfile
FROM node:22-alpine AS dev

WORKDIR /app
COPY package*.json ./
RUN npm install
EXPOSE 5173
CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]

FROM node:22-alpine AS build

WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build
```

- [ ] **Step 5: Validate Compose syntax**

Run:

```bash
docker compose config
```

Expected: command exits with status 0 and prints normalized Compose YAML.

- [ ] **Step 6: Checkpoint**

Record in the task notes that Docker/dev foundation files exist and Compose syntax validates. Do not create a Git commit in this non-Git directory.

---

## Task 2: Backend protocol and validation

**Files:**
- Create: `backend/go.mod`
- Create: `backend/internal/protocol/events_test.go`
- Create: `backend/internal/protocol/events.go`

- [ ] **Step 1: Create Go module**

Create `backend/go.mod`:

```go
module github.com/pedrovasalmeida/website-realtime-chat/backend

go 1.22

require nhooyr.io/websocket v1.8.17
```

- [ ] **Step 2: Write failing protocol tests**

Create `backend/internal/protocol/events_test.go`:

```go
package protocol

import (
	"strings"
	"testing"
)

func TestValidateJoinName(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{name: "trims name", input: "  Pedro  ", want: "Pedro"},
		{name: "rejects empty", input: "   ", wantErr: true},
		{name: "rejects long name", input: strings.Repeat("a", MaxNameLength+1), wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateJoinName(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if got != tt.want {
				t.Fatalf("expected %q, got %q", tt.want, got)
			}
		})
	}
}

func TestValidateMessageContent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{name: "trims content", input: "  hello  ", want: "hello"},
		{name: "rejects empty", input: "\n\t", wantErr: true},
		{name: "rejects long content", input: strings.Repeat("a", MaxMessageLength+1), wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateMessageContent(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if got != tt.want {
				t.Fatalf("expected %q, got %q", tt.want, got)
			}
		})
	}
}

func TestEventConstants(t *testing.T) {
	t.Parallel()

	if EventJoin != "join" {
		t.Fatalf("EventJoin = %q", EventJoin)
	}
	if EventMessage != "message" {
		t.Fatalf("EventMessage = %q", EventMessage)
	}
	if EventPresence != "presence" {
		t.Fatalf("EventPresence = %q", EventPresence)
	}
	if EventError != "error" {
		t.Fatalf("EventError = %q", EventError)
	}
}
```

- [ ] **Step 3: Run protocol tests to verify failure**

Run:

```bash
cd backend && go test ./internal/protocol -run 'TestValidate|TestEvent' -v
```

Expected: FAIL because `ValidateJoinName`, `ValidateMessageContent`, constants, and types are undefined.

- [ ] **Step 4: Implement protocol types and validation**

Create `backend/internal/protocol/events.go`:

```go
package protocol

import (
	"errors"
	"strings"
	"time"
)

const (
	EventJoin     = "join"
	EventMessage  = "message"
	EventPresence = "presence"
	EventError    = "error"

	MaxNameLength    = 40
	MaxMessageLength = 500
)

var (
	ErrNameRequired       = errors.New("name is required")
	ErrNameTooLong        = errors.New("name is too long")
	ErrMessageRequired    = errors.New("message content is required")
	ErrMessageContentLong = errors.New("message content is too long")
)

type ClientEvent struct {
	Type    string `json:"type"`
	Name    string `json:"name,omitempty"`
	Content string `json:"content,omitempty"`
}

type ServerEvent struct {
	Type    string   `json:"type"`
	Message *Message `json:"message,omitempty"`
	Users   []User   `json:"users,omitempty"`
	Error   string   `json:"error,omitempty"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Message struct {
	ID       string    `json:"id"`
	UserID   string    `json:"userId"`
	UserName string    `json:"userName"`
	Content  string    `json:"content"`
	SentAt   time.Time `json:"sentAt"`
}

func ValidateJoinName(input string) (string, error) {
	name := strings.TrimSpace(input)
	if name == "" {
		return "", ErrNameRequired
	}
	if len([]rune(name)) > MaxNameLength {
		return "", ErrNameTooLong
	}
	return name, nil
}

func ValidateMessageContent(input string) (string, error) {
	content := strings.TrimSpace(input)
	if content == "" {
		return "", ErrMessageRequired
	}
	if len([]rune(content)) > MaxMessageLength {
		return "", ErrMessageContentLong
	}
	return content, nil
}
```

- [ ] **Step 5: Run protocol tests to verify pass**

Run:

```bash
cd backend && go test ./internal/protocol -v
```

Expected: PASS.

- [ ] **Step 6: Checkpoint**

Record that protocol constants, JSON types, and validation rules are implemented and tested.

---

## Task 3: Backend in-memory chat hub

**Files:**
- Create: `backend/internal/chat/hub_test.go`
- Create: `backend/internal/chat/hub.go`

- [ ] **Step 1: Write failing hub tests**

Create `backend/internal/chat/hub_test.go`:

```go
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
```

- [ ] **Step 2: Run hub tests to verify failure**

Run:

```bash
cd backend && go test ./internal/chat -v
```

Expected: FAIL because `NewHub`, `WithClock`, `WithIDGenerator`, `Join`, `Leave`, and `SendMessage` are undefined.

- [ ] **Step 3: Implement chat hub**

Create `backend/internal/chat/hub.go`:

```go
package chat

import (
	"errors"
	"fmt"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	"github.com/pedrovasalmeida/website-realtime-chat/backend/internal/protocol"
)

var (
	ErrUserAlreadyConnected = errors.New("user is already connected")
	ErrUserNotConnected     = errors.New("user is not connected")
)

type Client struct {
	User protocol.User
	Send chan protocol.ServerEvent
}

type Hub struct {
	mu      sync.RWMutex
	clients map[string]*Client
	clock   func() time.Time
	newID   func() string
}

type Option func(*Hub)

func WithClock(clock func() time.Time) Option {
	return func(h *Hub) {
		h.clock = clock
	}
}

func WithIDGenerator(generator func() string) Option {
	return func(h *Hub) {
		h.newID = generator
	}
}

func NewHub(options ...Option) *Hub {
	var sequence uint64
	hub := &Hub{
		clients: make(map[string]*Client),
		clock:   time.Now,
		newID: func() string {
			next := atomic.AddUint64(&sequence, 1)
			return fmt.Sprintf("m%d", next)
		},
	}
	for _, option := range options {
		option(hub)
	}
	return hub
}

func (h *Hub) Join(userID string, rawName string) (*Client, error) {
	name, err := protocol.ValidateJoinName(rawName)
	if err != nil {
		return nil, err
	}

	h.mu.Lock()
	if _, exists := h.clients[userID]; exists {
		h.mu.Unlock()
		return nil, ErrUserAlreadyConnected
	}
	client := &Client{
		User: protocol.User{ID: userID, Name: name},
		Send: make(chan protocol.ServerEvent, 16),
	}
	h.clients[userID] = client
	users := h.usersLocked()
	h.mu.Unlock()

	h.broadcast(protocol.ServerEvent{Type: protocol.EventPresence, Users: users})
	return client, nil
}

func (h *Hub) Leave(userID string) {
	h.mu.Lock()
	client, exists := h.clients[userID]
	if exists {
		delete(h.clients, userID)
		close(client.Send)
	}
	users := h.usersLocked()
	h.mu.Unlock()

	if exists {
		h.broadcast(protocol.ServerEvent{Type: protocol.EventPresence, Users: users})
	}
}

func (h *Hub) Users() []protocol.User {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return h.usersLocked()
}

func (h *Hub) SendMessage(userID string, rawContent string) (protocol.Message, error) {
	content, err := protocol.ValidateMessageContent(rawContent)
	if err != nil {
		return protocol.Message{}, err
	}

	h.mu.RLock()
	client, exists := h.clients[userID]
	h.mu.RUnlock()
	if !exists {
		return protocol.Message{}, ErrUserNotConnected
	}

	message := protocol.Message{
		ID:       h.newID(),
		UserID:   client.User.ID,
		UserName: client.User.Name,
		Content:  content,
		SentAt:   h.clock().UTC(),
	}

	h.broadcast(protocol.ServerEvent{Type: protocol.EventMessage, Message: &message})
	return message, nil
}

func (h *Hub) SendTo(userID string, event protocol.ServerEvent) {
	h.mu.RLock()
	client := h.clients[userID]
	h.mu.RUnlock()
	if client == nil {
		return
	}
	deliver(client, event)
}

func (h *Hub) broadcast(event protocol.ServerEvent) {
	h.mu.RLock()
	clients := make([]*Client, 0, len(h.clients))
	for _, client := range h.clients {
		clients = append(clients, client)
	}
	h.mu.RUnlock()

	for _, client := range clients {
		deliver(client, event)
	}
}

func (h *Hub) usersLocked() []protocol.User {
	users := make([]protocol.User, 0, len(h.clients))
	for _, client := range h.clients {
		users = append(users, client.User)
	}
	sort.Slice(users, func(i, j int) bool {
		if users[i].Name == users[j].Name {
			return users[i].ID < users[j].ID
		}
		return users[i].Name < users[j].Name
	})
	return users
}

func deliver(client *Client, event protocol.ServerEvent) {
	select {
	case client.Send <- event:
	default:
	}
}
```

- [ ] **Step 4: Run hub tests to verify pass**

Run:

```bash
cd backend && go test ./internal/chat -v
```

Expected: PASS.

- [ ] **Step 5: Run all backend tests so far**

Run:

```bash
cd backend && go test ./...
```

Expected: PASS.

- [ ] **Step 6: Checkpoint**

Record that the in-memory hub handles join, leave, presence, and message broadcast with tests.

---

## Task 4: Backend WebSocket handler

**Files:**
- Create: `backend/internal/ws/handler_test.go`
- Create: `backend/internal/ws/handler.go`

- [ ] **Step 1: Write failing WebSocket handler tests**

Create `backend/internal/ws/handler_test.go`:

```go
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
```

- [ ] **Step 2: Run WebSocket tests to verify failure**

Run:

```bash
cd backend && go test ./internal/ws -v
```

Expected: FAIL because `NewHandler`, `WithUserIDGenerator`, and handler logic are undefined.

- [ ] **Step 3: Implement WebSocket handler**

Create `backend/internal/ws/handler.go`:

```go
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
```

- [ ] **Step 4: Run WebSocket tests**

Run:

```bash
cd backend && go test ./internal/ws -v
```

Expected: PASS.

- [ ] **Step 5: Run all backend tests**

Run:

```bash
cd backend && go test ./...
```

Expected: PASS.

- [ ] **Step 6: Checkpoint**

Record that `/ws` behavior is covered for join, message broadcast, presence updates, disconnect cleanup, and invalid payload errors.

---

## Task 5: Backend HTTP server entrypoint

**Files:**
- Create: `backend/internal/server/router_test.go`
- Create: `backend/internal/server/router.go`
- Create: `backend/cmd/chat-server/main.go`

- [ ] **Step 1: Write failing router test**

Create `backend/internal/server/router_test.go`:

```go
package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pedrovasalmeida/website-realtime-chat/backend/internal/chat"
)

func TestRouterHealthz(t *testing.T) {
	t.Parallel()

	router := NewRouter(chat.NewHub())
	request := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}
	if response.Body.String() != "ok\n" {
		t.Fatalf("expected ok body, got %q", response.Body.String())
	}
}
```

- [ ] **Step 2: Run router test to verify failure**

Run:

```bash
cd backend && go test ./internal/server -v
```

Expected: FAIL because `NewRouter` is undefined.

- [ ] **Step 3: Implement router**

Create `backend/internal/server/router.go`:

```go
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
```

- [ ] **Step 4: Create server entrypoint**

Create `backend/cmd/chat-server/main.go`:

```go
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
```

- [ ] **Step 5: Run backend tests and build**

Run:

```bash
cd backend && go mod tidy && go test ./... && go build ./cmd/chat-server
```

Expected: all commands exit with status 0. `go.sum` is created or updated by `go mod tidy`.

- [ ] **Step 6: Checkpoint**

Record that backend has a health endpoint, WebSocket route, entrypoint, tests, and build.

---

## Task 6: Frontend scaffold and test foundation

**Files:**
- Create: `frontend/package.json`
- Create: `frontend/index.html`
- Create: `frontend/tsconfig.json`
- Create: `frontend/tsconfig.node.json`
- Create: `frontend/vite.config.ts`
- Create: `frontend/vitest.setup.ts`
- Create: `frontend/tailwind.config.js`
- Create: `frontend/postcss.config.js`
- Create: `frontend/src/main.tsx`
- Create: `frontend/src/index.css`
- Create: `frontend/src/App.tsx`

- [ ] **Step 1: Create frontend package and config files**

Create `frontend/package.json`:

```json
{
  "name": "realtime-chat-frontend",
  "private": true,
  "version": "0.1.0",
  "type": "module",
  "scripts": {
    "dev": "vite --host 0.0.0.0",
    "build": "tsc -b && vite build",
    "test": "vitest",
    "test:run": "vitest run"
  },
  "dependencies": {
    "react": "^18.3.1",
    "react-dom": "^18.3.1"
  },
  "devDependencies": {
    "@testing-library/jest-dom": "^6.6.3",
    "@testing-library/react": "^16.1.0",
    "@testing-library/user-event": "^14.5.2",
    "@types/react": "^18.3.12",
    "@types/react-dom": "^18.3.1",
    "@vitejs/plugin-react": "^4.3.4",
    "autoprefixer": "^10.4.20",
    "jsdom": "^25.0.1",
    "postcss": "^8.4.49",
    "tailwindcss": "^3.4.17",
    "typescript": "^5.7.2",
    "vite": "^5.4.11",
    "vitest": "^2.1.8"
  }
}
```

Create `frontend/index.html`:

```html
<!doctype html>
<html lang="pt-BR">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Realtime Chat</title>
  </head>
  <body>
    <div id="root"></div>
    <script type="module" src="/src/main.tsx"></script>
  </body>
</html>
```

Create `frontend/tsconfig.json`:

```json
{
  "compilerOptions": {
    "target": "ES2020",
    "useDefineForClassFields": true,
    "lib": ["DOM", "DOM.Iterable", "ES2020"],
    "allowJs": false,
    "skipLibCheck": true,
    "esModuleInterop": true,
    "allowSyntheticDefaultImports": true,
    "strict": true,
    "forceConsistentCasingInFileNames": true,
    "module": "ESNext",
    "moduleResolution": "Node",
    "resolveJsonModule": true,
    "isolatedModules": true,
    "noEmit": true,
    "jsx": "react-jsx"
  },
  "include": ["src", "vitest.setup.ts"],
  "references": [{ "path": "./tsconfig.node.json" }]
}
```

Create `frontend/tsconfig.node.json`:

```json
{
  "compilerOptions": {
    "composite": true,
    "module": "ESNext",
    "moduleResolution": "Node",
    "allowSyntheticDefaultImports": true
  },
  "include": ["vite.config.ts"]
}
```

- [ ] **Step 2: Create Vite, Vitest, and Tailwind config**

Create `frontend/vite.config.ts`:

```ts
/// <reference types="vitest" />

import react from '@vitejs/plugin-react';
import { defineConfig } from 'vite';

export default defineConfig({
  plugins: [react()],
  test: {
    environment: 'jsdom',
    globals: true,
    passWithNoTests: true,
    setupFiles: './vitest.setup.ts',
  },
});
```

Create `frontend/vitest.setup.ts`:

```ts
import '@testing-library/jest-dom/vitest';
```

Create `frontend/tailwind.config.js`:

```js
/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{ts,tsx}'],
  theme: {
    extend: {},
  },
  plugins: [],
};
```

Create `frontend/postcss.config.js`:

```js
export default {
  plugins: {
    tailwindcss: {},
    autoprefixer: {},
  },
};
```

- [ ] **Step 3: Create minimal React entrypoint**

Create `frontend/src/index.css`:

```css
@tailwind base;
@tailwind components;
@tailwind utilities;

:root {
  color-scheme: light;
  font-family: Inter, ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif;
  background: #eef2ff;
  color: #0f172a;
}

body {
  margin: 0;
  min-width: 320px;
  min-height: 100vh;
}
```

Create `frontend/src/App.tsx`:

```tsx
export default function App() {
  return (
    <main className="min-h-screen bg-slate-100 p-6 text-slate-950">
      <h1 className="text-3xl font-bold">Realtime Chat</h1>
      <p className="mt-2 text-slate-600">Preparando interface do chat.</p>
    </main>
  );
}
```

Create `frontend/src/main.tsx`:

```tsx
import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';
import './index.css';

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
);
```

- [ ] **Step 4: Install dependencies**

Run:

```bash
cd frontend && npm install
```

Expected: command exits with status 0 and creates `package-lock.json`.

- [ ] **Step 5: Validate frontend foundation**

Run:

```bash
cd frontend && npm run test:run && npm run build
```

Expected: test command exits with status 0 because `passWithNoTests` is enabled, then build exits with status 0.

- [ ] **Step 6: Checkpoint**

Record that frontend scaffold, Tailwind, Vitest, and build foundation are ready.

---

## Task 7: Frontend protocol state reducer

**Files:**
- Create: `frontend/src/types/chat.ts`
- Create: `frontend/src/chat/chatReducer.test.ts`
- Create: `frontend/src/chat/chatReducer.ts`

- [ ] **Step 1: Create frontend protocol types**

Create `frontend/src/types/chat.ts`:

```ts
export type ConnectionStatus = 'idle' | 'connecting' | 'connected' | 'disconnected' | 'error';

export type User = {
  id: string;
  name: string;
};

export type ChatMessage = {
  id: string;
  userId: string;
  userName: string;
  content: string;
  sentAt: string;
};

export type ServerEvent =
  | { type: 'presence'; users: User[] }
  | { type: 'message'; message: ChatMessage }
  | { type: 'error'; error: string };

export type ClientEvent =
  | { type: 'join'; name: string }
  | { type: 'message'; content: string };

export type ChatState = {
  status: ConnectionStatus;
  users: User[];
  messages: ChatMessage[];
  error: string | null;
};
```

- [ ] **Step 2: Write failing reducer tests**

Create `frontend/src/chat/chatReducer.test.ts`:

```ts
import { describe, expect, it } from 'vitest';
import { chatReducer, initialChatState } from './chatReducer';

describe('chatReducer', () => {
  it('stores presence snapshots', () => {
    const state = chatReducer(initialChatState, {
      type: 'server-event',
      event: {
        type: 'presence',
        users: [
          { id: 'u1', name: 'Alice' },
          { id: 'u2', name: 'Bob' },
        ],
      },
    });

    expect(state.users).toEqual([
      { id: 'u1', name: 'Alice' },
      { id: 'u2', name: 'Bob' },
    ]);
  });

  it('appends broadcast messages', () => {
    const state = chatReducer(initialChatState, {
      type: 'server-event',
      event: {
        type: 'message',
        message: {
          id: 'm1',
          userId: 'u1',
          userName: 'Alice',
          content: 'hello',
          sentAt: '2026-06-09T18:30:00Z',
        },
      },
    });

    expect(state.messages).toHaveLength(1);
    expect(state.messages[0].content).toBe('hello');
  });

  it('stores server errors and connection state changes', () => {
    const connected = chatReducer(initialChatState, { type: 'status', status: 'connected' });
    const failed = chatReducer(connected, {
      type: 'server-event',
      event: { type: 'error', error: 'message content is required' },
    });

    expect(connected.status).toBe('connected');
    expect(failed.error).toBe('message content is required');
  });
});
```

- [ ] **Step 3: Run reducer tests to verify failure**

Run:

```bash
cd frontend && npm run test:run -- src/chat/chatReducer.test.ts
```

Expected: FAIL because `chatReducer` and `initialChatState` are undefined.

- [ ] **Step 4: Implement reducer**

Create `frontend/src/chat/chatReducer.ts`:

```ts
import type { ChatState, ConnectionStatus, ServerEvent } from '../types/chat';

export const initialChatState: ChatState = {
  status: 'idle',
  users: [],
  messages: [],
  error: null,
};

export type ChatAction =
  | { type: 'status'; status: ConnectionStatus }
  | { type: 'server-event'; event: ServerEvent }
  | { type: 'client-error'; error: string | null }
  | { type: 'reset' };

export function chatReducer(state: ChatState, action: ChatAction): ChatState {
  switch (action.type) {
    case 'status':
      return { ...state, status: action.status, error: action.status === 'connected' ? null : state.error };
    case 'server-event':
      return applyServerEvent(state, action.event);
    case 'client-error':
      return { ...state, error: action.error };
    case 'reset':
      return initialChatState;
    default:
      return state;
  }
}

function applyServerEvent(state: ChatState, event: ServerEvent): ChatState {
  switch (event.type) {
    case 'presence':
      return { ...state, users: event.users };
    case 'message':
      return { ...state, messages: [...state.messages, event.message] };
    case 'error':
      return { ...state, error: event.error };
    default:
      return state;
  }
}
```

- [ ] **Step 5: Run reducer tests to verify pass**

Run:

```bash
cd frontend && npm run test:run -- src/chat/chatReducer.test.ts
```

Expected: PASS.

- [ ] **Step 6: Checkpoint**

Record that frontend protocol types and reducer behavior are unit-tested.

---

## Task 8: Frontend chat components

**Files:**
- Create: `frontend/src/components/ConnectionStatus.test.tsx`
- Create: `frontend/src/components/ConnectionStatus.tsx`
- Create: `frontend/src/components/JoinForm.test.tsx`
- Create: `frontend/src/components/JoinForm.tsx`
- Create: `frontend/src/components/MessageComposer.test.tsx`
- Create: `frontend/src/components/MessageComposer.tsx`
- Create: `frontend/src/components/MessageList.test.tsx`
- Create: `frontend/src/components/MessageList.tsx`
- Create: `frontend/src/components/PeopleList.test.tsx`
- Create: `frontend/src/components/PeopleList.tsx`

- [ ] **Step 1: Write failing component tests**

Create `frontend/src/components/ConnectionStatus.test.tsx`:

```tsx
import { render, screen } from '@testing-library/react';
import { describe, expect, it } from 'vitest';
import ConnectionStatus from './ConnectionStatus';

describe('ConnectionStatus', () => {
  it('shows connected state', () => {
    render(<ConnectionStatus status="connected" error={null} />);
    expect(screen.getByText('Conectado')).toBeInTheDocument();
  });

  it('shows error text', () => {
    render(<ConnectionStatus status="error" error="falha de conexao" />);
    expect(screen.getByText('Erro: falha de conexao')).toBeInTheDocument();
  });
});
```

Create `frontend/src/components/JoinForm.test.tsx`:

```tsx
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { describe, expect, it, vi } from 'vitest';
import JoinForm from './JoinForm';

describe('JoinForm', () => {
  it('submits trimmed name', async () => {
    const onJoin = vi.fn();
    const user = userEvent.setup();
    render(<JoinForm onJoin={onJoin} />);

    await user.type(screen.getByLabelText('Seu nome'), '  Pedro  ');
    await user.click(screen.getByRole('button', { name: 'Entrar no chat' }));

    expect(onJoin).toHaveBeenCalledWith('Pedro');
  });
});
```

Create `frontend/src/components/MessageComposer.test.tsx`:

```tsx
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { describe, expect, it, vi } from 'vitest';
import MessageComposer from './MessageComposer';

describe('MessageComposer', () => {
  it('sends trimmed messages and clears the input', async () => {
    const onSend = vi.fn();
    const user = userEvent.setup();
    render(<MessageComposer disabled={false} onSend={onSend} />);

    await user.type(screen.getByLabelText('Mensagem'), '  ola  ');
    await user.click(screen.getByRole('button', { name: 'Enviar' }));

    expect(onSend).toHaveBeenCalledWith('ola');
    expect(screen.getByLabelText('Mensagem')).toHaveValue('');
  });

  it('disables send button when disconnected', () => {
    render(<MessageComposer disabled={true} onSend={() => undefined} />);
    expect(screen.getByRole('button', { name: 'Enviar' })).toBeDisabled();
  });
});
```

Create `frontend/src/components/MessageList.test.tsx`:

```tsx
import { render, screen } from '@testing-library/react';
import { describe, expect, it } from 'vitest';
import MessageList from './MessageList';

describe('MessageList', () => {
  it('renders author, content, and time', () => {
    render(
      <MessageList
        messages={[
          {
            id: 'm1',
            userId: 'u1',
            userName: 'Alice',
            content: 'ola',
            sentAt: '2026-06-09T18:30:00Z',
          },
        ]}
      />,
    );

    expect(screen.getByText('Alice')).toBeInTheDocument();
    expect(screen.getByText('ola')).toBeInTheDocument();
    expect(screen.getByText(/\d{2}:\d{2}/)).toBeInTheDocument();
  });
});
```

Create `frontend/src/components/PeopleList.test.tsx`:

```tsx
import { render, screen } from '@testing-library/react';
import { describe, expect, it } from 'vitest';
import PeopleList from './PeopleList';

describe('PeopleList', () => {
  it('renders online people', () => {
    render(<PeopleList users={[{ id: 'u1', name: 'Alice' }, { id: 'u2', name: 'Bob' }]} />);

    expect(screen.getByText('Pessoas online')).toBeInTheDocument();
    expect(screen.getByText('Alice')).toBeInTheDocument();
    expect(screen.getByText('Bob')).toBeInTheDocument();
  });
});
```

- [ ] **Step 2: Run component tests to verify failure**

Run:

```bash
cd frontend && npm run test:run -- src/components
```

Expected: FAIL because component modules are undefined.

- [ ] **Step 3: Implement components**

Create `frontend/src/components/ConnectionStatus.tsx`:

```tsx
import type { ConnectionStatus as Status } from '../types/chat';

type Props = {
  status: Status;
  error: string | null;
};

const labels: Record<Status, string> = {
  idle: 'Aguardando entrada',
  connecting: 'Conectando',
  connected: 'Conectado',
  disconnected: 'Desconectado',
  error: 'Erro',
};

export default function ConnectionStatus({ status, error }: Props) {
  return (
    <div className="rounded-full bg-white/80 px-3 py-1 text-sm font-medium text-slate-700 shadow-sm">
      <span>{labels[status]}</span>
      {error ? <span className="ml-2 text-red-600">Erro: {error}</span> : null}
    </div>
  );
}
```

Create `frontend/src/components/JoinForm.tsx`:

```tsx
import { FormEvent, useState } from 'react';

type Props = {
  onJoin: (name: string) => void;
};

export default function JoinForm({ onJoin }: Props) {
  const [name, setName] = useState('');
  const trimmedName = name.trim();

  function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    if (trimmedName) {
      onJoin(trimmedName);
    }
  }

  return (
    <form onSubmit={handleSubmit} className="mx-auto flex max-w-md flex-col gap-4 rounded-3xl bg-white p-8 shadow-xl">
      <div>
        <h1 className="text-3xl font-bold text-slate-950">Realtime Chat</h1>
        <p className="mt-2 text-slate-600">Informe seu nome para entrar na sala geral.</p>
      </div>
      <label className="flex flex-col gap-2 text-sm font-medium text-slate-700">
        Seu nome
        <input
          className="rounded-2xl border border-slate-300 px-4 py-3 text-base outline-none focus:border-indigo-500 focus:ring-2 focus:ring-indigo-200"
          value={name}
          onChange={(event) => setName(event.target.value)}
          maxLength={40}
        />
      </label>
      <button
        className="rounded-2xl bg-indigo-600 px-5 py-3 font-semibold text-white transition hover:bg-indigo-500 disabled:cursor-not-allowed disabled:bg-slate-300"
        disabled={!trimmedName}
        type="submit"
      >
        Entrar no chat
      </button>
    </form>
  );
}
```

Create `frontend/src/components/MessageComposer.tsx`:

```tsx
import { FormEvent, useState } from 'react';

type Props = {
  disabled: boolean;
  onSend: (content: string) => void;
};

export default function MessageComposer({ disabled, onSend }: Props) {
  const [content, setContent] = useState('');
  const trimmedContent = content.trim();

  function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    if (disabled || !trimmedContent) {
      return;
    }
    onSend(trimmedContent);
    setContent('');
  }

  return (
    <form onSubmit={handleSubmit} className="flex gap-3 border-t border-slate-200 bg-white p-4">
      <label className="sr-only" htmlFor="message-input">
        Mensagem
      </label>
      <input
        id="message-input"
        className="min-w-0 flex-1 rounded-2xl border border-slate-300 px-4 py-3 outline-none focus:border-indigo-500 focus:ring-2 focus:ring-indigo-200"
        value={content}
        onChange={(event) => setContent(event.target.value)}
        disabled={disabled}
        maxLength={500}
      />
      <button
        className="rounded-2xl bg-indigo-600 px-5 py-3 font-semibold text-white transition hover:bg-indigo-500 disabled:cursor-not-allowed disabled:bg-slate-300"
        disabled={disabled || !trimmedContent}
        type="submit"
      >
        Enviar
      </button>
    </form>
  );
}
```

Create `frontend/src/components/MessageList.tsx`:

```tsx
import type { ChatMessage } from '../types/chat';

type Props = {
  messages: ChatMessage[];
};

export default function MessageList({ messages }: Props) {
  if (messages.length === 0) {
    return <div className="flex flex-1 items-center justify-center text-slate-500">Nenhuma mensagem ainda.</div>;
  }

  return (
    <ol className="flex flex-1 flex-col gap-3 overflow-y-auto p-4">
      {messages.map((message) => (
        <li key={message.id} className="rounded-2xl bg-white p-4 shadow-sm">
          <div className="flex items-center justify-between gap-3 text-sm text-slate-500">
            <strong className="text-slate-800">{message.userName || message.userId}</strong>
            <time dateTime={message.sentAt}>{formatTime(message.sentAt)}</time>
          </div>
          <p className="mt-2 whitespace-pre-wrap text-slate-900">{message.content}</p>
        </li>
      ))}
    </ol>
  );
}

function formatTime(value: string) {
  return new Intl.DateTimeFormat('pt-BR', {
    hour: '2-digit',
    minute: '2-digit',
  }).format(new Date(value));
}
```

Create `frontend/src/components/PeopleList.tsx`:

```tsx
import type { User } from '../types/chat';

type Props = {
  users: User[];
};

export default function PeopleList({ users }: Props) {
  return (
    <aside className="rounded-3xl bg-white p-5 shadow-sm lg:w-72">
      <h2 className="text-lg font-semibold text-slate-950">Pessoas online</h2>
      <p className="mt-1 text-sm text-slate-500">{users.length} conectado(s)</p>
      <ul className="mt-4 space-y-2">
        {users.map((user) => (
          <li key={user.id} className="flex items-center gap-3 rounded-2xl bg-slate-50 px-3 py-2 text-slate-800">
            <span className="h-2.5 w-2.5 rounded-full bg-emerald-500" aria-hidden="true" />
            <span>{user.name || user.id}</span>
          </li>
        ))}
      </ul>
    </aside>
  );
}
```

- [ ] **Step 4: Run component tests to verify pass**

Run:

```bash
cd frontend && npm run test:run -- src/components
```

Expected: PASS.

- [ ] **Step 5: Run all frontend tests**

Run:

```bash
cd frontend && npm run test:run
```

Expected: PASS.

- [ ] **Step 6: Checkpoint**

Record that required frontend unit-tested components are implemented.

---

## Task 9: Frontend WebSocket hook and app UI

**Files:**
- Create: `frontend/src/hooks/useChatConnection.ts`
- Modify: `frontend/src/App.tsx`
- Modify: `frontend/src/index.css`

- [ ] **Step 1: Invoke frontend design skill**

Before editing UI code, invoke the `frontend-design` skill and use it to keep the app visually polished while preserving the approved classic layout.

- [ ] **Step 2: Implement WebSocket hook**

Create `frontend/src/hooks/useChatConnection.ts`:

```ts
import { useCallback, useEffect, useReducer, useRef } from 'react';
import { chatReducer, initialChatState } from '../chat/chatReducer';
import type { ClientEvent, ServerEvent } from '../types/chat';

const reconnectDelayMs = 1200;

type Options = {
  url: string;
};

export function useChatConnection({ url }: Options) {
  const [state, dispatch] = useReducer(chatReducer, initialChatState);
  const socketRef = useRef<WebSocket | null>(null);
  const nameRef = useRef<string | null>(null);
  const reconnectTimerRef = useRef<number | null>(null);
  const manualCloseRef = useRef(false);

  const clearReconnectTimer = useCallback(() => {
    if (reconnectTimerRef.current !== null) {
      window.clearTimeout(reconnectTimerRef.current);
      reconnectTimerRef.current = null;
    }
  }, []);

  const sendClientEvent = useCallback((event: ClientEvent) => {
    const socket = socketRef.current;
    if (!socket || socket.readyState !== WebSocket.OPEN) {
      dispatch({ type: 'client-error', error: 'conexao indisponivel' });
      return;
    }
    socket.send(JSON.stringify(event));
  }, []);

  const connect = useCallback(
    (name: string) => {
      clearReconnectTimer();
      manualCloseRef.current = false;
      nameRef.current = name;
      dispatch({ type: 'status', status: 'connecting' });

      const socket = new WebSocket(url);
      socketRef.current = socket;

      socket.onopen = () => {
        dispatch({ type: 'status', status: 'connected' });
        socket.send(JSON.stringify({ type: 'join', name } satisfies ClientEvent));
      };

      socket.onmessage = (event) => {
        try {
          const serverEvent = JSON.parse(event.data) as ServerEvent;
          dispatch({ type: 'server-event', event: serverEvent });
        } catch {
          dispatch({ type: 'client-error', error: 'evento invalido recebido do servidor' });
        }
      };

      socket.onerror = () => {
        dispatch({ type: 'status', status: 'error' });
        dispatch({ type: 'client-error', error: 'falha na conexao' });
      };

      socket.onclose = () => {
        socketRef.current = null;
        dispatch({ type: 'status', status: 'disconnected' });
        if (!manualCloseRef.current && nameRef.current) {
          reconnectTimerRef.current = window.setTimeout(() => connect(nameRef.current!), reconnectDelayMs);
        }
      };
    },
    [clearReconnectTimer, url],
  );

  const disconnect = useCallback(() => {
    manualCloseRef.current = true;
    clearReconnectTimer();
    socketRef.current?.close();
    socketRef.current = null;
    dispatch({ type: 'status', status: 'disconnected' });
  }, [clearReconnectTimer]);

  const sendMessage = useCallback(
    (content: string) => {
      sendClientEvent({ type: 'message', content });
    },
    [sendClientEvent],
  );

  useEffect(() => {
    return () => {
      manualCloseRef.current = true;
      clearReconnectTimer();
      socketRef.current?.close();
    };
  }, [clearReconnectTimer]);

  return {
    state,
    connect,
    disconnect,
    sendMessage,
  };
}
```

- [ ] **Step 3: Wire app to hook and components**

Modify `frontend/src/App.tsx`:

```tsx
import { useState } from 'react';
import { initialChatState } from './chat/chatReducer';
import ConnectionStatus from './components/ConnectionStatus';
import JoinForm from './components/JoinForm';
import MessageComposer from './components/MessageComposer';
import MessageList from './components/MessageList';
import PeopleList from './components/PeopleList';
import { useChatConnection } from './hooks/useChatConnection';

const wsUrl = import.meta.env.VITE_WS_URL ?? 'ws://localhost:8080/ws';

export default function App() {
  const [joined, setJoined] = useState(false);
  const { state, connect, sendMessage } = useChatConnection({ url: wsUrl });
  const chatState = joined ? state : initialChatState;

  function handleJoin(name: string) {
    setJoined(true);
    connect(name);
  }

  if (!joined) {
    return (
      <main className="flex min-h-screen items-center justify-center bg-gradient-to-br from-indigo-100 via-slate-100 to-sky-100 p-6">
        <JoinForm onJoin={handleJoin} />
      </main>
    );
  }

  const canSend = chatState.status === 'connected';

  return (
    <main className="min-h-screen bg-gradient-to-br from-indigo-100 via-slate-100 to-sky-100 p-4 text-slate-950 md:p-8">
      <div className="mx-auto flex max-w-6xl flex-col gap-4">
        <header className="flex flex-col gap-3 rounded-3xl bg-white/80 p-5 shadow-sm backdrop-blur md:flex-row md:items-center md:justify-between">
          <div>
            <p className="text-sm font-medium uppercase tracking-wide text-indigo-600">Sala geral</p>
            <h1 className="text-3xl font-bold">Realtime Chat</h1>
          </div>
          <ConnectionStatus status={chatState.status} error={chatState.error} />
        </header>

        <section className="grid min-h-[70vh] gap-4 lg:grid-cols-[1fr_18rem]">
          <div className="flex min-h-[34rem] flex-col overflow-hidden rounded-3xl bg-slate-50 shadow-xl">
            <MessageList messages={chatState.messages} />
            <MessageComposer disabled={!canSend} onSend={sendMessage} />
          </div>
          <PeopleList users={chatState.users} />
        </section>
      </div>
    </main>
  );
}
```

- [ ] **Step 4: Run frontend tests and build**

Run:

```bash
cd frontend && npm run test:run && npm run build
```

Expected: both commands exit with status 0.

- [ ] **Step 5: Checkpoint**

Record that frontend hook, app composition, classic layout, tests, and build are complete.

---

## Task 10: Full local integration with Docker Compose

**Files:**
- Modify only if validation exposes a direct config mismatch: `docker-compose.yml`, `backend/Dockerfile`, `frontend/Dockerfile`, `.env.example`

- [ ] **Step 1: Run complete backend validation**

Run:

```bash
cd backend && go test ./... && go build ./cmd/chat-server
```

Expected: PASS and build exits with status 0.

- [ ] **Step 2: Run complete frontend validation**

Run:

```bash
cd frontend && npm run test:run && npm run build
```

Expected: PASS and build exits with status 0.

- [ ] **Step 3: Validate Compose config**

Run:

```bash
docker compose config
```

Expected: exits with status 0.

- [ ] **Step 4: Start services**

Run:

```bash
docker compose up --build
```

Expected: backend logs show the server listening on `:8080`; frontend logs show Vite serving on `http://localhost:5173/`.

- [ ] **Step 5: Verify backend health from a second terminal**

Run:

```bash
curl -i http://localhost:8080/healthz
```

Expected:

```text
HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8

ok
```

- [ ] **Step 6: Verify browser chat flow**

Open `http://localhost:5173` in two browser windows. In the first window, join as `Alice`; in the second, join as `Bob`. Send `hello` from Alice.

Expected:

- Both windows show Alice and Bob in the people list.
- Both windows show Alice's `hello` message with author and time.
- Closing Bob's window removes Bob from Alice's people list.

- [ ] **Step 7: Stop services**

Run:

```bash
docker compose down
```

Expected: services stop and containers are removed.

- [ ] **Step 8: Checkpoint**

Record that backend, frontend, and Docker Compose validations all pass and the browser chat flow works.

---

## Task 11: Documentation and wiki ingest

**Files:**
- Modify: `README.md`
- Modify: `docs/wiki/project-overview.md`
- Modify: `docs/wiki/architecture.md`
- Modify: `docs/wiki/backend.md`
- Modify: `docs/wiki/frontend.md`
- Modify: `docs/wiki/realtime-protocol.md`
- Modify: `docs/wiki/docker.md`
- Modify: `docs/wiki/testing.md`
- Modify: `docs/wiki/index.md`
- Modify: `docs/wiki/log.md`

- [ ] **Step 1: Update README with local commands**

Modify `README.md` so it includes this section after the current stack section:

````markdown
## Como rodar localmente

```bash
docker compose up --build
```

- Frontend: http://localhost:5173
- Backend healthcheck: http://localhost:8080/healthz
- WebSocket: ws://localhost:8080/ws

## Testes e build

```bash
cd backend && go test ./...
cd frontend && npm run test:run
cd frontend && npm run build
```
````

- [ ] **Step 2: Update wiki implementation status**

Apply these wiki updates:

- `project-overview.md`: change `Status` to "Implementação inicial concluída com backend, frontend, Docker Compose e testes unitários." Add current notes for `/healthz`, `/ws`, frontend on `5173`, backend on `8080`, and validation commands.
- `architecture.md`: record the final layout `backend/`, `frontend/`, `docker-compose.yml`; describe runtime flow as browser -> Vite frontend -> WebSocket `/ws` -> Go hub -> broadcast to connected clients.
- `backend.md`: list implemented files `cmd/chat-server/main.go`, `internal/protocol/events.go`, `internal/chat/hub.go`, `internal/ws/handler.go`, and `internal/server/router.go`; record `go test ./...`.
- `frontend.md`: list implemented files `src/App.tsx`, `src/hooks/useChatConnection.ts`, `src/chat/chatReducer.ts`, `src/components/*`, and Tailwind config; record `npm run test:run` and `npm run build`.
- `realtime-protocol.md`: record JSON shapes for `join`, `message`, `presence`, and `error` using the same field names from `protocol/events.go` and `src/types/chat.ts`.
- `docker.md`: record `docker compose up --build`, backend port `8080`, frontend port `5173`, Air backend reload, Vite HMR, and `VITE_WS_URL=ws://localhost:8080/ws`.
- `testing.md`: record exact commands `cd backend && go test ./...`, `cd frontend && npm run test:run`, `cd frontend && npm run build`, `docker compose config`, and the manual two-browser smoke test.
- `index.md`: set page summaries to implemented-state summaries for architecture, backend, frontend, realtime protocol, Docker, and testing.

- [ ] **Step 3: Append wiki log entry**

Append to `docs/wiki/log.md`:

```markdown
## [2026-06-09] implementação | Primeira versão funcional do chat

- Implementado backend Go com `net/http`, `nhooyr.io/websocket`, hub em memória, presença, broadcast de mensagens e endpoint `/healthz`.
- Implementado frontend React/Vite/Tailwind com layout clássico, entrada por nome, lista de mensagens, composer, status de conexão e lista de pessoas online.
- Adicionados testes Go para protocolo, hub, WebSocket e servidor; adicionados testes unitários frontend para reducer e componentes.
- Configurado Docker Compose para desenvolvimento local com backend em `8080`, frontend em `5173`, Air no backend e Vite HMR no frontend.
- Validações executadas: `cd backend && go test ./...`, `cd frontend && npm run test:run`, `cd frontend && npm run build`, `docker compose config`, `docker compose up --build` e fluxo manual no browser.
```

- [ ] **Step 4: Validate documentation files and links**

Run:

```bash
test -f README.md \
  && test -f docs/wiki/index.md \
  && test -f docs/wiki/log.md \
  && test -f docs/wiki/project-overview.md \
  && test -f docs/wiki/architecture.md \
  && test -f docs/wiki/backend.md \
  && test -f docs/wiki/frontend.md \
  && test -f docs/wiki/realtime-protocol.md \
  && test -f docs/wiki/docker.md \
  && test -f docs/wiki/testing.md
```

Expected: exits with status 0.

- [ ] **Step 5: Run final validation**

Run:

```bash
cd backend && go test ./...
cd ../frontend && npm run test:run && npm run build
cd .. && docker compose config
```

Expected: all commands exit with status 0.

- [ ] **Step 6: Checkpoint**

Record that implementation, tests, Docker validation, README, and wiki ingest are complete. Do not create a Git commit in this non-Git directory.
