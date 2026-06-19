# Contribuindo com o Focuzen

Obrigado pelo interesse! Este guia resume como configurar o ambiente e abrir uma boa contribuição.

## Setup

1. Instale os pré-requisitos: Go 1.26+, Node 22+, pnpm 10+ (e `libasound2-dev` no Linux para o áudio da CLI).
2. Instale os hooks de commit (commitlint via Lefthook):
   ```sh
   pnpm install
   ```
3. Veja o [README](README.md) para rodar web, backend e CLI em desenvolvimento.

## Fluxo

1. Crie uma branch a partir da `main`.
2. Faça mudanças pequenas e coesas.
3. Antes de abrir o PR, garanta que passa localmente:
   ```sh
   make lint
   make test
   ```
4. Abra o PR preenchendo o template. O CI (lint, testes, build e scans) precisa passar.

## Commits — Conventional Commits

As mensagens seguem [Conventional Commits](https://www.conventionalcommits.org/) e são
validadas pelo commitlint no hook de `commit-msg`.

Formato: `tipo(escopo): descrição` (minúsculas, no imperativo).

- **Tipos**: `feat`, `fix`, `docs`, `test`, `refactor`, `chore`, `build`, `ci`, `perf`.
- **Escopos comuns**: `server`, `web`, `cli`, `infra`, `e2e`.

Exemplos: `feat(web): add repeat toggle`, `fix(server): correct module path`.

O versionamento e o CHANGELOG são automáticos (release-please) a partir desses commits.

## Estilo de código

- **Go**: idiomático, erros explícitos, interfaces pequenas, testes table-driven. `gofmt` + `golangci-lint` limpos.
- **TypeScript**: modo `strict`, sem `any` implícito.
- **Cores/tokens** no web sempre via CSS variables (nada hardcoded).
- Comentários só quando agregam (o *porquê*); identificadores em inglês; texto de usuário via i18n.

## Testes

- Backend: `testing` + testify (`httptest` para mocks).
- Web: Vitest + Testing Library + MSW; E2E com Playwright.
- CLI: teatest + unit, com a camada de áudio mockada.
