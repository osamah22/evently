package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/osamah22/evently/contracts"
	"github.com/osamah22/evently/datab"
	"go.uber.org/zap"
)

type EventController struct {
	logger *zap.Logger
	q      datab.Querier
}

func NewEventController(logger *zap.Logger, queries datab.Querier) *EventController {
	return &EventController{
		logger: logger,
		q:      queries,
	}
}

func (c *EventController) RegisterRoutes(router gin.IRouter) {
	router.GET("/events/categories", c.ListCategories)
}

func (c *EventController) CreateEvent(ctx *gin.Context) {

}

// ListCategories godoc
//
//	@Summary		List event categories
//	@Description	Returns all available event categories
//	@Tags			events
//	@Produce		json
//	@Success		200	{object}	contracts.CategoriesResponse
//	@Failure		500	{object}	contracts.ErrorResponse
//	@Router			/events/categories [get]
func (c *EventController) ListCategories(ctx *gin.Context) {
	categories, err := c.q.ListCategories(ctx.Request.Context())
	if err != nil {
		c.logger.Error("failed to list categories", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, contracts.ErrorResponse{Error: err.Error()})
		return
	}

	resp := contracts.CategoriesResponse{Categories: make([]contracts.CategoryResponse, len(categories))}
	for i, cat := range categories {
		resp.Categories[i] = contracts.CategoryResponse{
			ID:        cat.ID,
			Name:      cat.Name,
			CreatedAt: cat.CreatedAt.Time,
		}
	}
	ctx.JSON(http.StatusOK, resp)
}
