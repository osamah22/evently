output "superuser_role_id" {
  value       = auth0_role.superuser.id
  description = "Auth0 role_id of the Superuser role, for reference/debugging."
}

output "resource_server_id" {
  value       = data.auth0_resource_server.evently_api.id
  description = "Auth0 resource server ID resolved from resource_server_identifier."
}
