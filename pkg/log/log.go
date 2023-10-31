package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger = NewLazyLogger(func() (*zap.SugaredLogger, error) {
	conf := zap.NewDevelopmentConfig()
	conf.DisableCaller = true
	conf.DisableStacktrace = true
	conf.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	logger, err := conf.Build()
	if err != nil {
		return nil, err
	}

	return logger.Sugar(), nil
})

func SetLogger(l *zap.SugaredLogger) {
	Logger = NewLazyLogger(func() (*zap.SugaredLogger, error) { return l, nil })
}
