# Terraform: Auth0 authorization config

Manages Evently's Auth0 authorization surface: API scopes, the
"Superuser" role, and its assignment to the one superuser account. It
does **not** manage the tenant, the SPA application, or the resource
server (API) itself — those already exist and are referenced, not
recreated.

## Layout

```
terraform/
  provider.tf          # provider + terraform block, resource_server data source
  variables.tf          # inputs (domain, M2M creds, API identifier, superuser user_id)
  scopes.tf              # wires up the scopes/ module
  scopes/
    variables.tf
    event.tf             # event:* scopes
    category.tf           # category:* scopes
  role.tf                # Superuser role + its permissions
  user_roles.tf            # assigns Superuser to the one superuser account
  outputs.tf
  terraform.tfvars.example
```

Scopes are managed with `auth0_resource_server_scope` (one resource per
scope). Do not add the deprecated inline `scopes` block to an
`auth0_resource_server` resource for the same API, and don't mix in
`auth0_resource_server_scopes` (the bulk variant) — combining either
with `auth0_resource_server_scope` makes Terraform fight itself and
delete/recreate scopes on every apply.

## One-time manual setup (Auth0 Dashboard)

Terraform authenticates as its own Machine-to-Machine application —
never reuse the SPA's credentials. Before running `terraform init`:

1. **Create an M2M application** for Terraform: Dashboard → Applications
   → Create Application → Machine to Machine.
2. **Authorize it for the Auth0 Management API** and grant these scopes
   (least privilege for what this config manages):
   - `read:resource_servers`, `update:resource_servers`
   - `read:roles`, `create:roles`, `update:roles`, `delete:roles`
   - `read:role_members`, `create:role_members`, `delete:role_members`
   - `read:users`
3. **Note the existing API's identifier** (audience) — used for
   `resource_server_identifier`, and the existing SPA stays untouched.
4. **Find the superuser account's `user_id`** (Dashboard → User
   Management → Users → the account → copy the `user_id`, e.g.
   `auth0|64f1a2b3c4d5e6f7a8b9c0d1`) — used for `superuser_user_id`.

## Configuration

Copy the example vars file and fill in the non-secret values:

```bash
cp terraform.tfvars.example terraform.tfvars
```

Never put `auth0_client_id` / `auth0_client_secret` in a `.tf` or
`.tfvars` file. Export them as environment variables instead (Terraform
picks up `TF_VAR_*` automatically):

```bash
export TF_VAR_auth0_client_id="the Terraform M2M app's client id"
export TF_VAR_auth0_client_secret="the Terraform M2M app's client secret"
```

## Usage

```bash
terraform init
terraform plan
terraform apply
```

If the resource server was previously created outside Terraform, no
import is needed for it — it's only read via a `data` source. If a
"Superuser" role or the scopes already exist by hand in the dashboard,
import them once instead of letting `apply` create duplicates:

```bash
terraform import auth0_role.superuser <role_id>
terraform import auth0_resource_server_scope.event_create_own "<resource_server_identifier>::event:create:own"
# ...repeat per existing scope, using the API's identifier (audience), not its internal id
```

## Adding a new scope

1. Add a `resource "auth0_resource_server_scope" "..."` block in the
   relevant file under `scopes/` (or a new file for a new domain).
2. If the Superuser role should have it too, add a matching
   `permissions { }` block in `role.tf`. This is manual by design (see
   the note in `role.tf`) — the two lists aren't derived from each
   other, so double-check they stay in sync.

## Caveat: `auth0_user_roles` owns the whole role list

`auth0_user_roles` manages the *entire* set of roles on a user, not
just the ones Terraform adds. If the superuser account ever picks up
another role by hand in the dashboard, the next `terraform apply` will
remove it. Keep all of that user's role assignments in this config.
