package config

import (
	"dresser/internal/infrastructure/db"
	"encoding/json"
	"os"
	"path/filepath"
)

// Config is the application configuration
type Config struct {
	Database db.Config `json:"database"`
}

// Load loads configuration from the specified file path
func Load(configPath string) (*Config, error) {
	// Read config file
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	// Parse JSON
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// LoadDefault loads the default configuration from config.json
func LoadDefault() (*Config, error) {
	// Get executable directory
	execDir, err := os.Executable()
	if err != nil {
		return nil, err
	}

	// Use config.json in the same directory
	configPath := filepath.Join(filepath.Dir(execDir), "config.json")

	// Try to load from the path
	config, err := Load(configPath)
	if err != nil {
		// Fallback to current directory if not found
		config, err = Load("config.json")
		if err != nil {
			return nil, err
		}
	}

	return config, nil
}
