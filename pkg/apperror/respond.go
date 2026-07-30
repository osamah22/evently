package apperr

import (
	"errors"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/osamah22/evently/pkg/logger"
	"go.uber.org/zap"
)

func Respond(c fiber.Ctx, err error) error {
	logger := logger.FromContext(c.Context())
	if appErr, ok := errors.AsType[*AppError](err); ok {
		if appErr.Internal {
			logger.Error("internal error",
				zap.Error(appErr.Err),
				zap.String("code", string(appErr.Code)),
				zap.String("path", c.Path()),
				zap.String("method", c.Method()),
			)
			return c.Status(appErr.Status).JSON(fiber.Map{
				"code":    CodeInternal,
				"message": "internal server error", // never leak appErr.Err.Error() to the client
			})
		}

		logger.Info(string(appErr.Code), zap.String("error", appErr.Error()))
		return c.Status(appErr.Status).JSON(fiber.Map{
			"code":    appErr.Code,
			"message": appErr.Message,
		})
	}
	if fiberErr, ok := errors.AsType[*fiber.Error](err); ok {
		log.Debug("fiber error", zap.Int("status", fiberErr.Code), zap.String("path", c.Path()))
		return c.Status(fiberErr.Code).JSON(fiber.Map{
			"code":    CodeNotFound,
			"message": fiberErr.Message,
		})
	}
	// unknown/unwrapped error — treat as internal, always safe default

	logger.Error("unexpected error", zap.Error(err))
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"code":    CodeInternal,
		"message": "internal server error",
	})
}
