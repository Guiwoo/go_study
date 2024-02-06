package gin_log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var GinLogger *zap.Logger

func NewZapLog() {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "/Users/guiwoopark/Desktop/personal/study/gin_pra/logs/gin_pra.log",
		MaxSize:    100,
		MaxBackups: 3,
		MaxAge:     7, // days
	})

	pe := zap.NewProductionEncoderConfig()

	pe.EncodeTime = zapcore.RFC3339TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(pe)

	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		zap.InfoLevel,
	)
	consoleCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zap.InfoLevel)
	tees := zapcore.NewTee(fileCore, consoleCore)
	GinLogger = zap.New(tees, zap.AddCaller())

	GinLogger.Info("something work?")
}
