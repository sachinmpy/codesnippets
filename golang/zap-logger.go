package golang

import (
	"os"

	zap "go.uber.org/zap"
)

func initZapLogger(filename string) *zap.Logger { // Logger to save log to a file name

	config := zap.NewProductionConfig()
	logfile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	config.OutputPaths = []string{logfile.Name()}
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	l, err := config.Build()
	if err != nil {
		panic(err)
	}

	return l

}
func initZapLoggerConsole() *zap.Logger { // Logger to log on console
	loggerConfig := zap.NewDevelopmentConfig()

	// loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("")
	loggerConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	l, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}
	return l

}
