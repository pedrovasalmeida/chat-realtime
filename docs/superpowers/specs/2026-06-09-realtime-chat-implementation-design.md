# Realtime Chat Implementation Design

## Status

Approved on 2026-06-09.

## Goal

Plan the first implementation of a realtime chat with Go backend, React/Vite/Tailwind frontend, backend and frontend unit tests, and Docker-based local development.

This design is documentation-only. It does not scaffold backend, frontend, tests, or Docker files.

## Scope

In scope:

- A single implementation spec with separate stages for Docker, backend, frontend, integration, and wiki ingest.
- One global chat room.
- User identification by name entered on join, without authentication.
- Realtime communication over WebSocket.
- Messages, presence, and connections held in memory for the first version.
- Backend tests in Go.
- Frontend unit tests with the Vite/React testing stack.
- Docker Compose for local development of backend and frontend.

Out of scope:

- Message persistence after server restart.
- Authentication, registration, or login.
- Multiple rooms.
- End-to-end tests.
- Production deployment.
- Git initialization or commits.

## Architecture

The repository should remain a simple monorepo:

| Path | Responsibility |
| --- | --- |
| `backend/` | Go HTTP/WebSocket server, in-memory chat hub, backend tests. |
| `frontend/` | React/Vite/Tailwind app, chat UI, frontend unit tests. |
| `docker-compose.yml` | Local development orchestration. |
| `docs/wiki/` | Persistent project decisions and implementation ingest. |

The backend is the source of truth for connected users, presence, and messages while the process is running. The frontend renders the approved classic chat layout: messages and input on the left, people online on the right.

## Realtime Protocol

Use a WebSocket endpoint exposed by the backend, planned as `GET /ws`.

The implementation should use Go's standard `net/http` plus `nhooyr.io/websocket`, chosen for a small API surface, context-aware connection handling, and compatibility with `httptest`.

Initial events should be JSON:

| Event | Direction | Purpose |
| --- | --- | --- |
| `join` | client -> server | Send the chosen display name after opening the socket. |
| `message` | client -> server | Send chat message content. |
| `message` | server -> clients | Broadcast accepted chat messages. |
| `presence` | server -> clients | Broadcast the current users list after join/leave changes. |
| `error` | server -> client | Report invalid payloads or rejected actions. |

Server-generated messages should include:

- `id`
- `userId`
- `userName`
- `content`
- `sentAt`

The server generates message IDs and timestamps. The client may generate temporary UI state, but accepted messages come from the server broadcast.

## Backend Components

The backend should be split into small units:

- HTTP server/bootstrap: starts routes, reads configuration, and owns process lifecycle.
- WebSocket handler: upgrades connections, parses client events, writes server events.
- Chat hub: stores active clients in memory, handles join/leave, broadcasts messages and presence snapshots.
- Protocol types/validation: defines event payloads and validation rules.

The hub should not depend on HTTP. This keeps message distribution and presence behavior unit-testable without a running server.

## Frontend Components

The frontend should use the approved classic layout:

- Join/name screen before connecting or before sending the first `join` event.
- Chat shell with messages list, composer, send button, and people list on the right for desktop.
- Connection status display: connecting, connected, disconnected, and error.
- Message item showing user name/id and time.
- Responsive behavior that keeps the people list usable on smaller screens.

When the implementation reaches frontend work, use the `frontend-design` skill before producing UI code so the visual result stays polished while remaining simple.

## Docker Development

Docker Compose should provide a local development environment with at least:

- `backend` service exposing the Go server on port `8080`.
- `frontend` service exposing Vite on port `5173`.
- Volume mounts for source code during development.
- Frontend environment pointing to the backend WebSocket URL, e.g. `VITE_WS_URL=ws://localhost:8080/ws`.

The frontend should use Vite HMR. The backend development container should use Air for Go hot reload.

## Error Handling and Connection States

Backend:

- Validate join names and message content.
- Send explicit `error` events for invalid payloads.
- Do not crash or drop healthy clients because one client sends bad data.
- Remove disconnected clients from the hub and broadcast updated presence.

Frontend:

- Disable sending while disconnected or while the message is empty.
- Show connection state and last relevant error.
- Attempt simple reconnection after a short delay.
- Do not queue offline messages in the first version.

## Testing Strategy

Backend tests are required:

- Unit tests for hub join/leave/presence behavior.
- Unit tests for message validation and event payload parsing.
- WebSocket handler tests covering successful join, message broadcast, presence update, disconnect cleanup, and invalid payload errors.
- Validation command: `go test ./...`.

Frontend tests are required, but only unit-level:

- Component tests for message list rendering, composer validation, people list rendering, and connection status.
- Hook/state tests for WebSocket state management if that logic is extracted from components.
- Suggested stack: Vitest + Testing Library.
- No end-to-end tests in the first version.

Docker validation should confirm that Compose starts both services and that the frontend can connect to the backend.

## Implementation Stages

1. Docker/dev foundation: define service boundaries, ports, environment variables, and local commands.
2. Backend: create Go module, server, protocol types, in-memory hub, WebSocket endpoint, and tests.
3. Frontend: create React/Vite/Tailwind app, classic chat UI, WebSocket client state, and unit tests.
4. Integration: run both services, verify join, message broadcast, presence updates, disconnect handling, and validation errors.
5. Wiki ingest: update the relevant wiki pages with files changed, validation commands, and implementation status; append the log.

## Documentation and Git

Approved decisions from this design should be reflected in the wiki pages for architecture, backend, frontend, realtime protocol, Docker, and testing.

The current directory is not a Git repository. Do not initialize Git just to satisfy the design workflow. The design document can be written and reviewed, but the commit step is not applicable unless the user explicitly asks to initialize or use Git later.
