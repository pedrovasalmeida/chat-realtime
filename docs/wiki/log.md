# Wiki Log

Este arquivo é append-only. Use entradas no formato `## [YYYY-MM-DD] tipo | título`.

## [2026-06-09] preparação | Estrutura inicial da wiki

- Registrada a preparação documental do repositório Realtime Chat.
- Definida a wiki inicial com `index.md`, `log.md` e páginas core em formato template mínimo.
- Mantidos `docs/raw/idea.md` e `docs/wiki/llm-wiki.md` como fontes de contexto.
- Escopo desta etapa: documentação base, sem scaffold Go, React ou Docker.

## [2026-06-09] planejamento | Design inicial de implementação do chat

- Criada a spec principal em `docs/superpowers/specs/2026-06-09-realtime-chat-implementation-design.md`.
- Aprovada a abordagem de implementação em etapas: Docker/dev, backend Go/WebSocket, frontend React/Vite/Tailwind, integração e ingestão na wiki.
- Registradas decisões iniciais: sala global única, usuários por nome sem autenticação, mensagens em memória, WebSocket com `nhooyr.io/websocket`, layout clássico do chat e Docker Compose para desenvolvimento.
- Ajustada a estratégia de testes: backend com testes Go obrigatórios, frontend com testes unitários e sem testes e2e inicialmente.
- Atualizadas páginas da wiki de visão geral, arquitetura, backend, frontend, protocolo realtime, Docker e testes.

## [2026-06-09] planejamento | Plano de implementação do chat

- Criado o plano em `docs/superpowers/plans/2026-06-09-realtime-chat-implementation.md`.
- O plano detalha execução por tarefas para Docker/dev, backend Go/WebSocket, frontend React/Vite/Tailwind, integração local e ingestão na wiki.
- Registrado que o diretório ainda não é um repositório Git; o plano usa checkpoints em vez de commits.

## [2026-06-09] documentação | Ingestão da wiki como último step

- Ajustada a regra em `.github/copilot-instructions.md` para que a ingestão da wiki aconteça apenas no último step dos planos de implementação.
- Registrado que a ingestão deve ocorrer somente quando o plano estiver devidamente implementado.

## [2026-06-09] implementação | Primeira versão funcional do chat

- Implementado backend Go com `net/http`, `nhooyr.io/websocket`, hub em memória, presença, broadcast de mensagens e endpoint `/healthz`.
- Implementado frontend React/Vite/Tailwind com layout clássico refinado, entrada por nome, lista de mensagens, composer, status de conexão e lista de pessoas online.
- Adicionados testes Go para protocolo, hub, WebSocket e servidor; adicionados testes unitários frontend para reducer e componentes.
- Configurado Docker Compose para desenvolvimento local com backend em `8080`, frontend em `5173`, Air no backend, Vite HMR no frontend e overrides de porta via variáveis de ambiente.
- Validações executadas: `cd backend && go test ./...`, build backend com `GOFLAGS=-buildvcs=false`, `cd frontend && npm run test:run`, `cd frontend && npm run build`, `docker compose config`, `docker compose up --build` com portas alternativas por conflito local em `8080`, healthcheck e smoke com dois clientes WebSocket.
