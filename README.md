# Focuzen

App de foco minimalista em duas frentes que compõem o mesmo produto:

- **Web** — tela de foco no browser: relógio em destaque, busca no YouTube e player de áudio embutido (vídeo oculto, só o áudio toca).
- **CLI** — TUI de foco no terminal: relógio, interface limpa e rádios de foco via streams MP3.

MVP deliberadamente enxuto: **sem contas de usuário e sem banco de dados** (estado local).
O foco está na qualidade de engenharia (testes, segurança, CI/CD, IaC, documentação).

## Stack

- **Web**: React + TypeScript (strict) + TailwindCSS + Vite.
- **Backend**: Go + chi + `log/slog` — serve a SPA (`embed.FS`) e o proxy de busca do YouTube.
- **CLI**: Go + Bubble Tea + Lip Gloss + Beep.
- **Infra**: Docker (distroless) + Terraform (Koyeb). CLI distribuída via GitHub Releases.

Monorepo com dois módulos Go (`server/` sem CGO, `cli/` com CGO) ligados por `go.work`,
além do frontend em `web/`.

## Documentação

- [`PRD.md`](PRD.md) — especificação completa e roadmap.
- [`CLAUDE.md`](CLAUDE.md) — convenções e fluxo de trabalho.
- `docs/adr/` — decisões de arquitetura (ADRs).

## Licença

[MIT](LICENSE) © João Victor Ventura Martins
