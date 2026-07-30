# Scopes for the "event" domain (mirrors internal/event).
#
# There is no event:*:own scope: any authenticated user is implicitly
# allowed to create/update/delete their own events — that's enforced by
# comparing ownership in application code, not gated by an Auth0
# permission. Only the elevated ":any" actions (acting on someone else's
# event) are real, checked permissions.

# Matches internal/event/permissions.go's permReadAny constant.
resource "auth0_resource_server_scope" "event_read_any" {
  resource_server_identifier = var.resource_server_identifier
  scope                      = "event:read:any"
  description                = "Read any events, including DRAFT ones."
}

resource "auth0_resource_server_scope" "event_create_any" {
  resource_server_identifier = var.resource_server_identifier
  scope                      = "event:create:any"
  description                = "Create an event on behalf of any user (admin)."
}

resource "auth0_resource_server_scope" "event_update_any" {
  resource_server_identifier = var.resource_server_identifier
  scope                      = "event:update:any"
  description                = "Update any event regardless of ownership (admin)."
}

resource "auth0_resource_server_scope" "event_delete_any" {
  resource_server_identifier = var.resource_server_identifier
  scope                      = "event:delete:any"
  description                = "Delete any event regardless of ownership (admin)."
}
