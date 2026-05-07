package logger

import (
	"fmt"
	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, fmt.Errorf("не удалось начать логгирование: %w", err)
	}

	return logger, nil
}
