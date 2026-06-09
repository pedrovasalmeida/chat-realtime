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
