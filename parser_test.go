package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
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
	
	err = os.WriteFile(jsonFile, data, 0644)
	if err != nil {
		t.Fatalf("Failed to write test JSON file: %v", err)
	}
	
	// Test replacing values
	replacements := map[string]string{
		"debugMode":  "true",
		"maxRetries": "5",
		"serverUrl":  "production.example.com",
	}
	
	err = ParseAndReplace(jsonFile, replacements)
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
	
	if result["maxRetries"] != "5" {
		t.Errorf("Expected maxRetries to be '5', got %v", result["maxRetries"])
	}
	
	if result["serverUrl"] != "production.example.com" {
		t.Errorf("Expected serverUrl to be 'production.example.com', got %v", result["serverUrl"])
	}
	
	// Verify unchanged value
	if result["enableFeature"] != true {
		t.Errorf("Expected enableFeature to remain true, got %v", result["enableFeature"])
	}
}

func TestParseAndReplace_Properties(t *testing.T) {
	tempDir := t.TempDir()
	propsFile := filepath.Join(tempDir, "test.properties")
	
	// Create test properties file
	content := `# Configuration file
debug.enabled=false
server.port=8080
server.host=localhost
# Another comment
app.name=TestApp
database.url=jdbc:mysql://localhost:3306/test
`
	
	err := os.WriteFile(propsFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to write test properties file: %v", err)
	}
	
	// Test replacing values
	replacements := map[string]string{
		"debug.enabled": "true",
		"server.port":   "9090",
		"app.name":      "ProductionApp",
	}
	
	err = ParseAndReplace(propsFile, replacements)
	if err != nil {
		t.Fatalf("ParseAndReplace failed: %v", err)
	}
	
	// Verify changes
	modifiedContent, err := os.ReadFile(propsFile)
	if err != nil {
		t.Fatalf("Failed to read modified properties file: %v", err)
	}
	
	contentStr := string(modifiedContent)
	
	if !strings.Contains(contentStr, "debug.enabled=true") {
		t.Error("debug.enabled should be changed to true")
	}
	
	if !strings.Contains(contentStr, "server.port=9090") {
		t.Error("server.port should be changed to 9090")
	}
	
	if !strings.Contains(contentStr, "app.name=ProductionApp") {
		t.Error("app.name should be changed to ProductionApp")
	}
	
	// Verify unchanged value
	if !strings.Contains(contentStr, "server.host=localhost") {
		t.Error("server.host should remain unchanged")
	}
	
	// Verify comments are preserved
	if !strings.Contains(contentStr, "# Configuration file") {
		t.Error("Comments should be preserved")
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
  timeout: 5000,
  apiKey: "secret-key"
};

export const settings = {
  darkMode: false,
  notifications: true,
  version: "1.0.0"
};`
	
	err := os.WriteFile(tsFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to write test TypeScript file: %v", err)
	}
	
	// Test replacing values
	replacements := map[string]string{
		"debugMode":      "true",
		"maxRetries":     "10",
		"serverUrl":      "production.example.com",
		"enableFeature":  "false",
		"darkMode":       "true",
		"notifications":  "false",
	}
	
	err = ParseAndReplace(tsFile, replacements)
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
	
	if !strings.Contains(contentStr, "darkMode: true") {
		t.Error("darkMode should be changed to true (without quotes)")
	}
	
	if !strings.Contains(contentStr, "notifications: false") {
		t.Error("notifications should be changed to false (without quotes)")
	}
	
	// Test numeric values (should not have quotes)
	if !strings.Contains(contentStr, "maxRetries: 10") {
		t.Error("maxRetries should be changed to 10 (without quotes)")
	}
	
	// Test string values (should have quotes)
	if !strings.Contains(contentStr, `serverUrl: "production.example.com"`) {
		t.Error("serverUrl should be changed to production.example.com (with quotes)")
	}
	
	// Verify unchanged values
	if !strings.Contains(contentStr, "timeout: 5000") {
		t.Error("timeout should remain unchanged")
	}
	
	if !strings.Contains(contentStr, `apiKey: "secret-key"`) {
		t.Error("apiKey should remain unchanged")
	}
}

func TestParseAndReplace_UnsupportedFileType(t *testing.T) {
	tempDir := t.TempDir()
	txtFile := filepath.Join(tempDir, "test.txt")
	
	// Create test file with unsupported extension
	err := os.WriteFile(txtFile, []byte("some content"), 0644)
	if err != nil {
		t.Fatalf("Failed to write test txt file: %v", err)
	}
	
	replacements := map[string]string{
		"test": "value",
	}
	
	err = ParseAndReplace(txtFile, replacements)
	if err == nil {
		t.Error("Expected error for unsupported file type, got nil")
	}
	
	if !strings.Contains(err.Error(), "unsupported file type") {
		t.Errorf("Expected 'unsupported file type' error, got: %v", err)
	}
}

func TestParseAndReplace_NonExistentFile(t *testing.T) {
	nonExistentFile := filepath.Join(os.TempDir(), "nonexistent.json")
	
	replacements := map[string]string{
		"test": "value",
	}
	
	err := ParseAndReplace(nonExistentFile, replacements)
	if err == nil {
		t.Error("Expected error for non-existent file, got nil")
	}
}

func TestParseAndReplace_InvalidJSON(t *testing.T) {
	tempDir := t.TempDir()
	jsonFile := filepath.Join(tempDir, "invalid.json")
	
	// Create invalid JSON file
	err := os.WriteFile(jsonFile, []byte("invalid json content"), 0644)
	if err != nil {
		t.Fatalf("Failed to write invalid JSON file: %v", err)
	}
	
	replacements := map[string]string{
		"test": "value",
	}
	
	err = ParseAndReplace(jsonFile, replacements)
	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}

func TestParseAndReplace_EmptyReplacements(t *testing.T) {
	tempDir := t.TempDir()
	jsonFile := filepath.Join(tempDir, "test.json")
	
	// Create test JSON file
	testData := map[string]interface{}{
		"debugMode": false,
	}
	
	data, err := json.MarshalIndent(testData, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal test JSON: %v", err)
	}
	
	err = os.WriteFile(jsonFile, data, 0644)
	if err != nil {
		t.Fatalf("Failed to write test JSON file: %v", err)
	}
	
	// Test with empty replacements
	replacements := map[string]string{}
	
	err = ParseAndReplace(jsonFile, replacements)
	if err != nil {
		t.Fatalf("ParseAndReplace should not fail with empty replacements: %v", err)
	}
	
	// Verify file content remains unchanged
	content, err := os.ReadFile(jsonFile)
	if err != nil {
		t.Fatalf("Failed to read JSON file: %v", err)
	}
	
	var result map[string]interface{}
	err = json.Unmarshal(content, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}
	
	if result["debugMode"] != false {
		t.Error("File content should remain unchanged with empty replacements")
	}
}

func TestParseAndReplace_TypeScriptComplexPattern(t *testing.T) {
	tempDir := t.TempDir()
	tsFile := filepath.Join(tempDir, "complex.ts")
	
	// Create TypeScript file with complex patterns
	content := `export const locationWorkerSettings = {
  locationWorkerlogs: false,
  debugLocationWorker: false,
  debugPermissions: false,
  debugInitialization: false,
};

export const databaseSettings = {
  environment: "development",
  database: {
    name: "databaseName",
    resetDatabase: false,
    resetAsyncStorage: false,
    fakeDataCountWhenEmptyTodos: 0,
  },
};`
	
	err := os.WriteFile(tsFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to write test TypeScript file: %v", err)
	}
	
	// Test replacing nested values
	replacements := map[string]string{
		"resetDatabase":      "true",
		"resetAsyncStorage":  "true",
		"debugLocationWorker": "true",
		"environment":        "production",
	}
	
	err = ParseAndReplace(tsFile, replacements)
	if err != nil {
		t.Fatalf("ParseAndReplace failed: %v", err)
	}
	
	// Verify changes
	modifiedContent, err := os.ReadFile(tsFile)
	if err != nil {
		t.Fatalf("Failed to read modified TypeScript file: %v", err)
	}
	
	contentStr := string(modifiedContent)
	
	if !strings.Contains(contentStr, "resetDatabase: true") {
		t.Error("resetDatabase should be changed to true")
	}
	
	if !strings.Contains(contentStr, "resetAsyncStorage: true") {
		t.Error("resetAsyncStorage should be changed to true")
	}
	
	if !strings.Contains(contentStr, "debugLocationWorker: true") {
		t.Error("debugLocationWorker should be changed to true")
	}
	
	if !strings.Contains(contentStr, `environment: "production"`) {
		t.Error("environment should be changed to production (with quotes)")
	}
}
