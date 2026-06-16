# Hospedagem do backend na Koyeb

- Status: aceito
- Data: 2026-06-16
- Decisores: João Victor Ventura Martins

## Contexto e problema

O backend precisa de um host para um container Go de longa duração. Restrições de
projeto: custo **R$ 0** e **sem cartão de crédito**. Precisamos de uma plataforma que
sirva HTTP continuamente sem cobrar nem exigir cartão.

## Fatores de decisão

- Custo zero garantido, sem necessidade de cartão de crédito.
- Sem "sleep" / cold start agressivo que prejudique a experiência.
- Suporte a deploy de container e a Infrastructure as Code.

## Opções consideradas

- **Koyeb**: free tier sem cartão, sem sleep, deploy de container.
- **Render / Railway / Fly.io**: free tiers que exigem cartão e/ou colocam o serviço
  para dormir.
- **VPS própria**: custo mensal recorrente.

## Decisão

Escolhida: **Koyeb**, porque atende o requisito inflexível de custo zero sem cartão e
mantém o serviço acordado. Fornece uma URL pública gratuita, dispensando domínio próprio.
O provisionamento é feito com Terraform.

### Consequências

- Boa, porque elimina qualquer risco de cobrança surpresa.
- Ruim, porque amarra parte da infra ao provider Koyeb do Terraform.
- Neutra: o cache em memória é best-effort (a instância pode reiniciar/escalar) — aceito
  para o MVP, ver [ADR-0008](0008-budget-e-limites.md) e o PRD §5.3.
