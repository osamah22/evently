package user

import (
	"github.com/gofiber/fiber/v3"
	"github.com/osamah22/evently/internal/auth"
	apperr "github.com/osamah22/evently/pkg/apperror"
)

type UserController struct {
	svc *UserService
}

func NewUserController(UserService *UserService) *UserController {
	return &UserController{
		svc: UserService,
	}
}

// profile godoc
//
//	@Summary		Get current user
//	@Description	Return the profile of the authenticated user
//	@Tags			user
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{object}	models.User
//	@Failure		404	{object}	apperr.AppError
//	@Failure		500	{object}	apperr.AppError
//	@Router			/user/me [get]
func (h *UserController) profile(c fiber.Ctx) error {
	user, exists := auth.GetUser(c)
	if !exists {
		return apperr.NotFound("user")
	}
	dbUser, err := h.svc.q.GetUserByID(c.Context(), user.ID)
	if err != nil {
		return apperr.Internal(err)
	}

	return c.JSON(fiber.Map{
		"message": "You are logged in",
		"data":    dbUser,
	})
}
