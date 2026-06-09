# Realtime Chat Repo Preparation Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Create the documentation-only repository foundation for a Go + React realtime chat project.

**Architecture:** Bootstrap three documentation layers: `README.md` for human onboarding, `.github/copilot-instructions.md` for agent workflow rules, and `docs/wiki/` for the persistent project wiki. Do not scaffold backend, frontend, Docker, or Git in this plan.

**Tech Stack:** Markdown, GitHub repository metadata, shell validation.

---

## Scope Check

This plan implements one focused subsystem: repository documentation and wiki preparation. It does not implement Go backend code, React frontend code, Docker files, package manifests, or runtime configuration.

## Version Control Policy

The current directory is not a Git repository and the approved spec excludes Git initialization or commits. This plan has no commit steps. If the repository is initialized in a future task, commit messages must not mention IA, Copilot, assistants, or automated co-authorship unless the user explicitly requests it.

## File Structure

- Create `README.md`: project entry point, stack, current status, repo navigation, and next steps.
- Create `.github/copilot-instructions.md`: repository rules for agents, wiki workflow, language, validation, and commit-message constraints.
- Create `docs/wiki/index.md`: content-oriented wiki catalog.
- Create `docs/wiki/log.md`: append-only chronological wiki history.
- Create `docs/wiki/project-overview.md`: high-level project status and stack template.
- Create `docs/wiki/architecture.md`: architecture decision template.
- Create `docs/wiki/backend.md`: Go backend template.
- Create `docs/wiki/frontend.md`: React/Vite/Tailwind frontend template.
- Create `docs/wiki/realtime-protocol.md`: realtime events and payloads template.
- Create `docs/wiki/docker.md`: container and local environment template.
- Create `docs/wiki/testing.md`: testing strategy template.
- Keep `docs/raw/idea.md` unchanged as a raw source.
- Keep `docs/wiki/llm-wiki.md` unchanged as the wiki-pattern reference.

---

### Task 1: Repository Entrypoint and Agent Instructions

**Files:**
- Create: `README.md`
- Create: `.github/copilot-instructions.md`

- [ ] **Step 1: Confirm target files do not already exist**

Run:

```bash
test ! -e README.md && test ! -e .github/copilot-instructions.md
```

Expected: no output and exit code 0. If this fails, inspect the existing file before editing so user content is not overwritten.

- [ ] **Step 2: Create the GitHub metadata directory**

Run:

```bash
mkdir -p .github
```

Expected: no output and exit code 0.

- [ ] **Step 3: Create `README.md`**

Create `README.md` with this exact content:

```markdown
# Realtime Chat

Projeto de chat em tempo real construído com Go no backend e React no frontend.

## Status

Preparação inicial do repositório. Ainda não há scaffold de backend, frontend ou Docker.

## Stack prevista

| Camada | Tecnologia |
| --- | --- |
| Frontend | React, Vite, Tailwind CSS e testes |
| Backend | Go e testes |
| Infra local | Docker |

## Estrutura atual

| Caminho | Responsabilidade |
| --- | --- |
| `.github/copilot-instructions.md` | Regras de trabalho para agentes e manutenção da wiki. |
| `docs/raw/idea.md` | Fonte bruta com a ideia inicial do projeto. |
| `docs/wiki/llm-wiki.md` | Referência do padrão de wiki mantida por LLM. |
| `docs/wiki/index.md` | Catálogo navegável da wiki do projeto. |
| `docs/wiki/log.md` | Histórico cronológico de mudanças relevantes. |

## Como trabalhar neste repo

1. Leia `.github/copilot-instructions.md` antes de fazer mudanças.
2. Consulte `docs/wiki/index.md` para encontrar páginas relevantes.
3. Ao concluir uma implementação futura, atualize a wiki com o que mudou, como validar e quais arquivos foram afetados.
4. Registre mudanças relevantes em `docs/wiki/log.md`.

## Próximos passos

1. Planejar o backend em Go.
2. Planejar o frontend em React, Vite e Tailwind.
3. Definir a estratégia de Docker para ambiente local.
```

- [ ] **Step 4: Create `.github/copilot-instructions.md`**

Create `.github/copilot-instructions.md` with this exact content:

```markdown
# Instruções do repositório

## Contexto

Este repositório contém um projeto de chat em tempo real com backend em Go, frontend em React/Vite/Tailwind, testes e Docker.

## Regras gerais

- Trabalhe em português do Brasil, exceto quando o usuário pedir outro idioma.
- Faça mudanças pequenas e alinhadas ao pedido atual.
- Não implemente backend, frontend ou Docker quando a tarefa for apenas documental.
- Preserve `docs/raw/` como fonte bruta; não trate essa pasta como wiki gerada.
- Não crie commits com menções a IA, Copilot, assistentes ou coautoria automatizada, exceto quando o usuário pedir explicitamente.
- Quando o diretório não for um repositório Git, não inicialize Git sem aprovação do usuário.

## Fluxo da wiki

1. Leia `docs/wiki/index.md` antes de modificar ou usar páginas da wiki.
2. Atualize `docs/wiki/index.md` quando criar, renomear ou alterar materialmente páginas da wiki.
3. Acrescente entradas em `docs/wiki/log.md` após mudanças documentais relevantes, implementações concluídas ou ingestões.
4. Após implementar planos futuros, faça a ingestão na wiki registrando o que mudou, onde validar, arquivos relevantes e páginas atualizadas.
5. Mantenha `docs/wiki/log.md` append-only.

## Convenções de páginas da wiki

Use estas seções quando fizer sentido:

- `Status`
- `Purpose`
- `Current notes`
- `Decisions pending`
- `Related pages`

As páginas iniciais devem permanecer como templates mínimos até que decisões sejam aprovadas.

## Validação

- Para mudanças documentais, confirme que links e arquivos esperados existem.
- Para mudanças de código futuras, rode os testes e comandos de build já definidos no repo.
- Não adicione ferramentas de lint, build ou teste sem necessidade aprovada.
```

- [ ] **Step 5: Validate Task 1**

Run:

```bash
test -f README.md && test -f .github/copilot-instructions.md
grep -F "docs/wiki/index.md" README.md
grep -F "Não crie commits com menções a IA, Copilot" .github/copilot-instructions.md
```

Expected: both files exist, and the two `grep` commands print matching lines.

---

### Task 2: Wiki Catalog and Chronological Log

**Files:**
- Create: `docs/wiki/index.md`
- Create: `docs/wiki/log.md`

- [ ] **Step 1: Confirm wiki operation files do not already exist**

Run:

```bash
test ! -e docs/wiki/index.md && test ! -e docs/wiki/log.md
```

Expected: no output and exit code 0. If this fails, inspect the existing file before editing so user content is not overwritten.

- [ ] **Step 2: Create `docs/wiki/index.md`**

Create `docs/wiki/index.md` with this exact content:

```markdown
# Wiki Index

## Status

Wiki inicial do projeto Realtime Chat. As páginas abaixo são templates mínimos para registrar decisões futuras sem criar scaffold de aplicação.

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
| [Visão geral do projeto](project-overview.md) | Objetivo, stack e status atual do Realtime Chat. |
| [Arquitetura](architecture.md) | Espaço para decisões arquiteturais aprovadas. |
| [Backend](backend.md) | Espaço para decisões e notas do backend em Go. |
| [Frontend](frontend.md) | Espaço para decisões e notas do frontend em React/Vite/Tailwind. |
| [Protocolo realtime](realtime-protocol.md) | Espaço para eventos, payloads, presença e ciclo de conexão. |
| [Docker](docker.md) | Espaço para decisões de containers e ambiente local. |
| [Testes](testing.md) | Espaço para estratégia de testes backend, frontend e integração. |
| [Log](log.md) | Histórico cronológico append-only da wiki. |

## Manutenção

- `index.md` é o catálogo de navegação e deve ficar sincronizado com as páginas existentes.
- `log.md` é cronológico e não deve ser reescrito, apenas acrescido.
- Páginas em `docs/raw/` são fontes brutas e não substituem páginas da wiki.
```

- [ ] **Step 3: Create `docs/wiki/log.md`**

Create `docs/wiki/log.md` with this exact content:

```markdown
# Wiki Log

Este arquivo é append-only. Use entradas no formato `## [YYYY-MM-DD] tipo | título`.

## [2026-06-09] preparação | Estrutura inicial da wiki

- Registrada a preparação documental do repositório Realtime Chat.
- Definida a wiki inicial com `index.md`, `log.md` e páginas core em formato template mínimo.
- Mantidos `docs/raw/idea.md` e `docs/wiki/llm-wiki.md` como fontes de contexto.
- Escopo desta etapa: documentação base, sem scaffold Go, React ou Docker.
```

- [ ] **Step 4: Validate Task 2**

Run:

```bash
test -f docs/wiki/index.md && test -f docs/wiki/log.md
grep -F "[Log](log.md)" docs/wiki/index.md
grep -F "## [2026-06-09] preparação | Estrutura inicial da wiki" docs/wiki/log.md
```

Expected: both files exist, and the two `grep` commands print matching lines.

---

### Task 3: Project Overview and Architecture Wiki Pages

**Files:**
- Create: `docs/wiki/project-overview.md`
- Create: `docs/wiki/architecture.md`

- [ ] **Step 1: Confirm project-level wiki pages do not already exist**

Run:

```bash
test ! -e docs/wiki/project-overview.md && test ! -e docs/wiki/architecture.md
```

Expected: no output and exit code 0. If this fails, inspect the existing file before editing so user content is not overwritten.

- [ ] **Step 2: Create `docs/wiki/project-overview.md`**

Create `docs/wiki/project-overview.md` with this exact content:

```markdown
# Project Overview

## Status

Preparação inicial. O projeto ainda não possui scaffold de backend, frontend ou Docker.

## Purpose

Registrar a visão do Realtime Chat e manter o contexto de alto nível sincronizado com a evolução do repo.

## Current notes

- Objetivo: chat em tempo real.
- Backend previsto: Go com testes.
- Frontend previsto: React, Vite, Tailwind CSS e testes.
- Infra local prevista: Docker.
- Interface desejada: lista de mensagens, campo de envio, botão de envio e lista de pessoas no chat.

## Decisions pending

- Estrutura final do monorepo.
- Estratégia de transporte realtime.
- Modelo de mensagens e presença.
- Persistência de mensagens.
- Autenticação ou identificação de usuários.
- Comandos de teste e build.

## Related pages

- [Arquitetura](architecture.md)
- [Backend](backend.md)
- [Frontend](frontend.md)
- [Protocolo realtime](realtime-protocol.md)
- [Docker](docker.md)
- [Testes](testing.md)
```

- [ ] **Step 3: Create `docs/wiki/architecture.md`**

Create `docs/wiki/architecture.md` with this exact content:

```markdown
# Architecture

## Status

Template inicial. Nenhuma decisão arquitetural além da stack informada foi fechada.

## Purpose

Centralizar decisões sobre organização do sistema, limites entre frontend/backend e fluxo de dados.

## Current notes

- A aplicação será um chat em tempo real.
- O backend será escrito em Go.
- O frontend será escrito com React, Vite e Tailwind CSS.
- Docker será usado em etapa futura para ambiente local.

## Decisions pending

- Layout de pastas para backend, frontend e infraestrutura.
- Contrato entre frontend e backend.
- Estratégia de gerenciamento de conexão realtime.
- Estratégia de armazenamento e recuperação de mensagens.
- Separação entre configuração local, testes e produção.

## Related pages

- [Visão geral do projeto](project-overview.md)
- [Backend](backend.md)
- [Frontend](frontend.md)
- [Protocolo realtime](realtime-protocol.md)
- [Docker](docker.md)
```

- [ ] **Step 4: Validate Task 3**

Run:

```bash
test -f docs/wiki/project-overview.md && test -f docs/wiki/architecture.md
grep -F "[Arquitetura](architecture.md)" docs/wiki/project-overview.md
grep -F "[Visão geral do projeto](project-overview.md)" docs/wiki/architecture.md
```

Expected: both files exist, and the two `grep` commands print matching lines.

---

### Task 4: Core Technical Wiki Templates

**Files:**
- Create: `docs/wiki/backend.md`
- Create: `docs/wiki/frontend.md`
- Create: `docs/wiki/realtime-protocol.md`
- Create: `docs/wiki/docker.md`
- Create: `docs/wiki/testing.md`

- [ ] **Step 1: Confirm technical wiki pages do not already exist**

Run:

```bash
test ! -e docs/wiki/backend.md && test ! -e docs/wiki/frontend.md && test ! -e docs/wiki/realtime-protocol.md && test ! -e docs/wiki/docker.md && test ! -e docs/wiki/testing.md
```

Expected: no output and exit code 0. If this fails, inspect the existing file before editing so user content is not overwritten.

- [ ] **Step 2: Create `docs/wiki/backend.md`**

Create `docs/wiki/backend.md` with this exact content:

```markdown
# Backend

## Status

Template inicial. Nenhum módulo Go foi criado nesta etapa.

## Purpose

Registrar decisões e notas sobre o backend em Go.

## Current notes

- Backend previsto em Go.
- Testes são parte obrigatória da stack.
- O backend deverá expor recursos para envio, recebimento e distribuição de mensagens em tempo real.

## Decisions pending

- Versão do Go.
- Layout do módulo e pacotes.
- Biblioteca ou abordagem HTTP.
- Transporte realtime.
- Modelo de mensagem, sala e usuário.
- Estratégia de testes unitários e de integração.
- Estratégia de logs e tratamento de erros.

## Related pages

- [Arquitetura](architecture.md)
- [Protocolo realtime](realtime-protocol.md)
- [Testes](testing.md)
```

- [ ] **Step 3: Create `docs/wiki/frontend.md`**

Create `docs/wiki/frontend.md` with this exact content:

```markdown
# Frontend

## Status

Template inicial. Nenhum scaffold React/Vite foi criado nesta etapa.

## Purpose

Registrar decisões e notas sobre a interface web do chat.

## Current notes

- Frontend previsto com React, Vite e Tailwind CSS.
- Testes são parte obrigatória da stack.
- UI desejada: visualização das mensagens, campo de texto, botão de envio e lista de pessoas à direita.
- Cada mensagem deve exibir nome ou id do usuário e horário.

## Decisions pending

- Estrutura de pastas do frontend.
- Biblioteca de testes.
- Estratégia de estado da conexão realtime.
- Formato de componentes para chat, mensagens, input e lista de pessoas.
- Tratamento visual de conexão, reconexão e erros.

## Related pages

- [Arquitetura](architecture.md)
- [Protocolo realtime](realtime-protocol.md)
- [Testes](testing.md)
```

- [ ] **Step 4: Create `docs/wiki/realtime-protocol.md`**

Create `docs/wiki/realtime-protocol.md` with this exact content:

```markdown
# Realtime Protocol

## Status

Template inicial. O protocolo realtime ainda não foi definido.

## Purpose

Registrar eventos, payloads e estados de conexão usados entre frontend e backend.

## Current notes

- O projeto exige comunicação em tempo real para troca de mensagens.
- A lista de pessoas no chat depende de presença ou estado equivalente.
- Mensagens exibem usuário e horário.

## Decisions pending

- Transporte realtime.
- Eventos de entrada, saída, envio de mensagem, recebimento de mensagem e presença.
- Formato dos payloads.
- Identificação de usuário.
- Estratégia de reconexão.
- Tratamento de mensagens duplicadas ou fora de ordem.

## Related pages

- [Arquitetura](architecture.md)
- [Backend](backend.md)
- [Frontend](frontend.md)
- [Testes](testing.md)
```

- [ ] **Step 5: Create `docs/wiki/docker.md`**

Create `docs/wiki/docker.md` with this exact content:

```markdown
# Docker

## Status

Template inicial. Nenhum arquivo Docker foi criado nesta etapa.

## Purpose

Registrar decisões de containerização e ambiente local.

## Current notes

- Docker faz parte da stack prevista.
- A configuração será definida depois dos scaffolds de backend e frontend.

## Decisions pending

- Imagens para backend e frontend.
- Uso de Docker Compose.
- Estratégia de hot reload local.
- Variáveis de ambiente.
- Portas locais.
- Comandos de build e execução.

## Related pages

- [Arquitetura](architecture.md)
- [Backend](backend.md)
- [Frontend](frontend.md)
- [Testes](testing.md)
```

- [ ] **Step 6: Create `docs/wiki/testing.md`**

Create `docs/wiki/testing.md` with this exact content:

```markdown
# Testing

## Status

Template inicial. Nenhuma configuração de testes foi criada nesta etapa.

## Purpose

Registrar a estratégia de testes do projeto conforme backend, frontend e integração forem planejados.

## Current notes

- Backend em Go deve ter testes.
- Frontend em React/Vite/Tailwind deve ter testes.
- Validação documental desta etapa garante apenas estrutura e links básicos.

## Decisions pending

- Comandos de teste do backend.
- Framework e comandos de teste do frontend.
- Estratégia de testes de integração para comunicação realtime.
- Critérios mínimos para considerar uma implementação concluída.
- Como registrar validações no log da wiki.

## Related pages

- [Backend](backend.md)
- [Frontend](frontend.md)
- [Protocolo realtime](realtime-protocol.md)
- [Docker](docker.md)
```

- [ ] **Step 7: Validate Task 4**

Run:

```bash
test -f docs/wiki/backend.md && test -f docs/wiki/frontend.md && test -f docs/wiki/realtime-protocol.md && test -f docs/wiki/docker.md && test -f docs/wiki/testing.md
grep -F "[Protocolo realtime](realtime-protocol.md)" docs/wiki/backend.md
grep -F "Cada mensagem deve exibir nome ou id do usuário e horário." docs/wiki/frontend.md
grep -F "[Backend](backend.md)" docs/wiki/realtime-protocol.md
grep -F "Nenhum arquivo Docker foi criado nesta etapa." docs/wiki/docker.md
grep -F "Validação documental desta etapa garante apenas estrutura e links básicos." docs/wiki/testing.md
```

Expected: all files exist, and the five `grep` commands print matching lines.

---

### Task 5: Final Documentation Validation

**Files:**
- Validate: `README.md`
- Validate: `.github/copilot-instructions.md`
- Validate: `docs/wiki/*.md`
- Validate unchanged source context: `docs/raw/idea.md`, `docs/wiki/llm-wiki.md`

- [ ] **Step 1: Validate all required files exist**

Run:

```bash
for path in README.md .github/copilot-instructions.md docs/wiki/index.md docs/wiki/log.md docs/wiki/project-overview.md docs/wiki/architecture.md docs/wiki/backend.md docs/wiki/frontend.md docs/wiki/realtime-protocol.md docs/wiki/docker.md docs/wiki/testing.md docs/raw/idea.md docs/wiki/llm-wiki.md; do
  test -f "$path"
done
```

Expected: no output and exit code 0.

- [ ] **Step 2: Validate the wiki index catalogs every created page**

Run:

```bash
grep -F "[Visão geral do projeto](project-overview.md)" docs/wiki/index.md
grep -F "[Arquitetura](architecture.md)" docs/wiki/index.md
grep -F "[Backend](backend.md)" docs/wiki/index.md
grep -F "[Frontend](frontend.md)" docs/wiki/index.md
grep -F "[Protocolo realtime](realtime-protocol.md)" docs/wiki/index.md
grep -F "[Docker](docker.md)" docs/wiki/index.md
grep -F "[Testes](testing.md)" docs/wiki/index.md
grep -F "[Log](log.md)" docs/wiki/index.md
```

Expected: all eight `grep` commands print matching lines.

- [ ] **Step 3: Validate README and agent instructions reference the wiki workflow**

Run:

```bash
grep -F "Preparação inicial do repositório." README.md
grep -F "Ao concluir uma implementação futura, atualize a wiki" README.md
grep -F "Após implementar planos futuros" .github/copilot-instructions.md
grep -F "Mantenha \`docs/wiki/log.md\` append-only." .github/copilot-instructions.md
```

Expected: all four `grep` commands print matching lines.

- [ ] **Step 4: Validate no application scaffold was created**

Run:

```bash
find . -maxdepth 2 \( -name package.json -o -name go.mod -o -name Dockerfile -o -name docker-compose.yml -o -name docker-compose.yaml \) -print
```

Expected: no output.

- [ ] **Step 5: Validate Git was not initialized by this plan**

Run:

```bash
if git rev-parse --show-toplevel >/dev/null 2>&1; then echo "git-present"; else echo "git-absent"; fi
```

Expected: `git-absent`.

- [ ] **Step 6: Report completion**

Report these facts:

```text
Created README.md, .github/copilot-instructions.md, docs/wiki/index.md, docs/wiki/log.md, and the core wiki template pages.
Kept docs/raw/idea.md and docs/wiki/llm-wiki.md unchanged.
Did not create Go, React, Docker, package, or Git scaffold files.
Did not create commits.
```

