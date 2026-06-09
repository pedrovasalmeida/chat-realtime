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
