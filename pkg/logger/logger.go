package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var log *zap.SugaredLogger

func InitLogger() {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	writeSyncer := zapcore.AddSync(file)
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)

	log = zap.New(core).Sugar()
}

func L() *zap.SugaredLogger {
	return log
}
