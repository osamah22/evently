package logger

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ctxKey struct{}

var loggerKey = ctxKey{}

var fallback *zap.Logger

// New creates the application's base logger and stores it as the fallback.
func New() (*zap.Logger, error) {
	log, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	fallback = log

	return log, nil
}

// Middleware creates a request-scoped logger and stores it in the request context.
func Middleware(base *zap.Logger) fiber.Handler {
	return func(c fiber.Ctx) error {
		requestID := uuid.New().String()

		log := base.With(
			zap.String("request_id", requestID),
		)

		ctx := context.WithValue(
			c.Context(),
			loggerKey,
			log,
		)

		c.SetContext(ctx)

		return c.Next()
	}
}

// FromContext returns the request-scoped logger.
// If no logger exists, the application's base logger is returned.
func FromContext(ctx context.Context) *zap.Logger {
	if log, ok := ctx.Value(loggerKey).(*zap.Logger); ok {
		return log
	}

	return fallback
}

// WithFields returns a new context containing a logger enriched with additional fields.
func WithFields(ctx context.Context, fields ...zap.Field) context.Context {
	log := FromContext(ctx)

	log = log.With(fields...)

	return context.WithValue(ctx, loggerKey, log)
}
