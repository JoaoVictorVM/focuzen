# Dois módulos Go ligados por go.work

- Status: aceito
- Data: 2026-06-16
- Decisores: João Victor Ventura Martins

## Contexto e problema

O backend precisa compilar como binário **estático sem CGO** (para a imagem distroless),
enquanto a CLI **depende de CGO** (áudio via Beep). Um único módulo Go forçaria uma única
configuração de CGO para todo o código.

## Fatores de decisão

- Server deve ser `CGO_ENABLED=0` para um binário estático na distroless.
- CLI precisa de `CGO_ENABLED=1` por causa da camada de áudio.
- Manter o desenvolvimento local ergonômico (um clone, um editor).

## Opções consideradas

- **Dois módulos Go (`server/`, `cli/`) + `go.work`**.
- **Módulo único** com build tags para isolar o áudio.
- **Repositórios separados** (descartado pelo [ADR-0001](0001-monorepo.md)).

## Decisão

Escolhida: **dois módulos Go ligados por `go.work`**. Isso isola completamente as
exigências de CGO: `server/` compila sem CGO, `cli/` com CGO, cada um com seu `go.mod`. O
`go.work` mantém a ergonomia de desenvolvimento local (resolução de ambos os módulos de uma
vez) sem acoplar suas dependências.

### Consequências

- Boa, porque cada binário tem build independente e dependências isoladas.
- Boa, porque evita malabarismo de build tags em um módulo único.
- Ruim, porque há dois `go.mod` para manter e o `go.work` precisa ser entendido por quem
  contribui.

## Mais informações

PRD §7. Relacionado: [ADR-0001](0001-monorepo.md).
