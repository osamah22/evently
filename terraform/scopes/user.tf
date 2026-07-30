# Scopes for the "user" domain (mirrors internal/user).

resource "auth0_resource_server_scope" "user_read" {
  resource_server_identifier = var.resource_server_identifier
  scope                      = "user:read"
  description                = "Read a user's information."
}
