package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func L() *zap.Logger {
	return zap.L()
}

func S() *zap.SugaredLogger {
	return zap.S()
}

func New() (*zap.Logger, error) {
	core, err := getCore()
	if err != nil {
		return nil, err
	}

	logger := zap.New(zapcore.NewTee(core), zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	zap.ReplaceGlobals(logger)

	return logger, nil
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
}

// Hieu, I want you to think about core = writer (aka transportation) + encoder
func getCore() (zapcore.Core, error) {
	writer, err := getFileRotateWriteSyncer()
	if err != nil {
		return nil, err
	}

	return zapcore.NewCore(getEncoder(), writer, zapcore.DebugLevel), nil
}
