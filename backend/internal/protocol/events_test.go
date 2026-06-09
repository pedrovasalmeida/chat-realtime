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
