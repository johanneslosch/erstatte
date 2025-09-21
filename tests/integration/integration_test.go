package integration_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/johanneslosch/erstatte/pkg/parser"
	"github.com/johanneslosch/erstatte/pkg/settings"
)

func TestEndToEndWorkflow(t *testing.T) {
	tempDir := t.TempDir()

	// Create test TypeScript file
	testTSFile := filepath.Join(tempDir, "config.ts")
	tsContent := `export const config = {
  debugMode: false,
  enableFeature: true,
  serverUrl: "localhost",
  maxRetries: 3
};`

	err := os.WriteFile(testTSFile, []byte(tsContent), 0600)
	if err != nil {
		t.Fatalf("Failed to create test TypeScript file: %v", err)
	}

	// Create settings
	settingsFile := filepath.Join(tempDir, "test_erstatte.json")
	testSettings := settings.Settings{
		Boolean: []settings.InternalSetting{
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

	err = settings.SaveSettings(settingsFile, testSettings)
	if err != nil {
		t.Fatalf("Failed to save test settings: %v", err)
	}

	// Load settings and apply changes
	loadedSettings, err := settings.LoadSettings(settingsFile)
	if err != nil {
		t.Fatalf("Failed to load settings: %v", err)
	}

	// Apply each setting
	for _, s := range loadedSettings.Boolean {
		replacements := map[string]string{s.Field: s.Value}
		if parseErr := parser.ParseAndReplace(s.FilePath, replacements); parseErr != nil {
			t.Errorf("Failed to apply setting %s: %v", s.Field, parseErr)
		}
	}

	// Verify the changes were applied
	modifiedContent, err := os.ReadFile(testTSFile)
	if err != nil {
		t.Fatalf("Failed to read modified TypeScript file: %v", err)
	}

	modifiedStr := string(modifiedContent)

	expectedChanges := []string{
		"debugMode: true",
		"enableFeature: false",
		`serverUrl: "production.example.com"`,
	}

	for _, expected := range expectedChanges {
		if !strings.Contains(modifiedStr, expected) {
			t.Errorf("Expected to find '%s' in modified file", expected)
		}
	}

	// Verify unchanged value
	if !strings.Contains(modifiedStr, "maxRetries: 3") {
		t.Error("maxRetries should remain unchanged")
	}
}

func TestComplexScenario(t *testing.T) {
	tempDir := t.TempDir()

	// Create multiple files
	jsonFile := filepath.Join(tempDir, "config.json")
	tsFile := filepath.Join(tempDir, "settings.ts")
	propsFile := filepath.Join(tempDir, "app.properties")

	// JSON file
	jsonContent := `{
  "debug": false,
  "port": 8080
}`
	err := os.WriteFile(jsonFile, []byte(jsonContent), 0600)
	if err != nil {
		t.Fatalf("Failed to create JSON file: %v", err)
	}

	// TypeScript file
	tsContent := `export const config = {
  production: false,
  apiUrl: "localhost"
};`
	err = os.WriteFile(tsFile, []byte(tsContent), 0600)
	if err != nil {
		t.Fatalf("Failed to create TypeScript file: %v", err)
	}

	// Properties file
	propsContent := `environment=development
server.host=localhost`
	err = os.WriteFile(propsFile, []byte(propsContent), 0600)
	if err != nil {
		t.Fatalf("Failed to create properties file: %v", err)
	}

	// Create settings for all files
	testSettings := settings.Settings{
		Boolean: []settings.InternalSetting{
			{FilePath: jsonFile, Field: "debug", Value: "true"},
			{FilePath: tsFile, Field: "production", Value: "true"},
			{FilePath: tsFile, Field: "apiUrl", Value: "api.production.com"},
			{FilePath: propsFile, Field: "environment", Value: "production"},
		},
	}

	// Apply all settings
	for _, s := range testSettings.Boolean {
		replacements := map[string]string{s.Field: s.Value}
		err := parser.ParseAndReplace(s.FilePath, replacements)
		if err != nil {
			t.Errorf("Failed to apply setting %s in %s: %v", s.Field, s.FilePath, err)
		}
	}

	// Verify JSON changes
	jsonModified, _ := os.ReadFile(jsonFile)
	if !strings.Contains(string(jsonModified), `"debug": "true"`) {
		t.Error("JSON debug should be changed to true")
	}

	// Verify TypeScript changes
	tsModified, _ := os.ReadFile(tsFile)
	tsStr := string(tsModified)
	if !strings.Contains(tsStr, "production: true") {
		t.Error("TypeScript production should be changed to true")
	}
	if !strings.Contains(tsStr, `apiUrl: "api.production.com"`) {
		t.Error("TypeScript apiUrl should be changed")
	}

	// Verify Properties changes
	propsModified, _ := os.ReadFile(propsFile)
	if !strings.Contains(string(propsModified), "environment=production") {
		t.Error("Properties environment should be changed to production")
	}
}
