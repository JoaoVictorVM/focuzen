# C4 — Componentes (Backend)

Por dentro do container do backend. A busca passa por um cache LRU (decorator)
antes de chegar ao cliente do YouTube, protegendo a cota. As dependências externas
ficam atrás de interfaces (port `Searcher`).

```mermaid
C4Component
  title Diagrama de Componentes — Backend

  Container_Boundary(backend, "Backend (Go)") {
    Component(router, "Router", "chi", "Monta middlewares e rotas")
    Component(mw, "Middleware", "Go", "Security headers (CSP) e rate limit por IP")
    Component(health, "Health handlers", "Go", "/healthz e /readyz")
    Component(search, "Search handler", "Go", "/api/v1/search — valida e sanitiza o termo")
    Component(download, "Download handler", "Go", "/download — redireciona pro GitHub Releases")
    Component(webui, "WebUI", "embed.FS", "Serve a SPA buildada (rota coringa)")
    Component(cache, "Cache LRU", "Go", "Decorator do Searcher — protege a cota")
    Component(yt, "YouTube client", "Go", "Port Searcher: chama o search.list")
    Component(config, "Config", "Go", "Carrega env vars (chave, limites, URLs)")
  }

  System_Ext(youtube, "YouTube Data API v3")

  Rel(router, mw, "aplica")
  Rel(router, health, "roteia")
  Rel(router, search, "roteia (sob rate limit)")
  Rel(router, download, "roteia")
  Rel(router, webui, "roteia /*")
  Rel(search, cache, "consulta")
  Rel(cache, yt, "delega no cache miss")
  Rel(yt, youtube, "search.list", "HTTPS")
  Rel(config, yt, "fornece a API key")
```
