package category

import (
	"context"
	"strings"

	"github.com/osamah22/evently/internal/models"
	apperr "github.com/osamah22/evently/pkg/apperror"
)

type Service struct {
	q *models.Queries
}

func NewService(queries *models.Queries) *Service {
	return &Service{
		q: queries,
	}
}

func (s *Service) List(ctx context.Context) ([]models.Category, error) {
	categories, err := s.q.ListCategories(ctx)
	if err != nil {
		return nil, apperr.FromPgError(err, "category")
	}
	return categories, nil
}

func (s *Service) Create(ctx context.Context, name string) (models.Category, error) {
	category, err := s.q.CreateCategory(ctx,
		strings.TrimSpace(strings.ToUpper(name)))
	if err != nil {
		return models.Category{}, apperr.FromPgError(err, "category")
	}
	return category, nil
}

func (s *Service) Delete(ctx context.Context, id int) (models.Category, error) {
	category, err := s.q.DeleteById(ctx, int32(id))
	if err != nil {
		return models.Category{}, apperr.FromPgError(err, "category")
	}
	return category, nil
}
