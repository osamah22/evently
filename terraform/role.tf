# The single Superuser role, bundling every scope defined under scopes/.
# There is currently one manually managed superuser account; see
# user_roles.tf for how it's assigned this role.
#
# NOTE: when a new scope is added under scopes/, add a matching
# `permissions` block below by hand — it is not derived automatically,
# so the two can drift if you forget.
resource "auth0_role" "superuser" {
  name        = "Superuser"
  description = "Holds every Evently permission. Assigned to the single manually managed superuser account."
}

resource "auth0_role_permissions" "superuser" {
  role_id = auth0_role.superuser.id

  permissions {
    name                       = "event:read:any"
    resource_server_identifier = var.resource_server_identifier
  }

  permissions {
    name                       = "event:create:any"
    resource_server_identifier = var.resource_server_identifier
  }

  permissions {
    name                       = "event:update:any"
    resource_server_identifier = var.resource_server_identifier
  }

  permissions {
    name                       = "event:delete:any"
    resource_server_identifier = var.resource_server_identifier
  }

  permissions {
    name                       = "category:create"
    resource_server_identifier = var.resource_server_identifier
  }

  permissions {
    name                       = "category:delete"
    resource_server_identifier = var.resource_server_identifier
  }

  permissions {
    name                       = "user:read"
    resource_server_identifier = var.resource_server_identifier
  }

  # Auth0 requires the scopes to exist on the resource server before
  # they can be granted to a role.
  depends_on = [module.scopes]
}
