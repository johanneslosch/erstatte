package settings_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/johanneslosch/erstatte/pkg/settings"
)

func TestLoadSettings_FileExists(t *testing.T) {
	// Create a temporary settings file
	tempDir := t.TempDir()
	settingsPath := filepath.Join(tempDir, "test_settings.json")
	
	expectedSettings := settings.Settings{
		Boolean: []settings.BooleanSetting{
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
	
	err = os.WriteFile(settingsPath, data, 0644)
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
	
	if loadedSettings.Boolean[0].Field != "testField" {
		t.Errorf("Expected Field 'testField', got '%s'", loadedSettings.Boolean[0].Field)
	}
	
	if loadedSettings.Boolean[0].Value != "true" {
		t.Errorf("Expected Value 'true', got '%s'", loadedSettings.Boolean[0].Value)
	}
}

func TestLoadSettings_FileDoesNotExist(t *testing.T) {
	tempDir := t.TempDir()
	settingsPath := filepath.Join(tempDir, "nonexistent_settings.json")
	
	// Test loading non-existent settings file
	loadedSettings, err := settings.LoadSettings(settingsPath)
	if err != nil {
		t.Fatalf("LoadSettings failed: %v", err)
	}
	
	// Should create default settings
	if len(loadedSettings.Boolean) != 1 {
		t.Errorf("Expected 1 default boolean setting, got %d", len(loadedSettings.Boolean))
	}
	
	// Verify file was created
	if _, err := os.Stat(settingsPath); os.IsNotExist(err) {
		t.Error("Settings file should have been created")
	}
}

func TestSaveSettings(t *testing.T) {
	tempDir := t.TempDir()
	settingsPath := filepath.Join(tempDir, "save_test_settings.json")
	
	testSettings := settings.Settings{
		Boolean: []settings.BooleanSetting{
			{
				FilePath: "./example.ts",
				Field:    "debugMode",
				Value:    "false",
			},
			{
				FilePath: "./config.json",
				Field:    "enableFeature",
				Value:    "true",
			},
		},
	}
	
	// Test saving settings
	err := settings.SaveSettings(settingsPath, testSettings)
	if err != nil {
		t.Fatalf("SaveSettings failed: %v", err)
	}
	
	// Verify file was created
	if _, err := os.Stat(settingsPath); os.IsNotExist(err) {
		t.Error("Settings file should have been created")
	}
	
	// Load and verify content
	loadedSettings, err := settings.LoadSettings(settingsPath)
	if err != nil {
		t.Fatalf("Failed to load saved settings: %v", err)
	}
	
	if len(loadedSettings.Boolean) != 2 {
		t.Errorf("Expected 2 boolean settings, got %d", len(loadedSettings.Boolean))
	}
}

func TestDefaultSettings(t *testing.T) {
	defaultSettings := settings.DefaultSettings()
	
	if len(defaultSettings.Boolean) != 1 {
		t.Errorf("Expected 1 default boolean setting, got %d", len(defaultSettings.Boolean))
	}
	
	defaultSetting := defaultSettings.Boolean[0]
	if defaultSetting.FilePath != "string" {
		t.Errorf("Expected default FilePath 'string', got '%s'", defaultSetting.FilePath)
	}
	
	if defaultSetting.Field != "string" {
		t.Errorf("Expected default Field 'string', got '%s'", defaultSetting.Field)
	}
	
	if defaultSetting.Value != "string" {
		t.Errorf("Expected default Value 'string', got '%s'", defaultSetting.Value)
	}
}
