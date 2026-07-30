package auth

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gofiber/fiber/v3"
	"github.com/osamah22/evently/config"
	"github.com/osamah22/evently/internal/models"
	apperr "github.com/osamah22/evently/pkg/apperror"
	"go.uber.org/zap"
)

type UserProvisioner interface {
	EnsureUserExists(ctx context.Context, auth0ID string) (models.User, error)
}
type AuthMiddleware struct {
	config      config.EnvVars
	validator   *validator.Validator
	logger      *zap.Logger
	userService UserProvisioner
}

func NewAuthMiddleware(cfg config.EnvVars, logger *zap.Logger, userService UserProvisioner) *AuthMiddleware {
	issuerURL, err := url.Parse("https://" + cfg.AUTH0_DOMAIN + "/")
	if err != nil {
		panic(err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{cfg.AUTH0_AUDIENCE},
		validator.WithCustomClaims(func() validator.CustomClaims {
			return &customClaims{}
		}),
	)
	if err != nil {
		panic(err)
	}
	logger.Named("auth_middleware")

	return &AuthMiddleware{config: cfg, validator: jwtValidator, logger: logger, userService: userService}

}

func (a *AuthMiddleware) ValidateToken(c fiber.Ctx) error {

	// get the token from the request header
	authHeader := c.Get("Authorization")
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid authorization header",
		})
	}

	// Validate the token
	tokenInfo, err := a.validator.ValidateToken(c.Context(), authHeaderParts[1])
	if err != nil {
		fmt.Println(err)
		a.logger.Error("invalid token", zap.String("error", err.Error()))
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	claims, ok := tokenInfo.(*validator.ValidatedClaims)
	if !ok {
		a.logger.Error("failed to parse tokenInfo to claims")
		return apperr.Internal(err)
	}
	auth0ID := claims.RegisteredClaims.Subject
	dbUser, err := a.userService.EnsureUserExists(c.Context(), auth0ID)
	if err != nil {
		return apperr.Internal(err)
	}
	user := &AuthenticatedUser{
		auth0ID:     dbUser.Auth0ID,
		ID:          dbUser.ID,
		Permissions: map[string]struct{}{},
	}
	custom, ok := claims.CustomClaims.(*customClaims)
	if ok {
		for _, perm := range custom.Permissions {
			user.Permissions[perm] = struct{}{}
		}
	}
	setUser(c, user)

	// Go to next middleware:
	return c.Next()
}

type Perm string

func (a *AuthMiddleware) RequireAnyPerm(perms ...Perm) fiber.Handler {
	return func(c fiber.Ctx) error {
		user, ok := GetUser(c)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "unauthorized",
			})
		}

		for _, perm := range perms {
			if user.HasPermission(perm) {
				return c.Next()
			}
		}
		a.logger.Warn("permission denied",
			zap.String("user_id", user.ID.String()),
			zap.Strings("required_any_of", permsToStrings(perms)),
		)
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "forbidden",
		})
	}
}
func permsToStrings(perms []Perm) []string {
	out := make([]string, len(perms))
	for i, p := range perms {
		out[i] = string(p)
	}
	return out
}
