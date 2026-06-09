# Frontend

## Status

Frontend React/Vite/Tailwind implementado com reducer, hook WebSocket, componentes de chat e testes unitários.

## Purpose

Registrar decisões e notas sobre a interface web do chat.

## Current notes

- App principal: `src/App.tsx`.
- Hook WebSocket: `src/hooks/useChatConnection.ts`, com conexão, envio de mensagens, erro local e reconexão automática.
- Estado determinístico: `src/chat/chatReducer.ts`.
- Tipos compartilhados do frontend: `src/types/chat.ts`.
- Componentes: `src/components/ConnectionStatus.tsx`, `JoinForm.tsx`, `MessageComposer.tsx`, `MessageList.tsx` e `PeopleList.tsx`.
- Tailwind/Vite/Vitest configurados por `tailwind.config.js`, `postcss.config.js`, `vite.config.ts` e `vitest.setup.ts`.
- Validações: `cd frontend && npm run test:run` e `cd frontend && npm run build`.

## Decisions pending

- Testes e2e automatizados.
- Acessibilidade avançada e navegação por teclado completa.
- Estados de reconexão mais detalhados para UX.

## Related pages

- [Arquitetura](architecture.md)
- [Protocolo realtime](realtime-protocol.md)
- [Testes](testing.md)
