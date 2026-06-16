# Budget e limites de recursos desde o dia zero

- Status: aceito
- Data: 2026-06-16
- Decisores: João Victor Ventura Martins

## Contexto e problema

O projeto precisa rodar a custo **R$ 0** e a YouTube Data API tem cota diária de **10.000
unidades**, com cada `search.list` custando **100** (~100 buscas/dia). Sem limites desde o
início, há risco de estourar cota (degradando o serviço) ou de surpresas de custo na infra.

## Fatores de decisão

- Custo zero garantido, sem cartão de crédito.
- Proteção da cota da YouTube API contra abuso e picos.
- Previsibilidade de recursos no container (sem surpresas de billing).

## Opções consideradas

- **Definir limites/budget desde o dia zero** (cache, debounce, rate limit, limites do container).
- **Adicionar limites reativamente**, só quando um problema aparecer.

## Decisão

Escolhida: **limites desde o dia zero**. Combinamos três camadas para proteger a cota:
**cache LRU em memória**, **debounce** na digitação do frontend e **rate limit por IP** no
backend. Os recursos do container (CPU/memória) são limitados no Terraform/Koyeb. Isso torna
o custo previsível e protege a cota antes de qualquer incidente.

### Consequências

- Boa, porque elimina o risco de estourar a cota ou de custo inesperado.
- Boa, porque cache + debounce + rate limit se reforçam mutuamente.
- Neutra: o cache em memória é best-effort na Koyeb (não compartilhado nem persistente);
  suficiente para o MVP, com Firestore free tier como evolução futura em aberto.

## Mais informações

PRD §5.3 e §14. Relacionado: [ADR-0002](0002-hospedagem-koyeb.md), [ADR-0006](0006-proximo-via-fila-local.md).
