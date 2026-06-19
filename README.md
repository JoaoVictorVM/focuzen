# Focuzen

[![CI](https://github.com/JoaoVictorVM/focuzen/actions/workflows/ci.yml/badge.svg)](https://github.com/JoaoVictorVM/focuzen/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

App de foco minimalista em duas frentes que compõem o mesmo produto:

- **Web** — tela de foco no browser: relógio em destaque, busca no YouTube e player de áudio embutido (vídeo oculto, só o áudio toca).
- **CLI** — TUI de foco no terminal: relógio, interface limpa e rádios de foco via streams MP3.

MVP deliberadamente enxuto: **sem contas de usuário e sem banco de dados** (estado local).
O foco está na qualidade de engenharia: testes, segurança, CI/CD, IaC e documentação.

> **Demo**: https://focuzen.onrender.com — no free tier o serviço hiberna; o primeiro acesso pode levar ~30-60s (cold start).

## Funcionalidades

- **Web**: relógio + data por extenso; busca no YouTube (com debounce e cache); player oculto com play/pause, anterior/próximo, repetir, mudo, volume e tela cheia; tema claro/escuro; i18n pt-BR/en.
- **CLI**: relógio centralizado; menu de rádios de foco (áudio via Beep); i18n pt-BR/en.

## Stack

- **Web**: React + TypeScript (strict) + TailwindCSS (cores via CSS variables) + Vite.
- **Backend**: Go + chi + `log/slog` — serve a SPA (`embed.FS`) e o proxy de busca do YouTube.
- **CLI**: Go + Bubble Tea + Lip Gloss + Beep.
- **Infra**: Docker (distroless, non-root) + Render (Blueprint `render.yaml`). CLI distribuída via GitHub Releases.
- **CI/CD**: GitHub Actions; versionamento com Conventional Commits + release-please.

## Estrutura

Monorepo com dois módulos Go (`server/` sem CGO, `cli/` com CGO) ligados por `go.work`,
além do frontend em `web/`.

```
focuzen/
├── server/   # backend Go (chi, embed da SPA, proxy de busca)
├── cli/      # CLI Go (Bubble Tea + Beep)
├── web/      # frontend React + TS + Tailwind + Vite
├── docs/     # ADRs, diagramas C4 e guia de deploy
└── render.yaml, Dockerfile, Makefile, .github/
```

## Pré-requisitos

- **Go** 1.26+
- **Node** 22+ e **pnpm** 10+
- Para a CLI no Linux: `libasound2-dev` (CGO do áudio)
- Uma chave da **YouTube Data API v3** (apenas no backend) — veja [como obter](https://console.cloud.google.com/apis/library/youtube.googleapis.com).

## Desenvolvimento

Backend e web rodam separados em dev (o Vite faz proxy de `/api` e `/download` para o backend):

```sh
# Terminal 1 — backend
cd server
YOUTUBE_API_KEY="sua-chave" go run ./cmd/server

# Terminal 2 — web
cd web
pnpm install
pnpm dev
```

CLI:

```sh
cd cli
go run ./cmd/focus
```

## Build, testes e lint

```sh
make build   # builda a SPA, embute no binário do server, compila tudo
make test    # testes de server, cli e web
make lint    # golangci-lint + typecheck do web
```

Ou por módulo: `cd server && go test ./...`, `cd web && pnpm test`, etc. E2E do web: `cd web && pnpm e2e`.

## Deploy

A web roda na Render via Blueprint (`render.yaml`), com deploy disparado pelo CD.
Veja o guia em [docs/deployment.md](docs/deployment.md).

## CLI

Baixe o binário do seu sistema em [Releases](https://github.com/JoaoVictorVM/focuzen/releases)
(ou pelo link `/download` no rodapé da web) e rode `./focus`.

## Documentação

- [`docs/c4/`](docs/c4/) — diagramas C4 (contexto, containers, componentes).
- [`docs/adr/`](docs/adr/) — decisões de arquitetura (ADRs).
- [`docs/deployment.md`](docs/deployment.md) — deploy na Render.

## Contribuindo

Veja [CONTRIBUTING.md](CONTRIBUTING.md). Reportes de segurança: [SECURITY.md](SECURITY.md).
Esperamos um ambiente acolhedor — ver [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md).

## Licença

[MIT](LICENSE) © João Victor Ventura Martins
