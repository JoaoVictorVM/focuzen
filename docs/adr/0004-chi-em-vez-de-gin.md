# Roteador chi em vez de Gin

- Status: aceito
- Data: 2026-06-16
- Decisores: João Victor Ventura Martins

## Contexto e problema

O backend precisa de um roteador HTTP com middlewares (rate limit, security headers,
request ID). Precisamos escolher entre um roteador idiomático sobre a stdlib ou um
framework mais opinativo.

## Fatores de decisão

- Compatibilidade com `net/http` / `http.Handler` (testes com `httptest`).
- Idiomático e leve, sem abstrações próprias que afastem da stdlib.
- Ecossistema de middlewares maduro.

## Opções consideradas

- **chi**: roteador fino, 100% compatível com `http.Handler`.
- **Gin**: framework com contexto próprio e mais conveniências, porém menos aderente à stdlib.
- **net/http puro**: sem dependência, mas roteamento/middlewares manuais.

## Decisão

Escolhida: **chi**, porque é idiomático e totalmente compatível com `http.Handler`, o que
mantém handlers e middlewares triviais de testar com `httptest` e alinhados à stdlib. Gin
introduz um `Context` próprio que reduz essa aderência; net/http puro exigiria reimplementar
roteamento e middlewares.

### Consequências

- Boa, porque handlers permanecem `http.HandlerFunc` simples e testáveis.
- Boa, porque os middlewares do chi (RequestID, etc.) cobrem boa parte das necessidades.
- Neutra: abrimos mão das conveniências do Gin (binding/validação embutidos), que não são
  necessárias para um proxy de busca enxuto.
