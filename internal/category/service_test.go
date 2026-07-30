package category

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/osamah22/evently/internal/models"
	apperr "github.com/osamah22/evently/pkg/apperror"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// mockRepository is a testify/mock test double for the repository interface.
type mockRepository struct {
	mock.Mock
}

func (m *mockRepository) ListCategories(ctx context.Context) ([]models.Category, error) {
	args := m.Called(ctx)
	categories, _ := args.Get(0).([]models.Category)
	return categories, args.Error(1)
}

func (m *mockRepository) CreateCategory(ctx context.Context, name string) (models.Category, error) {
	args := m.Called(ctx, name)
	return args.Get(0).(models.Category), args.Error(1)
}

func (m *mockRepository) DeleteById(ctx context.Context, id int32) (models.Category, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(models.Category), args.Error(1)
}

func TestService_List(t *testing.T) {
	t.Run("returns categories from the repository", func(t *testing.T) {
		want := []models.Category{{ID: 1, Name: "MUSIC"}, {ID: 2, Name: "TECH"}}
		repo := new(mockRepository)
		repo.On("ListCategories", mock.Anything).Return(want, nil)

		svc := NewService(repo)
		got, err := svc.List(context.Background())

		require.NoError(t, err)
		require.Equal(t, want, got)
		repo.AssertExpectations(t)
	})

	t.Run("wraps a DB error as an internal AppError", func(t *testing.T) {
		repo := new(mockRepository)
		repo.On("ListCategories", mock.Anything).Return(nil, errors.New("connection reset"))

		svc := NewService(repo)
		_, err := svc.List(context.Background())

		appErr, ok := errors.AsType[*apperr.AppError](err)
		require.True(t, ok, "expected *apperr.AppError, got %T", err)
		require.Equal(t, apperr.CodeInternal, appErr.Code)
	})
}

func TestService_Create(t *testing.T) {
	t.Run("normalizes the name before storing it", func(t *testing.T) {
		repo := new(mockRepository)
		repo.On("CreateCategory", mock.Anything, "MUSIC").
			Return(models.Category{ID: 1, Name: "MUSIC"}, nil)

		svc := NewService(repo)
		_, err := svc.Create(context.Background(), "  music  ")

		require.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("translates a unique violation into a conflict AppError", func(t *testing.T) {
		repo := new(mockRepository)
		repo.On("CreateCategory", mock.Anything, "MUSIC").
			Return(models.Category{}, &pgconn.PgError{Code: "23505"})

		svc := NewService(repo)
		_, err := svc.Create(context.Background(), "music")

		appErr, ok := errors.AsType[*apperr.AppError](err)
		require.True(t, ok, "expected *apperr.AppError, got %T", err)
		require.Equal(t, apperr.CodeConflict, appErr.Code)
	})
}

func TestService_Delete(t *testing.T) {
	t.Run("returns the deleted category", func(t *testing.T) {
		repo := new(mockRepository)
		repo.On("DeleteById", mock.Anything, int32(7)).
			Return(models.Category{ID: 7, Name: "MUSIC"}, nil)

		svc := NewService(repo)
		got, err := svc.Delete(context.Background(), 7)

		require.NoError(t, err)
		require.EqualValues(t, 7, got.ID)
		repo.AssertExpectations(t)
	})

	t.Run("translates a missing row into a not-found AppError", func(t *testing.T) {
		repo := new(mockRepository)
		repo.On("DeleteById", mock.Anything, int32(7)).
			Return(models.Category{}, pgx.ErrNoRows)

		svc := NewService(repo)
		_, err := svc.Delete(context.Background(), 7)

		appErr, ok := errors.AsType[*apperr.AppError](err)
		require.True(t, ok, "expected *apperr.AppError, got %T", err)
		require.Equal(t, apperr.CodeNotFound, appErr.Code)
	})
}
