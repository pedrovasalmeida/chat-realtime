# Architecture

## Status

Primeira implementação funcional criada com monorepo, backend Go, frontend React/Vite e Docker Compose.

## Purpose

Centralizar decisões sobre organização do sistema, limites entre frontend/backend e fluxo de dados.

## Current notes

- Estrutura final da primeira versão: `backend/`, `frontend/`, `docker-compose.yml`, `.env.example` e Dockerfiles de desenvolvimento.
- O backend em Go mantém conexões, presença e mensagens em memória na primeira versão.
- O frontend em React/Vite/Tailwind renderiza entrada por nome, mensagens, composer, status de conexão e pessoas online.
- Fluxo de runtime: browser -> Vite frontend -> WebSocket `/ws` -> Go hub -> broadcast para clientes conectados.
- `docker-compose.yml` orquestra backend com Air e frontend com Vite HMR.

## Decisions pending

- Estratégia de deploy/produção.
- Persistência e recuperação de histórico.
- Separação de salas e autenticação.

## Related pages

- [Visão geral do projeto](project-overview.md)
- [Backend](backend.md)
- [Frontend](frontend.md)
- [Protocolo realtime](realtime-protocol.md)
- [Docker](docker.md)
