package category

import (
	"context"
	"strings"

	"github.com/osamah22/evently/internal/models"
	apperr "github.com/osamah22/evently/pkg/apperror"
)

type repository interface {
	ListCategories(ctx context.Context) ([]models.Category, error)
	CreateCategory(ctx context.Context, name string) (models.Category, error)
	DeleteById(ctx context.Context, id int32) (models.Category, error)
}

type Service struct {
	repo repository
}

func NewService(repo repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) List(ctx context.Context) ([]models.Category, error) {
	categories, err := s.repo.ListCategories(ctx)
	if err != nil {
		return nil, apperr.FromPgError(err, "category")
	}
	return categories, nil
}

func (s *Service) Create(ctx context.Context, name string) (models.Category, error) {
	category, err := s.repo.CreateCategory(ctx,
		strings.TrimSpace(strings.ToUpper(name)))
	if err != nil {
		return models.Category{}, apperr.FromPgError(err, "category")
	}
	return category, nil
}

func (s *Service) Delete(ctx context.Context, id int) (models.Category, error) {
	category, err := s.repo.DeleteById(ctx, int32(id))
	if err != nil {
		return models.Category{}, apperr.FromPgError(err, "category")
	}
	return category, nil
}
