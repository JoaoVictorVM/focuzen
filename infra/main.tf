terraform {
  required_version = ">= 1.6"

  required_providers {
    koyeb = {
      source  = "koyeb/koyeb"
      version = "~> 0.1"
    }
  }
}

# Authenticates via the KOYEB_TOKEN environment variable.
provider "koyeb" {}

# The API key is stored as a Koyeb secret and referenced by the service env,
# so it never ends up in the image or in plain text on the instance.
resource "koyeb_secret" "youtube_api_key" {
  name  = "youtube-api-key"
  value = var.youtube_api_key
}

resource "koyeb_app" "focuzen" {
  name = var.app_name
}

resource "koyeb_service" "focuzen" {
  app_name = koyeb_app.focuzen.name

  definition {
    name    = var.service_name
    regions = [var.region]

    # Free tier: a single always-on instance (no sleep) — see ADR-0002.
    instance_types {
      type = "free"
    }

    scalings {
      min = 1
      max = 1
    }

    ports {
      port     = 8080
      protocol = "http"
    }

    routes {
      path = "/"
      port = 8080
    }

    health_checks {
      http {
        port = 8080
        path = "/healthz"
      }
    }

    env {
      key   = "PORT"
      value = "8080"
    }

    env {
      key    = "YOUTUBE_API_KEY"
      secret = koyeb_secret.youtube_api_key.name
    }

    docker {
      image = var.image
    }
  }
}
