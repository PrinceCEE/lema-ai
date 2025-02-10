package config

import (
	"errors"
	"os"

	"github.com/rs/zerolog"
)

const (
	EnvDevelopment = "development"
	EnvTest        = "test"
	EnvProduction  = "production"
)

type Config struct {
	PORT      string
	DSN       string
	ENV       string
	LOG_LEVEL string
}

func NewConfig(env, loglevel string) *Config {
	return &Config{
		PORT:      getEnv("PORT", "5001"),
		DSN:       getEnv("DSN", "file::memory:?cache=shared"),
		ENV:       getEnv("ENV", env),
		LOG_LEVEL: getEnv("LOG_LEVEL", loglevel),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func GetLoggerLevel(loglevel string) zerolog.Level {
	switch loglevel {
	case zerolog.LevelTraceValue:
		return zerolog.TraceLevel

	case zerolog.LevelDebugValue:
		return zerolog.DebugLevel

	case zerolog.LevelInfoValue:
		return zerolog.InfoLevel

	case zerolog.LevelWarnValue:
		return zerolog.WarnLevel

	case zerolog.LevelErrorValue:
		return zerolog.ErrorLevel

	case zerolog.LevelFatalValue:
		return zerolog.FatalLevel

	case zerolog.LevelPanicValue:
		return zerolog.PanicLevel
	default:
		panic(errors.New("invalid loglevel value"))
	}
}
