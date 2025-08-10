package log

import (
	"log/slog"
	"os"
	"strings"

	"github.com/spf13/viper"
)

const (
	keyValueEncoding string = "key-value"
)

func InitLog() {
	opts := &slog.HandlerOptions{Level: getLevel()}

	var handler slog.Handler = slog.NewJSONHandler(os.Stdout, opts)
	if viper.GetString("log.encoding") == keyValueEncoding {
		handler = slog.NewTextHandler(os.Stderr, opts)
	}

	slog.SetDefault(slog.New(handler))
}

func getLevel() slog.Level {
	logLevelConfig := strings.ToUpper(os.Getenv("LOG_LEVEL"))
	if logLevelConfig == "DEBUG" {
		return slog.LevelDebug
	}

	if logLevelConfig == "INFO" {
		return slog.LevelInfo
	}

	if logLevelConfig == "WARN" {
		return slog.LevelWarn
	}

	if logLevelConfig == "ERROR" {
		return slog.LevelError
	}

	return slog.LevelInfo
}
