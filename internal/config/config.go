package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// this will be the configuration for the tui program
// TODO: set the configuration for the filetree component to store the users selected paths
const (
	tuiConfigPath = "/.config/mangareadertui"
	configFile    = "config.json"
)

type TuiConfig struct {
	CollectionPath []string            `json:"collection_path"`
	MangaSeries    map[string][]string `json:"manga_series"`
}

func DefaultConfig() *TuiConfig {
	return &TuiConfig{CollectionPath: make([]string, 0)}
}

func getConfigFile() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("couldn't find users home directory:\n%w", err)
	}

	configFilePath := filepath.Join(homeDir, tuiConfigPath, configFile)
	return configFilePath, nil
}

func LoadConfig() (*TuiConfig, error) {
	cfgFile, err := getConfigFile()
	if err != nil {
		return nil, fmt.Errorf("could not locate users config:\n%w", err)
	}
	configDir := filepath.Dir(cfgFile)

	if err := os.MkdirAll(configDir, 0o755); err != nil {
		return nil, fmt.Errorf("failed to create mangareadertui config directory:\n%w", err)
	}

	data, err := os.ReadFile(cfgFile)
	if os.IsNotExist(err) {
		cfg := DefaultConfig()
		if saveErr := SaveConfig(cfg); saveErr != nil {
			return nil, fmt.Errorf("failed to save default config:\n%w", saveErr)
		}
		return cfg, nil
	}

	var tuiCfg TuiConfig
	if err := json.Unmarshal(data, &tuiCfg); err != nil {
		return nil, fmt.Errorf("failed to parse json config file:\n%w", err)
	}
	return &tuiCfg, nil
}

func SaveConfig(cfg *TuiConfig) error {
	cfgFile, err := getConfigFile()
	if err != nil {
		return fmt.Errorf("could not locate users config:\n%w", err)
	}

	jsonBytes, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to encode json:\n%w", err)
	}

	if err := os.WriteFile(cfgFile, jsonBytes, 0o644); err != nil {
		return fmt.Errorf("failed to write to json config file:\n%w", err)
	}

	return nil
}
