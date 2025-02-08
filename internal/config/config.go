package config

import (
	"log/slog"
	"os"
)

// LogConfig stores Logger Configuration
type LogConfig struct {
	Level slog.Level
	Group string
}

func (lc LogConfig) Apply() {
	logger := slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: lc.Level,
		}),
	).WithGroup(lc.Group)

	slog.SetDefault(logger)
}

type Config struct {
	Destination string
	Log         LogConfig
}

func (c Config) Clone() Config {
	return c
}

func (c Config) Apply() {
	c.Log.Apply()
}

var config = &DefaultConfig

func GetConfig() *Config {
	return config
}

var DefaultConfig = Config{
	Destination: "output",
	Log: LogConfig{
		Level: slog.LevelInfo,
		Group: "data",
	},
}
