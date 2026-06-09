# Realtime Chat

Projeto de chat em tempo real construído com Go no backend e React no frontend.

## Status

Preparação inicial do repositório. Ainda não há scaffold de backend, frontend ou Docker.

## Stack prevista

| Camada      | Tecnologia                         |
| ----------- | ---------------------------------- |
| Frontend    | React, Vite, Tailwind CSS e testes |
| Backend     | Go e testes                        |
| Infra local | Docker                             |

## Como rodar localmente

```bash
docker compose up --build
```

- Frontend: http://localhost:5173
- Backend healthcheck: http://localhost:8080/healthz
- WebSocket: ws://localhost:8080/ws

Se as portas padrão estiverem ocupadas, use overrides:

```bash
BACKEND_PORT=18080 FRONTEND_PORT=15173 VITE_WS_URL=ws://localhost:18080/ws docker compose up --build
```

## Testes e build

```bash
cd backend && go test ./...
cd frontend && npm run test:run
cd frontend && npm run build
```

## Estrutura atual

| Caminho                           | Responsabilidade                                      |
| --------------------------------- | ----------------------------------------------------- |
| `.github/copilot-instructions.md` | Regras de trabalho para agentes e manutenção da wiki. |
| `docs/raw/idea.md`                | Fonte bruta com a ideia inicial do projeto.           |
| `docs/wiki/llm-wiki.md`           | Referência do padrão de wiki mantida por LLM.         |
| `docs/wiki/index.md`              | Catálogo navegável da wiki do projeto.                |
| `docs/wiki/log.md`                | Histórico cronológico de mudanças relevantes.         |

## Como trabalhar neste repo

1. Leia `.github/copilot-instructions.md` antes de fazer mudanças.
2. Consulte `docs/wiki/index.md` para encontrar páginas relevantes.
3. Ao concluir uma implementação futura, atualize a wiki com o que mudou, como validar e quais arquivos foram afetados.
4. Registre mudanças relevantes em `docs/wiki/log.md`.
