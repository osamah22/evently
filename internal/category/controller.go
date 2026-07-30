package category

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	apperr "github.com/osamah22/evently/pkg/apperror"
)

type CategoryController struct {
	svc *Service
}

func NewController(categoryService *Service) *CategoryController {
	return &CategoryController{
		svc: categoryService,
	}
}

// list godoc
//
//	@Summary		List categories
//	@Description	List all event categories
//	@Tags			categories
//	@Produce		json
//	@Success		200	{object}	contracts.CategoriesResponse
//	@Failure		500	{object}	apperr.AppError
//	@Router			/categories [get]
func (h *CategoryController) list(c fiber.Ctx) error {
	categories, err := h.svc.List(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"data": categories,
	})
}

// create godoc
//
//	@Summary		Create category
//	@Description	Create a new event category
//	@Tags			categories
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			request	body		createRequest	true	"Category to create"
//	@Success		201		{object}	models.Category
//	@Failure		400		{object}	apperr.AppError
//	@Failure		403		{object}	apperr.AppError
//	@Failure		500		{object}	apperr.AppError
//	@Router			/categories [post]
func (h *CategoryController) create(c fiber.Ctx) error {
	var req createRequest
	if err := c.Bind().Body(&req); err != nil {
		return apperr.FromBindErr(err)
	}

	category, err := h.svc.Create(c.Context(), req.Name)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": category,
	})
}

// delete godoc
//
//	@Summary		Delete category
//	@Description	Delete an event category by id
//	@Tags			categories
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path	int	true	"Category ID"
//	@Success		204
//	@Failure		400	{object}	apperr.AppError
//	@Failure		403	{object}	apperr.AppError
//	@Failure		500	{object}	apperr.AppError
//	@Router			/categories/{id} [delete]
func (h *CategoryController) delete(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil || id <= 0 {
		return apperr.Validation(apperr.CodeValidation, "invalid category id")
	}

	_, err = h.svc.Delete(c.Context(), int(id))
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
