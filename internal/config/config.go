package config

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Theme         string `toml:"theme"`
	Color         bool   `toml:"color"`
	LineNumbers   bool   `toml:"line_numbers"`
	Pager         bool   `toml:"pager"`
	MaxSize       int64  `toml:"max_size"`
	MaxLineLength int    `toml:"max_line_length"`
	HTML          bool   // není v konfiguraci, používá se jen z CLI
}

func Default() *Config {
	return &Config{
		Theme:         "default",
		Color:         true,
		LineNumbers:   false,
		Pager:         true,
		MaxSize:       10 * 1024 * 1024, // 10 MB
		MaxLineLength: 4096,
	}
}

func Load() (*Config, error) {
	cfg := Default()
	configPath := filepath.Join(os.Getenv("HOME"), ".config", "ccat", "config.toml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return cfg, nil
	}
	if _, err := toml.DecodeFile(configPath, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
