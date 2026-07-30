# Scopes for the "category" domain (mirrors internal/category).

resource "auth0_resource_server_scope" "category_create" {
  resource_server_identifier = var.resource_server_identifier
  scope                      = "category:create"
  description                = "Create a category."
}

resource "auth0_resource_server_scope" "category_delete" {
  resource_server_identifier = var.resource_server_identifier
  scope                      = "category:delete"
  description                = "Delete a category."
}
