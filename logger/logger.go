package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logs *zap.Logger
var err error

func init() {
	cfg := zap.NewProductionConfig()
	encoderCfg := zap.NewProductionEncoderConfig()

	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.StacktraceKey = ""

	cfg.EncoderConfig = encoderCfg

	logs, err = cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	defer logs.Sync()
}

func Info(msg string) {
	logs.Info(msg)
}

func Error(msg string) {
	logs.Error(msg)
}
