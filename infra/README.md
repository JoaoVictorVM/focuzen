# Infra — Koyeb (Terraform)

Provisiona o Focuzen na [Koyeb](https://www.koyeb.com) (free tier, sem cartão, sem
sleep — ver [ADR-0002](../docs/adr/0002-hospedagem-koyeb.md)): um `koyeb_app` e um
`koyeb_service` que roda a imagem Docker, com a chave da YouTube API guardada como
**secret**.

## Pré-requisitos

- [Terraform](https://developer.hashicorp.com/terraform/install) >= 1.6.
- Um token de API da Koyeb, exportado no ambiente:

  ```sh
  export KOYEB_TOKEN="seu-token"
  ```

- A imagem publicada num registry acessível (ex.: GitHub Container Registry).

## Uso

```sh
cd infra
terraform init

# Crie um arquivo de variáveis a partir do exemplo (NÃO commitado):
cp example.tfvars prod.tfvars   # edite image e youtube_api_key

terraform plan  -var-file=prod.tfvars
terraform apply -var-file=prod.tfvars
```

A URL pública (`https://<app>-<org>.koyeb.app`) aparece no dashboard da Koyeb após
o deploy.

## Variáveis

| Variável | Obrigatória | Default | Descrição |
|---|---|---|---|
| `image` | sim | — | Imagem do container a implantar. |
| `youtube_api_key` | sim | — | Chave da YouTube Data API v3 (vira secret na Koyeb). |
| `app_name` | não | `focuzen` | Nome do app. |
| `service_name` | não | `focuzen` | Nome do service. |
| `region` | não | `was` | Região da Koyeb. |

## Notas

- **Segredos**: `youtube_api_key` é `sensitive` e vira um `koyeb_secret`; nunca vai
  para a imagem. Arquivos `*.tfvars` e o state são gitignored.
- **Health check**: HTTP em `/healthz` na porta 8080.
- **Escala**: uma instância sempre ligada (`min = max = 1`) no tipo `free`.
