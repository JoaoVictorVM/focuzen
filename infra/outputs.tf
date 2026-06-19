output "app_name" {
  description = "Deployed Koyeb app name."
  value       = koyeb_app.focuzen.name
}

output "service_id" {
  description = "Deployed Koyeb service id."
  value       = koyeb_service.focuzen.id
}
