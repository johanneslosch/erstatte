package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadSettings_FileExists(t *testing.T) {
	// Create a temporary settings file
	tempDir := t.TempDir()
	settingsPath := filepath.Join(tempDir, "test_settings.json")
	
	expectedSettings := Settings{
		Boolean: []BooleanSetting{
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
	settings, err := LoadSettings(settingsPath)
	if err != nil {
		t.Fatalf("LoadSettings failed: %v", err)
	}
	
	// Verify the loaded settings
	if len(settings.Boolean) != 1 {
		t.Errorf("Expected 1 boolean setting, got %d", len(settings.Boolean))
	}
	
	if settings.Boolean[0].FilePath != "./test.ts" {
		t.Errorf("Expected FilePath './test.ts', got '%s'", settings.Boolean[0].FilePath)
	}
	
	if settings.Boolean[0].Field != "testField" {
		t.Errorf("Expected Field 'testField', got '%s'", settings.Boolean[0].Field)
	}
	
	if settings.Boolean[0].Value != "true" {
		t.Errorf("Expected Value 'true', got '%s'", settings.Boolean[0].Value)
	}
}

func TestLoadSettings_FileDoesNotExist(t *testing.T) {
	tempDir := t.TempDir()
	settingsPath := filepath.Join(tempDir, "nonexistent_settings.json")
	
	// Test loading non-existent settings file
	settings, err := LoadSettings(settingsPath)
	if err != nil {
		t.Fatalf("LoadSettings failed: %v", err)
	}
	
	// Should create default settings
	if len(settings.Boolean) != 1 {
		t.Errorf("Expected 1 default boolean setting, got %d", len(settings.Boolean))
	}
	
	// Verify file was created
	if _, err := os.Stat(settingsPath); os.IsNotExist(err) {
		t.Error("Settings file should have been created")
	}
}

func TestLoadSettings_InvalidJSON(t *testing.T) {
	tempDir := t.TempDir()
	settingsPath := filepath.Join(tempDir, "invalid_settings.json")
	
	// Write invalid JSON
	err := os.WriteFile(settingsPath, []byte("invalid json content"), 0644)
	if err != nil {
		t.Fatalf("Failed to write invalid JSON file: %v", err)
	}
	
	// Test loading invalid JSON
	_, err = LoadSettings(settingsPath)
	if err == nil {
		t.Error("Expected error when loading invalid JSON, got nil")
	}
}

func TestSaveSettings(t *testing.T) {
	tempDir := t.TempDir()
	settingsPath := filepath.Join(tempDir, "save_test_settings.json")
	
	settings := Settings{
		Boolean: []BooleanSetting{
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
	err := SaveSettings(settingsPath, settings)
	if err != nil {
		t.Fatalf("SaveSettings failed: %v", err)
	}
	
	// Verify file was created
	if _, err := os.Stat(settingsPath); os.IsNotExist(err) {
		t.Error("Settings file should have been created")
	}
	
	// Load and verify content
	loadedSettings, err := LoadSettings(settingsPath)
	if err != nil {
		t.Fatalf("Failed to load saved settings: %v", err)
	}
	
	if len(loadedSettings.Boolean) != 2 {
		t.Errorf("Expected 2 boolean settings, got %d", len(loadedSettings.Boolean))
	}
	
	// Verify first setting
	if loadedSettings.Boolean[0].FilePath != "./example.ts" {
		t.Errorf("Expected FilePath './example.ts', got '%s'", loadedSettings.Boolean[0].FilePath)
	}
	
	if loadedSettings.Boolean[0].Field != "debugMode" {
		t.Errorf("Expected Field 'debugMode', got '%s'", loadedSettings.Boolean[0].Field)
	}
	
	if loadedSettings.Boolean[0].Value != "false" {
		t.Errorf("Expected Value 'false', got '%s'", loadedSettings.Boolean[0].Value)
	}
}

func TestDefaultSettings(t *testing.T) {
	settings := defaultSettings()
	
	if len(settings.Boolean) != 1 {
		t.Errorf("Expected 1 default boolean setting, got %d", len(settings.Boolean))
	}
	
	defaultSetting := settings.Boolean[0]
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

func TestBooleanSettingJSONMarshaling(t *testing.T) {
	setting := BooleanSetting{
		FilePath: "./test.ts",
		Field:    "testField",
		Value:    "true",
	}
	
	// Test marshaling
	data, err := json.Marshal(setting)
	if err != nil {
		t.Fatalf("Failed to marshal BooleanSetting: %v", err)
	}
	
	// Test unmarshaling
	var unmarshaled BooleanSetting
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal BooleanSetting: %v", err)
	}
	
	if unmarshaled.FilePath != setting.FilePath {
		t.Errorf("FilePath mismatch after marshaling/unmarshaling")
	}
	
	if unmarshaled.Field != setting.Field {
		t.Errorf("Field mismatch after marshaling/unmarshaling")
	}
	
	if unmarshaled.Value != setting.Value {
		t.Errorf("Value mismatch after marshaling/unmarshaling")
	}
}
