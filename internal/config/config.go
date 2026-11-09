package config

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// Config contains the small, explicit set of runtime settings.
type Config struct {
	Workspace string
	LogLevel  string
}

func Load() (Config, error) {
	workspace := strings.TrimSpace(os.Getenv("ORBITA_WORKSPACE"))
	if workspace == "" {
		workspace = "./workspace"
	}
	abs, err := filepath.Abs(workspace)
	if err != nil {
		return Config{}, err
	}
	level := strings.ToLower(strings.TrimSpace(os.Getenv("ORBITA_LOG_LEVEL")))
	if level == "" {
		level = "info"
	}
	switch level {
	case "debug", "info", "warn", "error":
	default:
		return Config{}, errors.New("ORBITA_LOG_LEVEL must be debug, info, warn, or error")
	}
	return Config{Workspace: abs, LogLevel: level}, nil
}
