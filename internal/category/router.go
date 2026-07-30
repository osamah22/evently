package category

import (
	"github.com/gofiber/fiber/v3"
	"github.com/osamah22/evently/internal/auth"
)

func NewCategoryGroup(app *fiber.App,
	controller *CategoryController,
	authmw *auth.AuthMiddleware) {

	categoryGroup := app.Group("/categories")
	// public routes
	categoryGroup.Get("/", controller.list)

	// protected routes
	protected := categoryGroup.Group("/")
	protected.Use(authmw.ValidateToken)
	protected.Post("/",
		authmw.RequireAnyPerm(permCreate),
		controller.create)
	protected.Delete("/:id",
		authmw.RequireAnyPerm(permDelete),
		controller.delete)
}
