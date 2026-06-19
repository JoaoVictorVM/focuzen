# Deploy — Render (Blueprint)

Hospedagem do Focuzen na [Render](https://render.com) (free web service, **sem cartão**;
ver [ADR-0009](adr/0009-hospedagem-render.md)). A infra é declarada no
[`render.yaml`](../render.yaml) na raiz do repositório (Render Blueprint): um web service
que builda o `Dockerfile` (SPA embarcada no binário Go), com health check em `/healthz`.

> Trade-off: o free tier **hiberna** após ~15 min de inatividade (cold start ~30-60s).

## Setup (uma vez)

1. Em [dashboard.render.com](https://dashboard.render.com), **New → Blueprint** e conecte
   este repositório. A Render lê o `render.yaml` e cria o serviço.
2. No serviço, em **Environment**, defina o secret **`YOUTUBE_API_KEY`** (está como
   `sync: false` no blueprint, então não vem versionado).
3. Em **Settings → Deploy Hook**, copie a URL do hook.
4. No GitHub, crie o repository secret **`RENDER_DEPLOY_HOOK_URL`** com essa URL.

## Deploy

`autoDeploy` está **desligado** no blueprint; os deploys são disparados pelo workflow de
CD ([`.github/workflows/cd.yml`](../.github/workflows/cd.yml)) em push na `main`, que
chama o Deploy Hook. A Render builda o `Dockerfile` e publica a nova versão.

Para um deploy manual: rode o workflow **CD** (aba Actions → Run workflow) ou bata na URL
do hook (`curl -fsS "$RENDER_DEPLOY_HOOK_URL"`).

## Variáveis de ambiente

| Variável | Onde | Descrição |
|---|---|---|
| `YOUTUBE_API_KEY` | secret na Render | Chave da YouTube Data API v3 (obrigatória). |
| `PORT` | injetada pela Render | A app escuta na porta que a Render define. |
