# Realtime Chat Repo Preparation Design

## Status

Approved on 2026-06-09.

## Goal

Prepare the current repository for a realtime chat project built with Go, React, Vite, Tailwind, tests, and Docker. This preparation is documentation-only: it creates the base project documentation and an LLM-maintained wiki structure, without scaffolding application code yet.

## Scope

In scope:

- Create `README.md` as the project entry point.
- Create `.github/copilot-instructions.md` as the operating guide for Copilot/LLM agents.
- Instantiate the `docs/wiki/llm-wiki.md` pattern with `index.md`, `log.md`, and core template pages.
- Keep existing source/context files under `docs/raw/` and `docs/wiki/llm-wiki.md`.
- Record future implementation work in the wiki after plans are implemented.

Out of scope:

- Go backend scaffold.
- React frontend scaffold.
- Docker configuration.
- Concrete architecture decisions beyond the stack already provided.
- Git initialization or commits.

## Repository Documentation Architecture

The prepared repository will use three documentation layers:

1. `README.md` for human onboarding and quick project context.
2. `.github/copilot-instructions.md` for agent workflow rules and repository conventions.
3. `docs/wiki/` for the persistent, incrementally maintained LLM wiki.

The wiki follows the structure described in `docs/wiki/llm-wiki.md`:

- `docs/wiki/index.md` is the content-oriented catalog.
- `docs/wiki/log.md` is the chronological append-only history.
- Additional wiki pages capture important project areas and core functionality.

The existing `docs/raw/idea.md` remains a raw source. The existing `docs/wiki/llm-wiki.md` remains the reference pattern for wiki maintenance.

## Wiki Pages

The initial wiki pages should be minimal templates, not full architecture documents. They should create places for future decisions without prematurely making those decisions.

Initial pages to create:

- `project-overview.md`: objective, stack, current status, and links to core areas.
- `architecture.md`: template for future architecture decisions.
- `backend.md`: template for Go backend notes, API shape, realtime transport, and tests.
- `frontend.md`: template for React, Vite, Tailwind, UI states, and tests.
- `realtime-protocol.md`: template for events, payloads, connection lifecycle, and presence.
- `docker.md`: template for container and local environment notes.
- `testing.md`: template for backend, frontend, and integration testing strategy.

Each page should include stable sections such as `Status`, `Purpose`, `Current notes`, `Decisions pending`, and `Related pages`.

## Agent Workflow Rules

`.github/copilot-instructions.md` should make these rules explicit:

- Work in Portuguese unless the user asks otherwise.
- Do not implement beyond the current request.
- Do not create commits that mention AI, Copilot, or similar AI attribution unless explicitly requested.
- Read `docs/wiki/index.md` before modifying or relying on wiki content.
- Update `docs/wiki/index.md` whenever wiki pages are created, renamed, or materially changed.
- Append to `docs/wiki/log.md` after meaningful documentation changes, completed implementation plans, or wiki ingests.
- After implementing future plans, ingest the implementation into the wiki by recording what changed, where it lives, how to validate it, and which pages were updated.

## Data Flow

For documentation work:

1. Read `README.md`, `.github/copilot-instructions.md`, and `docs/wiki/index.md` when relevant.
2. Update the target documentation or wiki page.
3. Update `docs/wiki/index.md` if page inventory or summaries changed.
4. Append a dated entry to `docs/wiki/log.md`.

For future implementation work:

1. Execute the approved implementation plan.
2. Run the relevant validation commands.
3. Ingest the result into the wiki.
4. Update affected wiki pages and `index.md`.
5. Append a completion entry to `log.md`.

## Error Handling and Consistency

Because this preparation is documentation-only, error handling means avoiding ambiguous or stale project guidance:

- Avoid loose `TODO` markers. Use explicit sections like `Decisions pending` and `Next steps`.
- Do not invent unapproved architecture decisions.
- Keep links relative and consistent.
- Keep `log.md` append-only.
- Treat `docs/raw/` as source material, not generated wiki content.
- If Git is unavailable or intentionally unused, skip commits and state that clearly.

## Validation

The completed preparation should be checked by confirming:

- `README.md` exists and links to the wiki.
- `.github/copilot-instructions.md` exists and includes the workflow rules above.
- `docs/wiki/index.md` exists and catalogs all wiki pages.
- `docs/wiki/log.md` exists and includes an entry for this repository preparation.
- Core wiki template pages exist and cross-link through `index.md`.
- No application scaffold files are created as part of this step.
