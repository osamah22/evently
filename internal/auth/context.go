package auth

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type AuthenticatedUser struct {
	auth0ID     string
	ID          uuid.UUID
	Permissions map[string]struct{}
}

const userContextKey = "auth_user"

func setUser(c fiber.Ctx, user *AuthenticatedUser) {
	c.Locals(userContextKey, user)
}

// GetUser retrieves the authenticated user from context.
// Returns nil, false if called on an unprotected route.
func GetUser(c fiber.Ctx) (*AuthenticatedUser, bool) {
	user, ok := c.Locals(userContextKey).(*AuthenticatedUser)
	return user, ok
}

func (u *AuthenticatedUser) HasPermission(perm Perm) bool {
	_, ok := u.Permissions[string(perm)]
	return ok
}
