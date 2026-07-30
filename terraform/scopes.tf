# Scope definitions live under scopes/, split one file per domain to
# mirror the Go package layout (internal/event, internal/category, ...).
module "scopes" {
  source = "./scopes"

  resource_server_identifier = data.auth0_resource_server.evently_api.identifier
}
