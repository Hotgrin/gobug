package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Config holds the user's own API key for the BYOK AI fallback. Stored
// locally in the OS config dir — never bundled with the repo, never sent
// anywhere except directly to Anthropic's API when the user runs a query.
type Config struct {
	APIKey string `json:"apiKey"`
	Model  string `json:"model"`
}

// configDirOverride lets tests point config storage at a temp dir instead
// of the real OS config dir. Empty in normal operation.
var configDirOverride string

func path() (string, error) {
	dir := configDirOverride
	if dir == "" {
		d, err := os.UserConfigDir()
		if err != nil {
			return "", err
		}
		dir = d
	}
	gobugDir := filepath.Join(dir, "gobug")
	if err := os.MkdirAll(gobugDir, 0700); err != nil {
		return "", err
	}
	return filepath.Join(gobugDir, "config.json"), nil
}

// Load returns the saved config, or a zero-value Config if none exists yet.
func Load() Config {
	p, err := path()
	if err != nil {
		return Config{}
	}
	data, err := os.ReadFile(p)
	if err != nil {
		return Config{}
	}
	var c Config
	_ = json.Unmarshal(data, &c)
	return c
}

// Save writes the config to disk with owner-only permissions (0600) since
// it contains an API key.
func Save(c Config) error {
	p, err := path()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(p, data, 0600)
}
