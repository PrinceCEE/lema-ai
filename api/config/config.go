package config

import (
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog"
)

const (
	EnvDevelopment = "development"
	EnvTest        = "test"
	EnvProduction  = "production"
)

type Config struct {
	PORT              string
	DSN               string
	ENV               string
	LOG_LEVEL         string
	MAX_IDLE_CONNS    int
	MAX_OPEN_CONNS    int
	CONN_MAX_LIFETIME time.Duration
}

func NewConfig(env, loglevel string) *Config {
	return &Config{
		PORT:              getEnv("PORT", "5001"),
		DSN:               getEnv("DSN", "file::memory:?cache=shared"),
		ENV:               getEnv("ENV", env),
		LOG_LEVEL:         getEnv("LOG_LEVEL", loglevel),
		MAX_IDLE_CONNS:    getEnvAsInt("MAX_IDLE_CONNS", 10),
		MAX_OPEN_CONNS:    getEnvAsInt("MAX_OPEN_CONNS", 100),
		CONN_MAX_LIFETIME: getEnvAsDuration("CONN_MAX_LIFETIME", time.Hour),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(name string, defaultVal int) int {
	if value, ok := os.LookupEnv(name); ok {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultVal
}

func getEnvAsDuration(name string, defaultVal time.Duration) time.Duration {
	if value, ok := os.LookupEnv(name); ok {
		if durationValue, err := time.ParseDuration(value); err == nil {
			return durationValue
		}
	}
	return defaultVal
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
		return zerolog.InfoLevel
	}
}
