package apperr

import (
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// FromPgError translates a raw pgx/postgres error into an *AppError.
// resource is used for friendlier NotFound messages, e.g. "event".
func FromPgError(err error, resource string) *AppError {
	if errors.Is(err, pgx.ErrNoRows) {
		return NotFound(resource)
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505": // unique_violation
			return Conflict(resource + " already exists")
		case "23503": // foreign_key_violation
			// Log the real detail, never leak it to the client.
			return NotFound("referenced resource")
		}
	}

	return Internal(err)
}
