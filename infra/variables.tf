variable "app_name" {
  description = "Koyeb app name."
  type        = string
  default     = "focuzen"
}

variable "service_name" {
  description = "Koyeb service name."
  type        = string
  default     = "focuzen"
}

variable "image" {
  description = "Container image to deploy (e.g. ghcr.io/joaovictorvm/focuzen:latest)."
  type        = string
}

variable "region" {
  description = "Koyeb region to deploy to."
  type        = string
  default     = "was"
}

variable "youtube_api_key" {
  description = "YouTube Data API v3 key, stored as a Koyeb secret (never committed)."
  type        = string
  sensitive   = true
}
