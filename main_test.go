package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMainIntegration(t *testing.T) {
	// Create a temporary directory for the integration test
	tempDir := t.TempDir()
	
	// Create a test TypeScript file
	testTSFile := filepath.Join(tempDir, "test.ts")
	tsContent := `export const config = {
  debugMode: false,
  enableFeature: true,
  serverUrl: "localhost",
  maxRetries: 3
};`
	
	err := os.WriteFile(testTSFile, []byte(tsContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test TypeScript file: %v", err)
	}
	
	// Create a test settings file
	settingsFile := filepath.Join(tempDir, "test_settings.json")
	settings := Settings{
		Boolean: []BooleanSetting{
			{
				FilePath: testTSFile,
				Field:    "debugMode",
				Value:    "true",
			},
			{
				FilePath: testTSFile,
				Field:    "enableFeature",
				Value:    "false",
			},
			{
				FilePath: testTSFile,
				Field:    "serverUrl",
				Value:    "production.example.com",
			},
		},
	}
	
	err = SaveSettings(settingsFile, settings)
	if err != nil {
		t.Fatalf("Failed to save test settings: %v", err)
	}
	
	// Test loading settings
	loadedSettings, err := LoadSettings(settingsFile)
	if err != nil {
		t.Fatalf("Failed to load settings: %v", err)
	}
	
	// Verify settings were loaded correctly
	if len(loadedSettings.Boolean) != 3 {
		t.Errorf("Expected 3 boolean settings, got %d", len(loadedSettings.Boolean))
	}
	
	// Test applying settings (simulate main function logic)
	for _, s := range loadedSettings.Boolean {
		replacements := map[string]string{s.Field: s.Value}
		err := ParseAndReplace(s.FilePath, replacements)
		if err != nil {
			t.Errorf("Failed to apply setting %s: %v", s.Field, err)
		}
	}
	
	// Verify the changes were applied to the TypeScript file
	modifiedContent, err := os.ReadFile(testTSFile)
	if err != nil {
		t.Fatalf("Failed to read modified TypeScript file: %v", err)
	}
	
	modifiedStr := string(modifiedContent)
	
	// Check that all values were replaced correctly
	expectedChanges := []string{
		"debugMode: true",
		"enableFeature: false",
		`serverUrl: "production.example.com"`,
	}
	
	for _, expected := range expectedChanges {
		if !containsString(modifiedStr, expected) {
			t.Errorf("Expected to find '%s' in modified file", expected)
		}
	}
	
	// Verify unchanged value
	if !containsString(modifiedStr, "maxRetries: 3") {
		t.Error("maxRetries should remain unchanged")
	}
}

func TestMainWithMissingFile(t *testing.T) {
	tempDir := t.TempDir()
	
	// Create settings pointing to non-existent file
	settingsFile := filepath.Join(tempDir, "test_settings.json")
	settingsContent := `{
  "boolean": [
    {
      "filePath": "/nonexistent/file.ts",
      "field": "debugMode",
      "value": "true"
    }
  ]
}`
	
	err := os.WriteFile(settingsFile, []byte(settingsContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test settings file: %v", err)
	}
	
	// Test loading settings
	settings, err := LoadSettings(settingsFile)
	if err != nil {
		t.Fatalf("Failed to load settings: %v", err)
	}
	
	// Test applying settings to non-existent file (should fail)
	for _, s := range settings.Boolean {
		replacements := map[string]string{s.Field: s.Value}
		err := ParseAndReplace(s.FilePath, replacements)
		if err == nil {
			t.Error("Expected error when trying to parse non-existent file")
		}
	}
}

func TestMainWithEmptySettings(t *testing.T) {
	tempDir := t.TempDir()
	
	// Create empty settings file
	settingsFile := filepath.Join(tempDir, "empty_settings.json")
	settingsContent := `{
  "boolean": []
}`
	
	err := os.WriteFile(settingsFile, []byte(settingsContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create empty settings file: %v", err)
	}
	
	// Test loading empty settings
	settings, err := LoadSettings(settingsFile)
	if err != nil {
		t.Fatalf("Failed to load empty settings: %v", err)
	}
	
	// Verify empty settings
	if len(settings.Boolean) != 0 {
		t.Errorf("Expected 0 boolean settings, got %d", len(settings.Boolean))
	}
	
	// Test that no operations are performed (should not error)
	for _, s := range settings.Boolean {
		replacements := map[string]string{s.Field: s.Value}
		err := ParseAndReplace(s.FilePath, replacements)
		if err != nil {
			t.Errorf("Unexpected error with empty settings: %v", err)
		}
	}
}

func TestRealWorldScenario(t *testing.T) {
	tempDir := t.TempDir()
	
	// Create a realistic TypeScript configuration file (similar to your testfile.ts)
	testTSFile := filepath.Join(tempDir, "config.ts")
	tsContent := `export const locationWorkerSettings = {
  locationWorkerlogs: false,
  debugLocationWorker: false,
  debugPermissions: false,
  debugInitialization: false,
};

export const databaseSettings = {
  environment: "development",
  database: {
    name: "databaseName",
    location: "default",
    resetDatabase: false,
    resetAsyncStorage: false,
    fakeDataCountWhenEmptyTodos: 0,
    fakeDataCountWhenEmptyTags: 0,
  },
  logging: {
    databaseProviderLogs: false,
    todoCreationLogs: false,
    todoGetLogs: false,
  },
};`
	
	err := os.WriteFile(testTSFile, []byte(tsContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test TypeScript file: %v", err)
	}
	
	// Create realistic settings
	settingsFile := filepath.Join(tempDir, "config_settings.json")
	settings := Settings{
		Boolean: []BooleanSetting{
			{
				FilePath: testTSFile,
				Field:    "resetDatabase",
				Value:    "true",
			},
			{
				FilePath: testTSFile,
				Field:    "resetAsyncStorage",
				Value:    "true",
			},
			{
				FilePath: testTSFile,
				Field:    "debugLocationWorker",
				Value:    "true",
			},
			{
				FilePath: testTSFile,
				Field:    "environment",
				Value:    "production",
			},
		},
	}
	
	err = SaveSettings(settingsFile, settings)
	if err != nil {
		t.Fatalf("Failed to save test settings: %v", err)
	}
	
	// Load and apply settings
	loadedSettings, err := LoadSettings(settingsFile)
	if err != nil {
		t.Fatalf("Failed to load settings: %v", err)
	}
	
	for _, s := range loadedSettings.Boolean {
		replacements := map[string]string{s.Field: s.Value}
		err := ParseAndReplace(s.FilePath, replacements)
		if err != nil {
			t.Errorf("Failed to apply setting %s: %v", s.Field, err)
		}
	}
	
	// Verify all changes were applied correctly
	modifiedContent, err := os.ReadFile(testTSFile)
	if err != nil {
		t.Fatalf("Failed to read modified file: %v", err)
	}
	
	modifiedStr := string(modifiedContent)
	
	expectedChanges := []string{
		"resetDatabase: true",
		"resetAsyncStorage: true",
		"debugLocationWorker: true",
		`environment: "production"`,
	}
	
	for _, expected := range expectedChanges {
		if !containsString(modifiedStr, expected) {
			t.Errorf("Expected to find '%s' in modified file", expected)
		}
	}
	
	// Verify some values remained unchanged
	unchangedValues := []string{
		"locationWorkerlogs: false",
		"debugPermissions: false",
		"fakeDataCountWhenEmptyTodos: 0",
	}
	
	for _, unchanged := range unchangedValues {
		if !containsString(modifiedStr, unchanged) {
			t.Errorf("Expected unchanged value '%s' to remain in file", unchanged)
		}
	}
}

// Helper function to check if a string contains a substring
func containsString(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && 
		(findSubstring(s, substr) != -1))
}

// Simple substring search function
func findSubstring(s, substr string) int {
	if len(substr) == 0 {
		return 0
	}
	if len(substr) > len(s) {
		return -1
	}
	
	for i := 0; i <= len(s)-len(substr); i++ {
		match := true
		for j := 0; j < len(substr); j++ {
			if s[i+j] != substr[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}
