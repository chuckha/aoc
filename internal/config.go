package internal

import (
	"os"
	"path/filepath"
	"strings"

	"gitlab.com/tozd/go/errors"
)

type Config struct {
	ConfigDir string
}

func NewConfig() Config {
	home := os.Getenv("HOME")
	defaultConfigDir := filepath.Join(home, ".aoc")
	return Config{
		ConfigDir: defaultConfigDir,
	}
}

func (c Config) GetToken() (string, error) {
	tkn, err := os.ReadFile(filepath.Join(c.ConfigDir, "token"))
	if err != nil {
		return "", errors.WithStack(err)
	}
	return strings.TrimSpace(string(tkn)), nil
}
