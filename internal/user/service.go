package user

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/osamah22/evently/internal/models"
	apperr "github.com/osamah22/evently/pkg/apperror"
)

type UserService struct {
	q *models.Queries
}

func NewService(querier *models.Queries) *UserService {
	if querier == nil {
		panic("querier is null!")
	}
	return &UserService{
		q: querier,
	}
}

// EnsureUserExists(ctx context.Context, auth0ID string) (datab.User, error)
func (u *UserService) EnsureUserExists(ctx context.Context, auth0ID string) (models.User, error) {
	user, err := u.q.GetUserByAuth0ID(ctx, auth0ID)
	if err == nil {
		return user, nil
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return models.User{}, apperr.FromPgError(err, "user")
	}
	created, err := u.q.CreateUser(ctx, auth0ID)
	if err != nil {
		appError := apperr.FromPgError(err, "user")
		if appError.Code == apperr.CodeConflict {
			return u.q.GetUserByAuth0ID(ctx, auth0ID)
		}
		return models.User{}, appError
	}

	return created, nil
}
