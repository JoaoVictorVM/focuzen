# Política de Segurança

## Como reportar uma vulnerabilidade

**Não abra uma issue pública** para vulnerabilidades. Use o canal privado do GitHub:

1. Acesse a aba **Security** do repositório.
2. Clique em **Report a vulnerability** (Private vulnerability reporting).

Descreva o problema, o impacto e os passos para reproduzir. Respondemos o quanto antes
e creditamos quem reportar, se desejar.

## Versões suportadas

Este é um projeto de portfólio sem releases LTS; correções de segurança são aplicadas na
linha mais recente (`main` e o último release).

## Medidas de segurança do projeto

- A **chave da YouTube API nunca vai para o frontend** — fica só no backend, via secret.
- **Rate limiting** por IP nas rotas de API (protege cota e contra abuso).
- **Security headers**: CSP (libera o player do YouTube de forma consciente),
  `X-Content-Type-Options`, `X-Frame-Options`, `Referrer-Policy`, HSTS.
- **Validação/sanitização** do parâmetro de busca.
- **Scans** no CI: `govulncheck` (Go), `pnpm audit` (web) e Trivy (imagem).
- **Container** distroless, multi-stage, usuário non-root.
