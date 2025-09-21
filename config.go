// Copyright 2025 Kristian Whittick
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Configuration structure
type Config struct {
	ArchiveMode   bool     `json:"archiveMode"`
	OutputFolders []string `json:"outputFolders"`
	VerboseMode   bool     `json:"verboseMode"`
	DebugMode     bool     `json:"debugMode"`
}

// Default configuration
func getDefaultConfig() *Config {
	return &Config{
		ArchiveMode:   true,
		OutputFolders: []string{"output"},
		VerboseMode:   false,
		DebugMode:     false,
	}
}

// Load configuration from blendpdf.json
func loadConfig(watchDir string) (*Config, error) {
	configPath := filepath.Join(watchDir, "blendpdf.json")

	// If config file doesn't exist, return default config
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return getDefaultConfig(), nil
	}

	// Read config file
	data, err := os.ReadFile(configPath)
	if err != nil {
		return getDefaultConfig(), err
	}

	// Parse JSON
	config := getDefaultConfig()
	if err := json.Unmarshal(data, config); err != nil {
		return getDefaultConfig(), err
	}

	// Validate config
	if err := validateConfig(config); err != nil {
		return getDefaultConfig(), err
	}

	return config, nil
}

// Validate configuration
func validateConfig(config *Config) error {
	// Ensure at least one output folder
	if len(config.OutputFolders) == 0 {
		config.OutputFolders = []string{"output"}
	}

	return nil
}
