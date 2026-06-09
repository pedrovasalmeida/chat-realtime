# Wiki Index

## Status

Wiki do projeto Realtime Chat atualizada após a primeira implementação funcional com backend, frontend, Docker Compose e testes.

## Como usar

1. Leia este índice para localizar a página relevante.
2. Atualize a página específica com decisões aprovadas ou resultados de implementação.
3. Atualize este índice quando páginas forem criadas, renomeadas ou materialmente alteradas.
4. Registre mudanças relevantes em [Log](log.md).

## Fontes

| Página | Resumo |
| --- | --- |
| [LLM Wiki](llm-wiki.md) | Referência do padrão usado para manter esta wiki. |
| [Ideia inicial](../raw/idea.md) | Fonte bruta com escopo e stack informados no início do projeto. |

## Páginas do projeto

| Página | Resumo |
| --- | --- |
| [Visão geral do projeto](project-overview.md) | Status implementado, endpoints, portas e validações principais do Realtime Chat. |
| [Arquitetura](architecture.md) | Layout final da primeira versão e fluxo browser -> Vite -> WebSocket -> hub Go. |
| [Backend](backend.md) | Módulos Go implementados para protocolo, hub, WebSocket, roteador e servidor. |
| [Frontend](frontend.md) | App React/Vite/Tailwind implementado com hook WebSocket, reducer e componentes testados. |
| [Protocolo realtime](realtime-protocol.md) | Formatos JSON implementados para `join`, `message`, `presence` e `error`. |
| [Docker](docker.md) | Docker Compose de desenvolvimento com Air, Vite HMR, portas e overrides locais. |
| [Testes](testing.md) | Comandos reais de teste/build e smoke de integração da primeira versão. |
| [Log](log.md) | Histórico cronológico append-only da wiki. |

## Manutenção

- `index.md` é o catálogo de navegação e deve ficar sincronizado com as páginas existentes.
- `log.md` é cronológico e não deve ser reescrito, apenas acrescido.
- Páginas em `docs/raw/` são fontes brutas e não substituem páginas da wiki.
