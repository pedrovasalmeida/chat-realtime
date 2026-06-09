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
