package utils

import (
	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

func InitializeLogger() {
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()
	Logger = zapLogger.Sugar()
}
