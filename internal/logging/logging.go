package logging

import (
	"context"

	"github.com/joematpal/go-logger"
)

type Logger = logger.CorrelationLogger

type Debugger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
}

type Field = logger.Field

func GetCorrelationIDFromMetadata(ctx context.Context) (string, error) {
	return logger.GetCorrelationIDFromMetadata(ctx)
}

type KV = logger.KV

type Map = logger.Map
