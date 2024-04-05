package database

import (
	"time"

	"github.com/omegaatt36/film36exp/logging"
	"gorm.io/gorm/logger"
)

func newGormLogger() logger.Interface {
	slowThreshold := 100 * time.Millisecond
	cfg := logger.Config{
		SlowThreshold:             slowThreshold,
		LogLevel:                  logger.Warn,
		IgnoreRecordNotFoundError: true,
		Colorful:                  true,
	}

	return logger.New(
		&gormLogger{},
		cfg,
	)
}

type gormLogger struct{}

func (l *gormLogger) Printf(message string, params ...any) {
	logging.Infof(message, params...)
}
