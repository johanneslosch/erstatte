package parser_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/johanneslosch/erstatte/pkg/parser"
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
	
	err := os.WriteFile(tsFile, []byte(content), 0644)
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
	err := os.WriteFile(txtFile, []byte("some content"), 0644)
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
