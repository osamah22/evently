package logger

import "go.uber.org/zap"

type LoggerMW struct {
	baseLogger *zap.Logger
}

// New builds the application's production zap logger.
func NewWithLogger(logger *zap.Logger) *LoggerMW {
	return &LoggerMW{
		baseLogger: logger,
	}
}
