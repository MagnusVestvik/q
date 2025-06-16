package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// ColorConfig represents the color configuration
type ColorConfig struct {
	Box     string `json:"box"`     // Box border color
	Title   string `json:"title"`   // Title color
	Default string `json:"default"` // Default text color
	Empty   string `json:"empty"`   // Empty directory message color
	Types   struct {
		Directory string `json:"directory"` // Directory color
		File      string `json:"file"`      // Default file color
		Image     string `json:"image"`     // Image file color
		Video     string `json:"video"`     // Video file color
		Custom    map[string]string `json:"custom"` // Custom file type colors
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
	// Get user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get user home directory: %v", err)
	}

	// Construct config file path
	configDir := filepath.Join(homeDir, ".config", "q")
	configFile := filepath.Join(configDir, "q.json")

	// Create default config
	config := DefaultConfig()

	// Check if config file exists
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		// Create config directory if it doesn't exist
		if err := os.MkdirAll(configDir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create config directory: %v", err)
		}

		// Write default config
		data, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("failed to marshal default config: %v", err)
		}

		if err := os.WriteFile(configFile, data, 0644); err != nil {
			return nil, fmt.Errorf("failed to write default config: %v", err)
		}

		return config, nil
	}

	// Read and parse config file
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
	// Get user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user home directory: %v", err)
	}

	// Construct config file path
	configDir := filepath.Join(homeDir, ".config", "q")
	configFile := filepath.Join(configDir, "q.json")

	// Create config directory if it doesn't exist
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %v", err)
	}

	// Marshal config to JSON
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %v", err)
	}

	// Write config file
	if err := os.WriteFile(configFile, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %v", err)
	}

	return nil
} 