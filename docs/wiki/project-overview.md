# Project Overview

## Status

Implementação inicial concluída com backend, frontend, Docker Compose e testes unitários.

## Purpose

Registrar a visão do Realtime Chat e manter o contexto de alto nível sincronizado com a evolução do repo.

## Current notes

- Backend Go expõe `GET /healthz` e WebSocket em `/ws`.
- Frontend React/Vite/Tailwind roda em `5173` e usa `VITE_WS_URL` para conectar ao backend.
- Backend roda em `8080` por padrão; portas podem ser sobrescritas via `BACKEND_PORT`, `FRONTEND_PORT` e `VITE_WS_URL`.
- Interface implementada: layout clássico com mensagens, campo de envio, status de conexão e lista de pessoas no chat à direita.
- Primeira versão: sala global única, usuários identificados por nome, sem autenticação e sem persistência após restart.
- Validações principais: `cd backend && go test ./...`, `cd frontend && npm run test:run`, `cd frontend && npm run build` e `docker compose config`.

## Decisions pending

- Persistência de mensagens e histórico.
- Autenticação e salas múltiplas.
- Testes e2e automatizados em browser.

## Related pages

- [Arquitetura](architecture.md)
- [Backend](backend.md)
- [Frontend](frontend.md)
- [Protocolo realtime](realtime-protocol.md)
- [Docker](docker.md)
- [Testes](testing.md)
