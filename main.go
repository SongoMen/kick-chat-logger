package main

import (
	"github.com/SongoMen/kickchatlogger/chatlogger"
	"go.uber.org/zap"
)

func main() {
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()
	sugar := zapLogger.Sugar()
	chatlogger.StartLogger(sugar)
}
