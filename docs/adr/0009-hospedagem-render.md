# Hospedagem do backend na Render

- Status: aceito
- Data: 2026-06-18
- Decisores: João Victor Ventura Martins

## Contexto e problema

O backend precisa de um host para um container Go. A restrição inflexível continua:
custo **R$ 0** e **sem cartão de crédito**. A Koyeb (ver [ADR-0002](0002-hospedagem-koyeb.md))
era a escolha original, mas removeu o free tier sem cartão — agora exige cartão e um
plano pago. Precisamos de um substituto que mantenha o custo zero sem cartão.

## Fatores de decisão

- Custo zero garantido, **sem cartão de crédito** (requisito inflexível).
- Suporte a deploy de container (Docker) e a Infrastructure as Code.
- URL pública gratuita (dispensa domínio próprio).

## Opções consideradas

- **Render (free web service)**: sem cartão; builda a imagem do nosso Dockerfile;
  health checks, env secrets e Blueprint (`render.yaml`) como IaC. Trade-off: **dorme**
  após ~15 min de inatividade (cold start ~30-60s) e 750 h/mês.
- **Hugging Face Spaces (Docker)**: sem cartão, mas é plataforma de demos (menos
  convencional para um app web "de verdade").
- **Fly.io / Google Cloud Run / Railway**: ótimos, mas passaram a **exigir cartão**.

## Decisão

Escolhida: **Render**, porque é a opção que mantém o requisito inflexível (grátis, sem
cartão) e é a mais próxima do modelo de container da Koyeb (Docker + health check + env
secrets). O "no sleep", que era um *desejo* e não requisito, é abandonado conscientemente:
o serviço hiberna quando ocioso e tem cold start no primeiro acesso. A IaC passa a ser o
**Blueprint `render.yaml`** (nativo da Render) no lugar do Terraform da Koyeb.

### Consequências

- Boa, porque preserva o custo zero sem cartão.
- Ruim, porque há **cold start** após inatividade (degrada o primeiro acesso) — aceitável
  para um MVP de portfólio.
- Neutra: o cache em memória continua best-effort (a instância hiberna/reinicia) — o que
  reforça o já registrado em [ADR-0008](0008-budget-e-limites.md).
- Neutra: deploy via Render (build do Dockerfile + deploy hook), substituindo o fluxo de
  imagem no GHCR + Koyeb.

## Mais informações

Substitui [ADR-0002](0002-hospedagem-koyeb.md). Blueprint em `render.yaml`; deploy
disparado pelo workflow de CD via Deploy Hook.
