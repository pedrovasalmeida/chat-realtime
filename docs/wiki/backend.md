# Backend

## Status

Backend Go implementado com protocolo JSON, hub em memória, handler WebSocket, roteador HTTP e testes.

## Purpose

Registrar decisões e notas sobre o backend em Go.

## Current notes

- Módulo Go: `backend/go.mod`.
- Entrypoint: `cmd/chat-server/main.go`.
- Protocolo e validação: `internal/protocol/events.go`.
- Hub em memória: `internal/chat/hub.go`, com join, leave, presença, broadcast e validação de mensagem.
- WebSocket: `internal/ws/handler.go`, com primeiro evento obrigatório `join`, mensagens, erros e cleanup de presença.
- HTTP: `internal/server/router.go`, com `GET /healthz` e `GET /ws`.
- Testes: `internal/protocol/events_test.go`, `internal/chat/hub_test.go`, `internal/ws/handler_test.go` e `internal/server/router_test.go`.
- Comando principal: `cd backend && go test ./...`.

## Decisions pending

- Persistência e recuperação de mensagens.
- Autenticação/autorização de usuários.
- Observabilidade e logs estruturados.

## Related pages

- [Arquitetura](architecture.md)
- [Protocolo realtime](realtime-protocol.md)
- [Testes](testing.md)
