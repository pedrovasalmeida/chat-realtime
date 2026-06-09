# Realtime Chat

Chat em tempo real com backend em Go, WebSocket, frontend React/Vite/Tailwind e ambiente local via Docker Compose.

## Status

Primeira versão funcional implementada. O projeto já possui backend, frontend, Docker Compose, testes unitários e wiki técnica em `docs/wiki/`.

## Funcionalidades

- Sala global única, sem autenticação e sem persistência após restart.
- Entrada por nome do usuário.
- Presença em tempo real com lista de pessoas online.
- Broadcast de mensagens para todos os clientes conectados.
- Tratamento de erros simples para entrada/mensagem inválida.
- Interface web com status de conexão, lista de mensagens, composer e lista de pessoas.

## Stack

| Camada | Tecnologia |
| --- | --- |
| Backend | Go, `net/http`, `nhooyr.io/websocket` |
| Frontend | React, Vite, TypeScript, Tailwind CSS |
| Testes | Go test, Vitest, Testing Library |
| Infra local | Docker Compose, Air, Vite HMR |

## Estrutura

| Caminho | Responsabilidade |
| --- | --- |
| `backend/` | Servidor Go, protocolo JSON, hub em memória, handler WebSocket e testes. |
| `frontend/` | App React/Vite/Tailwind, hook WebSocket, reducer, componentes e testes. |
| `docker-compose.yml` | Orquestra backend e frontend em desenvolvimento local. |
| `.env.example` | Variáveis locais esperadas para portas e URL do WebSocket. |
| `docs/wiki/` | Wiki técnica do projeto e histórico append-only. |
| `docs/raw/` | Fontes brutas de contexto; não substituir pela wiki gerada. |

## Como rodar localmente

```bash
docker compose up --build
```

Endpoints padrão:

- Frontend: http://localhost:5173
- Backend healthcheck: http://localhost:8080/healthz
- WebSocket: `ws://localhost:8080/ws`

Para encerrar:

```bash
docker compose down
```

Se as portas padrão estiverem ocupadas, use overrides:

```bash
BACKEND_PORT=18080 FRONTEND_PORT=15173 VITE_WS_URL=ws://localhost:18080/ws docker compose up --build
```

Com esses overrides:

- Frontend: http://localhost:15173
- Backend healthcheck: http://localhost:18080/healthz
- WebSocket: `ws://localhost:18080/ws`

## Testes e build

Backend:

```bash
cd backend
go test ./...
go build ./cmd/chat-server
```

Frontend:

```bash
cd frontend
npm run test:run
npm run build
```

Docker Compose:

```bash
docker compose config
```

## Fluxo manual de smoke test

1. Suba o ambiente.
2. Abra o frontend em duas janelas.
3. Entre como `Alice` em uma janela e `Bob` na outra.
4. Envie `hello` como Alice.
5. Confirme que ambas as janelas exibem Alice e Bob online e a mensagem enviada.
6. Feche a janela do Bob e confirme que Bob sai da lista de pessoas online da Alice.

## Wiki e manutenção

1. Leia `.github/copilot-instructions.md` antes de fazer mudanças.
2. Consulte `docs/wiki/index.md` para localizar páginas técnicas relevantes.
3. Ao concluir mudanças relevantes, atualize a página específica da wiki e registre a entrada em `docs/wiki/log.md`.
4. Mantenha `docs/wiki/log.md` append-only.
