# Assigns the Superuser role to the one existing superuser account.
#
# auth0_user_roles manages the FULL set of roles on the user, not just
# this one — if that account ever picks up other roles by hand in the
# dashboard, this resource will remove them on the next apply.
resource "auth0_user_roles" "superuser" {
  user_id = var.superuser_user_id
  roles   = [auth0_role.superuser.id]
}
