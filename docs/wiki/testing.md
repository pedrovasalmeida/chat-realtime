# Testing

## Status

Testes Go, testes unitários frontend, builds e validação Docker configurados para a primeira versão.

## Purpose

Registrar a estratégia de testes do projeto conforme backend, frontend e integração forem planejados.

## Current notes

- Backend:

```bash
cd backend && go test ./...
```

- Frontend:

```bash
cd frontend && npm run test:run
cd frontend && npm run build
```

- Docker:

```bash
docker compose config
```

- Smoke manual recomendado: abrir `http://localhost:5173` em duas janelas, entrar como `Alice` e `Bob`, enviar `hello` como Alice, confirmar presença/mensagem em ambas e confirmar remoção de Bob ao fechar a janela.
- Nesta implementação, a integração também foi validada com dois clientes WebSocket automatizados em portas alternativas por conflito local na porta `8080`.

## Decisions pending

- Automatizar smoke em browser.
- Cobrir reconexão do hook WebSocket com testes específicos.

## Related pages

- [Backend](backend.md)
- [Frontend](frontend.md)
- [Protocolo realtime](realtime-protocol.md)
- [Docker](docker.md)
