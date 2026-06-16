# Contraste de texto abaixo do WCAG AA (escolha estética)

- Status: aceito
- Data: 2026-06-16
- Decisores: João Victor Ventura Martins

## Contexto e problema

A identidade visual usa a cor primária `#969696` sobre o fundo `#f3f3f3`. Essa combinação
resulta em contraste de ~2.7:1, abaixo do mínimo do WCAG AA (4.5:1 para texto normal).
Precisamos registrar conscientemente esse desvio em vez de "corrigi-lo" silenciosamente.

## Fatores de decisão

- A estética minimalista e suave é central à identidade do produto.
- Acessibilidade é um valor do projeto (HTML semântico, teclado, foco visível).
- Decisões de design devem ser explícitas e revisáveis, não acidentais.

## Opções consideradas

- **Manter `#969696` sobre `#f3f3f3`** como escolha estética consciente.
- **Escurecer o texto** para atingir 4.5:1, alterando a identidade visual.

## Decisão

Escolhida: **manter o contraste atual** como decisão estética consciente e documentada.
As cores vêm de CSS variables num único lugar, então um ajuste futuro é trivial e será
**alinhado**, nunca feito silenciosamente. As demais práticas de acessibilidade (semântica,
navegação por teclado, foco visível, `aria-*`) permanecem em vigor.

### Consequências

- Boa, porque preserva a identidade visual pretendida.
- Ruim, porque o texto de baixo contraste reprova no WCAG AA — um custo de acessibilidade
  assumido de forma explícita.
- Neutra: por virem de CSS variables, as cores são revisáveis num único ponto.

## Mais informações

PRD §3.5 e §9. CLAUDE.md (gotcha de contraste).
