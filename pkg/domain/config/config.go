package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// ColorConfig represents the color configuration
type ColorConfig struct {
	Box     string `json:"box"`
	Title   string `json:"title"`
	Default string `json:"default"`
	Empty   string `json:"empty"`
	Types   struct {
		Directory string            `json:"directory"`
		File      string            `json:"file"`
		Image     string            `json:"image"`
		Video     string            `json:"video"`
		Custom    map[string]string `json:"custom"`
	} `json:"types"`
}

// Config represents the user configuration
type Config struct {
	Colors ColorConfig `json:"colors"`
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		Colors: ColorConfig{
			Box:     "#00FFFF", // Cyan
			Title:   "#FFFFFF", // White
			Default: "#FFFFFF", // White
			Empty:   "#FFFF00", // Yellow
			Types: struct {
				Directory string            `json:"directory"`
				File      string            `json:"file"`
				Image     string            `json:"image"`
				Video     string            `json:"video"`
				Custom    map[string]string `json:"custom"`
			}{
				Directory: "#0000FF", // Blue
				File:      "#FFFFFF", // White
				Image:     "#00FF00", // Green
				Video:     "#FF00FF", // Magenta
				Custom:    make(map[string]string),
			},
		},
	}
}

// LoadConfig loads the user configuration from the config file
func LoadConfig() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get user home directory: %v", err)
	}

	configDir := filepath.Join(homeDir, ".config", "q")
	configFile := filepath.Join(configDir, "q.json")

	config := DefaultConfig()

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		if err := os.MkdirAll(configDir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create config directory: %v", err)
		}

		data, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("failed to marshal default config: %v", err)
		}

		if err := os.WriteFile(configFile, data, 0644); err != nil {
			return nil, fmt.Errorf("failed to write default config: %v", err)
		}

		return config, nil
	}

	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	if err := json.Unmarshal(data, config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	return config, nil
}

// SaveConfig saves the configuration to the config file
func SaveConfig(config *Config) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user home directory: %v", err)
	}

	configDir := filepath.Join(homeDir, ".config", "q")
	configFile := filepath.Join(configDir, "q.json")

	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %v", err)
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %v", err)
	}

	if err := os.WriteFile(configFile, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %v", err)
	}

	return nil
}
