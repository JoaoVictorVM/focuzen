# Monorepo único para web, backend e CLI

- Status: aceito
- Data: 2026-06-16
- Decisores: João Victor Ventura Martins

## Contexto e problema

O Focuzen tem três artefatos — SPA web, backend Go e CLI Go — que formam um único
produto. Precisamos decidir se eles vivem em um repositório único ou em repositórios
separados.

## Fatores de decisão

- As três frentes evoluem juntas e compartilham contexto (docs, ADRs, convenções).
- Projeto é peça de portfólio: a navegação por um único repo facilita a avaliação.
- Tamanho pequeno do time (uma pessoa) — overhead de múltiplos repos não se paga.

## Opções consideradas

- **Monorepo**: tudo em um repositório.
- **Polirepo**: um repositório por artefato.

## Decisão

Escolhida: **Monorepo**, porque web, backend e CLI são o mesmo produto e versionam
juntos. Um único histórico, um único CI e documentação centralizada reduzem o atrito.
O isolamento técnico necessário (CGO da CLI vs. binário estático do server) é resolvido
com dois módulos Go ligados por `go.work` (ver [ADR-0005](0005-dois-modulos-go.md)),
sem precisar de repositórios separados.

### Consequências

- Boa, porque há um único ponto de verdade para issues, PRs, CI e docs.
- Boa, porque mudanças que cruzam camadas (ex.: contrato da API) ficam em um único PR.
- Ruim, porque o CI precisa de filtros por caminho para não rodar tudo a cada mudança.

## Mais informações

Relacionado: [ADR-0005](0005-dois-modulos-go.md).
