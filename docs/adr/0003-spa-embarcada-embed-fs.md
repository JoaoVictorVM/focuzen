# SPA embarcada no binário via embed.FS

- Status: aceito
- Data: 2026-06-16
- Decisores: João Victor Ventura Martins

## Contexto e problema

A SPA React precisa ser servida em produção. Podemos hospedá-la separadamente (CDN/host
estático) ou embuti-la no próprio binário Go que já roda o backend.

## Fatores de decisão

- Simplicidade de deploy (idealmente um único artefato).
- Evitar CORS e configuração de origem entre front e API.
- Custo zero (sem host estático adicional).

## Opções consideradas

- **Embarcar a SPA no binário** com `embed.FS`.
- **Host estático separado** (ex.: CDN/Pages) + backend à parte.

## Decisão

Escolhida: **`embed.FS`**, porque produz um único artefato que serve tanto a API quanto a
SPA na **mesma origem** — zero CORS em produção e um único deploy. O Vite gera `web/dist`,
que é copiado para `server/internal/webui/dist` (gitignored) antes do build do Go.

### Consequências

- Boa, porque o deploy é de um binário só, mesma origem, sem CORS.
- Ruim, porque o `server` só compila com a SPA já buildada e copiada para o diretório do
  embed — uma dependência de ordem no build a ser respeitada no Dockerfile e no Makefile.
- Neutra: CORS ainda é configurado, mas só relevante no desenvolvimento local.

## Mais informações

PRD §5.1 e §7. Relacionado: [ADR-0001](0001-monorepo.md).
