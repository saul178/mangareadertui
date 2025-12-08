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
	CollectionPath []string `json:"collection_path"`
}

// might not need this not sure?
func DefaultConfig() *TuiConfig {
	return &TuiConfig{CollectionPath: make([]string, 0)}
}

// TODO: this is bugged and exits out early before even creating a directory, causing it to fail
func isConfigExist() (bool, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return false, fmt.Errorf("couldnt find users home directory: %v\n", err)
	}

	targetPath := filepath.Join(homeDir, tuiConfigPath)

	info, err := os.Stat(targetPath)
	if os.IsNotExist(err) {
		return false, fmt.Errorf("directory does not exist: %v\n", os.ErrNotExist)
	}

	if err != nil {
		return false, fmt.Errorf("far worse error: %v", err)
	}

	if !info.IsDir() {
		return false, fmt.Errorf("path exist but is not a valid directory: %v\n", err)
	}
	return true, nil
}

func createConfigFile(path string) error {
	if err := os.MkdirAll(path, 0o755); err != nil {
		return fmt.Errorf("failed to create configuration path: %v", err)
	}
	fullPath := filepath.Join(path + configFile)

	var config TuiConfig
	jsonBytes, err := json.MarshalIndent(&config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to encode json: %v", err)
	}

	if err := os.WriteFile(fullPath, jsonBytes, 0o755); err != nil {
		return fmt.Errorf("failed to write to json: %v", err)
	}

	return nil
}

func GetConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error %v: ", err)
	}

	configDir := filepath.Join(homeDir, tuiConfigPath)
	configFilePath := filepath.Join(configDir, configFile)

	exist, err := isConfigExist()
	if err != nil {
		return "", err
	}

	if !exist {
		// make the directory and config file
		createConfigFile(configDir)
		return configFilePath, nil
	}
	return configFilePath, nil
}
