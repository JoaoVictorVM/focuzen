# "Próximo" via fila local em vez de recomendação do YouTube

- Status: aceito
- Data: 2026-06-16
- Decisores: João Victor Ventura Martins

## Contexto e problema

O controle de áudio tem um botão "Próximo". O comportamento natural seria pedir a próxima
recomendação ao YouTube, mas a API removeu o `relatedToVideoId` do `search.list` em 2023,
deixando de expor recomendações relacionadas.

## Fatores de decisão

- A API do YouTube não oferece mais recomendações de "próximo vídeo".
- O comportamento precisa ser previsível e não custar cota extra.
- Não introduzir dependências de scraping ou de endpoints não suportados.

## Opções consideradas

- **Avançar numa fila local** (resultados da busca atual ou playlist de foco curada).
- **Recomendação algorítmica do YouTube** — inviável (recurso removido da API).

## Decisão

Escolhida: **fila local**. O "Próximo" avança para o próximo item da lista de resultados da
busca ou de uma playlist de foco curada (modo "rádio"). Não é recomendação algorítmica — é
uma fila determinística no cliente, que não consome cota adicional.

### Consequências

- Boa, porque o comportamento é previsível e não gasta cota da API.
- Ruim, porque não há "descoberta" automática de conteúdo novo como no app do YouTube.
- Neutra: playlists curadas podem ser adicionadas depois para enriquecer a fila.

## Mais informações

PRD §3.3 e §18. Cota: ver [ADR-0008](0008-budget-e-limites.md).
