package parser_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/johanneslosch/erstatte/pkg/parser"
	"github.com/johanneslosch/erstatte/pkg/settings"
)

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

func TestParseAndReplace_TypeScript(t *testing.T) {
	tempDir := t.TempDir()
	tsFile := filepath.Join(tempDir, "test.ts")

	// Create test TypeScript file
	content := `export const config = {
  debugMode: false,
  maxRetries: 3,
  serverUrl: "localhost",
  enableFeature: true,
};`

	err := os.WriteFile(tsFile, []byte(content), 0600)
	if err != nil {
		t.Fatalf("Failed to write test TypeScript file: %v", err)
	}

	// Test replacing values
	replacements := map[string]string{
		"debugMode":     "true",
		"maxRetries":    "10",
		"serverUrl":     "production.example.com",
		"enableFeature": "false",
	}

	err = parser.ParseAndReplace(tsFile, replacements)
	if err != nil {
		t.Fatalf("ParseAndReplace failed: %v", err)
	}

	// Verify changes
	modifiedContent, err := os.ReadFile(tsFile)
	if err != nil {
		t.Fatalf("Failed to read modified TypeScript file: %v", err)
	}

	contentStr := string(modifiedContent)

	// Test boolean values (should not have quotes)
	if !strings.Contains(contentStr, "debugMode: true") {
		t.Error("debugMode should be changed to true (without quotes)")
	}

	if !strings.Contains(contentStr, "enableFeature: false") {
		t.Error("enableFeature should be changed to false (without quotes)")
	}

	// Test string values (should have quotes)
	if !strings.Contains(contentStr, `serverUrl: "production.example.com"`) {
		t.Error("serverUrl should be changed to production.example.com (with quotes)")
	}
}

func TestParseAndReplace_UnsupportedFileType(t *testing.T) {
	tempDir := t.TempDir()
	txtFile := filepath.Join(tempDir, "test.txt")

	// Create test file with unsupported extension
	err := os.WriteFile(txtFile, []byte("some content"), 0600)
	if err != nil {
		t.Fatalf("Failed to write test txt file: %v", err)
	}

	replacements := map[string]string{
		"test": "value",
	}

	err = parser.ParseAndReplace(txtFile, replacements)
	if err == nil {
		t.Error("Expected error for unsupported file type, got nil")
	}

	if !strings.Contains(err.Error(), "unsupported file type") {
		t.Errorf("Expected 'unsupported file type' error, got: %v", err)
	}
}

func TestParseAndReplace_Properties(t *testing.T) {
	tempDir := t.TempDir()
	propFile := filepath.Join(tempDir, "test.properties")

	// Create test properties file
	content := "# sample properties file\nusername=olduser\npassword=secret\n# comment\n"

	err := os.WriteFile(propFile, []byte(content), 0600)
	if err != nil {
		t.Fatalf("Failed to write test properties file: %v", err)
	}

	// Test replacing values
	replacements := map[string]string{
		"username": "newuser",
		"password": "newsecret",
	}

	err = parser.ParseAndReplace(propFile, replacements)
	if err != nil {
		t.Fatalf("ParseAndReplace failed: %v", err)
	}

	// Verify changes
	modifiedContent, err := os.ReadFile(propFile)
	if err != nil {
		t.Fatalf("Failed to read modified properties file: %v", err)
	}

	contentStr := string(modifiedContent)

	if !strings.Contains(contentStr, "username=newuser") {
		t.Errorf("Expected username to be changed to newuser, got: %s", contentStr)
	}

	if !strings.Contains(contentStr, "password=newsecret") {
		t.Errorf("Expected password to be changed to newsecret, got: %s", contentStr)
	}
}

func TestEndToEnd_ReplaceViaErstatteJSON(t *testing.T) {
	tempDir := t.TempDir()
	// properties file that will be modified
	propFile := filepath.Join(tempDir, "app.properties")

	// create properties file
	content := "host=localhost\nport=8080\nuser=olduser\n"
	if err := os.WriteFile(propFile, []byte(content), 0600); err != nil {
		t.Fatalf("Failed to write properties file: %v", err)
	}

		// create erstatte.json settings that point to the properties file
		settingsPath := filepath.Join(tempDir, "erstatte.json")

	// Use the settings struct and helper to write a valid JSON file (handles escaping)
	settingsObj := settings.Settings{
		Boolean: []settings.InternalSetting{
			{
				FilePath: propFile,
				Field:    "user",
				Value:    "deployuser",
			},
		},
	}

		if err := settings.SaveSettings(settingsPath, settingsObj); err != nil {
				t.Fatalf("Failed to write erstatte.json: %v", err)
		}

	// Load settings from the settings file and process them like main
	loaded, err := settings.LoadSettings(settingsPath)
	if err != nil {
		t.Fatalf("LoadSettings failed: %v", err)
	}

	for _, setting := range loaded.Boolean {
		replacements := map[string]string{setting.Field: setting.Value}
		if err := parser.ParseAndReplace(setting.FilePath, replacements); err != nil {
			t.Fatalf("ParseAndReplace failed: %v", err)
		}
	}

	// Verify the property was updated
	modified, err := os.ReadFile(propFile)
	if err != nil {
		t.Fatalf("Failed to read modified properties file: %v", err)
	}

	if !strings.Contains(string(modified), "user=deployuser") {
		t.Fatalf("Expected user=deployuser in properties file, got: %s", string(modified))
	}
}

// --- Additional string-focused tests ---

func TestParseAndReplace_JSON_String(t *testing.T) {
	tempDir := t.TempDir()
	jsonFile := filepath.Join(tempDir, "strings.json")

	testData := map[string]interface{}{
		"name": "old-name",
		"env":  "dev",
	}
	data, _ := json.MarshalIndent(testData, "", "  ")
	if err := os.WriteFile(jsonFile, data, 0600); err != nil {
		t.Fatalf("Failed to write json file: %v", err)
	}

	replacements := map[string]string{"name": "new-name"}
	if err := parser.ParseAndReplace(jsonFile, replacements); err != nil {
		t.Fatalf("ParseAndReplace failed: %v", err)
	}

	modified, _ := os.ReadFile(jsonFile)
	var result map[string]interface{}
	if err := json.Unmarshal(modified, &result); err != nil {
		t.Fatalf("Failed to unmarshal modified JSON: %v", err)
	}

	if result["name"] != "new-name" {
		t.Fatalf("Expected name to be 'new-name', got: %v", result["name"])
	}
}

func TestParseAndReplace_TypeScript_StringOnly(t *testing.T) {
	tempDir := t.TempDir()
	tsFile := filepath.Join(tempDir, "strings.ts")

	content := `export const cfg = { title: "Old Title", description: "desc" };`
	if err := os.WriteFile(tsFile, []byte(content), 0600); err != nil {
		t.Fatalf("Failed to write ts file: %v", err)
	}

	replacements := map[string]string{"title": "New Title"}
	if err := parser.ParseAndReplace(tsFile, replacements); err != nil {
		t.Fatalf("ParseAndReplace failed: %v", err)
	}

	modified, _ := os.ReadFile(tsFile)
	if !strings.Contains(string(modified), `title: "New Title"`) {
		t.Fatalf("Expected title to be replaced with quoted string, got: %s", string(modified))
	}
}

func TestParseAndReplace_Properties_StringOnly(t *testing.T) {
	tempDir := t.TempDir()
	propFile := filepath.Join(tempDir, "app.properties")
	content := "mode=development\nservice=old-service\n"
	if err := os.WriteFile(propFile, []byte(content), 0600); err != nil {
		t.Fatalf("Failed to write properties file: %v", err)
	}

	replacements := map[string]string{"service": "new-service"}
	if err := parser.ParseAndReplace(propFile, replacements); err != nil {
		t.Fatalf("ParseAndReplace failed: %v", err)
	}

	modified, _ := os.ReadFile(propFile)
	if !strings.Contains(string(modified), "service=new-service") {
		t.Fatalf("Expected service to be new-service, got: %s", string(modified))
	}
}

func TestEndToEnd_StringsViaErstatteJSON(t *testing.T) {
	tempDir := t.TempDir()
	propFile := filepath.Join(tempDir, "svc.properties")
	if err := os.WriteFile(propFile, []byte("svc=old\n"), 0600); err != nil {
		t.Fatalf("Failed to write properties file: %v", err)
	}

	settingsPath := filepath.Join(tempDir, "erstatte.json")
	settingsObj := settings.Settings{
		Boolean: []settings.InternalSetting{
			{FilePath: propFile, Field: "svc", Value: "newsvc"},
		},
	}
	if err := settings.SaveSettings(settingsPath, settingsObj); err != nil {
		t.Fatalf("Failed to write erstatte.json: %v", err)
	}

	loaded, err := settings.LoadSettings(settingsPath)
	if err != nil {
		t.Fatalf("LoadSettings failed: %v", err)
	}

	for _, s := range loaded.Boolean {
		if err := parser.ParseAndReplace(s.FilePath, map[string]string{s.Field: s.Value}); err != nil {
			t.Fatalf("ParseAndReplace failed: %v", err)
		}
	}

	modified, _ := os.ReadFile(propFile)
	if !strings.Contains(string(modified), "svc=newsvc") {
		t.Fatalf("Expected svc=newsvc, got: %s", string(modified))
	}
}
