variable "auth0_domain" {
  type        = string
  description = "Auth0 tenant domain (e.g. evently-dev.us.auth0.com), used only by the dedicated Terraform M2M app."
}

variable "auth0_client_id" {
  type        = string
  description = "Client ID of the dedicated Terraform Management API M2M application."
}

variable "auth0_client_secret" {
  type        = string
  description = "Client secret of the dedicated Terraform Management API M2M application."
  sensitive   = true
}

variable "resource_server_identifier" {
  type        = string
  description = "Identifier (audience) of the existing Evently API resource server in Auth0, e.g. https://api.evently.example.com."
}

variable "superuser_user_id" {
  type        = string
  description = "Auth0 user_id (e.g. auth0|64f1a2b3c4d5e6f7a8b9c0d1) of the existing, manually managed superuser account."
}
