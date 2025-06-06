package config

import (
	"errors"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	AgentType string `yaml:"agent_type"`
	APIKey    string `yaml:"api_key"`
	Language  string `yaml:"language"`
}

// ConfigPath allows overriding the default configuration file location.
// Tests can set this value to write config files in temporary directories.
var ConfigPath string

func Load() (*Config, error) {
	path := configPath()
	f, err := os.Open(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return &Config{}, nil
		}

		return nil, err
	}

	defer f.Close()

	var config Config
	if err := yaml.NewDecoder(f).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) Save() error {
	path := configPath()
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer f.Close()

	return yaml.NewEncoder(f).Encode(c)
}

func configPath() string {
	if ConfigPath != "" {
		return ConfigPath
	}
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".easycommit.yaml")
}

func loadDefaultConfig() *Config {
	return &Config{
		AgentType: "openai",
		Language:  "en-US",
	}
}
