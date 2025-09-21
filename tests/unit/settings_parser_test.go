package main_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/johanneslosch/erstatte/pkg/parser"
	"github.com/johanneslosch/erstatte/pkg/settings"
)

func TestLoadSettings_FileExists(t *testing.T) {
	// Create a temporary settings file
	tempDir := t.TempDir()
	settingsPath := filepath.Join(tempDir, "test_erstatte.json")

	expectedSettings := settings.Settings{
		Boolean: []settings.InternalSetting{
			{
				FilePath: "./test.ts",
				Field:    "testField",
				Value:    "true",
			},
		},
	}

	// Write test settings to file
	data, err := json.MarshalIndent(expectedSettings, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal test settings: %v", err)
	}

	err = os.WriteFile(settingsPath, data, 0600)
	if err != nil {
		t.Fatalf("Failed to write test settings file: %v", err)
	}

	// Test loading the settings
	loadedSettings, err := settings.LoadSettings(settingsPath)
	if err != nil {
		t.Fatalf("LoadSettings failed: %v", err)
	}

	// Verify the loaded settings
	if len(loadedSettings.Boolean) != 1 {
		t.Errorf("Expected 1 boolean setting, got %d", len(loadedSettings.Boolean))
	}

	if loadedSettings.Boolean[0].FilePath != "./test.ts" {
		t.Errorf("Expected FilePath './test.ts', got '%s'", loadedSettings.Boolean[0].FilePath)
	}
}

func TestParseAndReplace_JSON(t *testing.T) {
	tempDir := t.TempDir()
	jsonFile := filepath.Join(tempDir, "test.json")

	// Create test JSON file
	testData := map[string]interface{}{
		"debugMode":     false,
		"maxRetries":    3,
		"serverUrl":     "localhost",
		"enableFeature": true,
	}

	data, err := json.MarshalIndent(testData, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal test JSON: %v", err)
	}

	err = os.WriteFile(jsonFile, data, 0600)
	if err != nil {
		t.Fatalf("Failed to write test JSON file: %v", err)
	}

	// Test replacing values
	replacements := map[string]string{
		"debugMode":  "true",
		"maxRetries": "5",
		"serverUrl":  "production.example.com",
	}

	err = parser.ParseAndReplace(jsonFile, replacements)
	if err != nil {
		t.Fatalf("ParseAndReplace failed: %v", err)
	}

	// Verify changes
	content, err := os.ReadFile(jsonFile)
	if err != nil {
		t.Fatalf("Failed to read modified JSON file: %v", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(content, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal modified JSON: %v", err)
	}

	if result["debugMode"] != "true" {
		t.Errorf("Expected debugMode to be 'true', got %v", result["debugMode"])
	}
}
