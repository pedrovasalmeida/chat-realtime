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
3. Acrescente entradas em `docs/wiki/log.md` após mudanças documentais relevantes, implementações concluídas ou ingestões finais.
4. Em planos de implementação, deixe a ingestão da wiki apenas para o último step, quando o plano já estiver devidamente implementado; registre o que mudou, onde validar, arquivos relevantes e páginas atualizadas.
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
