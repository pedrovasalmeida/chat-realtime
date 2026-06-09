# Docker

## Status

Docker Compose implementado para desenvolvimento local com backend em Go/Air e frontend em Vite.

## Purpose

Registrar decisões de containerização e ambiente local.

## Current notes

- Comando padrão:

```bash
docker compose up --build
```

- Backend expõe `8080` por padrão e usa Air para hot reload.
- Frontend expõe `5173` por padrão e usa Vite HMR.
- O frontend recebe `VITE_WS_URL=ws://localhost:8080/ws` por padrão.
- Portas podem ser sobrescritas quando houver conflito local:

```bash
BACKEND_PORT=18080 FRONTEND_PORT=15173 VITE_WS_URL=ws://localhost:18080/ws docker compose up --build
```

- Validação de sintaxe:

```bash
docker compose config
```

## Decisions pending

- Separar configuração de produção.
- Healthchecks de containers.

## Related pages

- [Arquitetura](architecture.md)
- [Backend](backend.md)
- [Frontend](frontend.md)
- [Testes](testing.md)
