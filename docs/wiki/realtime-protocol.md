# Realtime Protocol

## Status

Protocolo JSON inicial implementado para eventos `join`, `message`, `presence` e `error`.

## Purpose

Registrar eventos, payloads e estados de conexão usados entre frontend e backend.

## Current notes

- O projeto usa WebSocket em `/ws` para comunicação em tempo real.
- A primeira versão tem uma sala global única.
- Cliente entra enviando:

```json
{ "type": "join", "name": "Alice" }
```

- Cliente envia mensagem com:

```json
{ "type": "message", "content": "hello" }
```

- Servidor transmite mensagens com:

```json
{
  "type": "message",
  "message": {
    "id": "m1",
    "userId": "u1",
    "userName": "Alice",
    "content": "hello",
    "sentAt": "2026-06-09T18:30:00Z"
  }
}
```

- Servidor transmite snapshots de presença com:

```json
{
  "type": "presence",
  "users": [{ "id": "u1", "name": "Alice" }]
}
```

- Servidor reporta erros com:

```json
{ "type": "error", "error": "message content is required" }
```

## Decisions pending

- Códigos de erro padronizados além das mensagens textuais.
- Tratamento de mensagens duplicadas ou fora de ordem.
- Versionamento do protocolo.

## Related pages

- [Arquitetura](architecture.md)
- [Backend](backend.md)
- [Frontend](frontend.md)
- [Testes](testing.md)
