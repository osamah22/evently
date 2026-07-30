terraform {
  required_version = ">= 1.5.0"

  required_providers {
    auth0 = {
      source  = "auth0/auth0"
      version = "~> 1.0"
    }
  }
}

# Authenticates as a dedicated Machine-to-Machine application, not the
# SPA used by end users. That M2M app must be authorized for the Auth0
# Management API (see terraform/README.md for the required scopes).
provider "auth0" {
  domain        = var.auth0_domain
  client_id     = var.auth0_client_id
  client_secret = var.auth0_client_secret
}

# The API (resource server) already exists in the tenant and is managed
# by hand; we only look it up so its identifier can be passed down to
# the scopes module and the Superuser role.
data "auth0_resource_server" "evently_api" {
  identifier = var.resource_server_identifier
}
