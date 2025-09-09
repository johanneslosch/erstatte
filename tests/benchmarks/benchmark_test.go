package benchmarks_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/johanneslosch/erstatte/pkg/parser"
	"github.com/johanneslosch/erstatte/pkg/settings"
)

func BenchmarkParseAndReplace_JSON(b *testing.B) {
	tempDir := b.TempDir()
	jsonFile := filepath.Join(tempDir, "benchmark.json")
	
	content := `{
  "debugMode": false,
  "maxRetries": 3,
  "serverUrl": "localhost",
  "enableFeature": true,
  "timeout": 5000,
  "apiKey": "secret-key"
}`
	
	replacements := map[string]string{
		"debugMode":     "true",
		"maxRetries":    "10",
		"serverUrl":     "production.example.com",
		"enableFeature": "false",
	}
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		// Reset file content for each iteration
		err := os.WriteFile(jsonFile, []byte(content), 0644)
		if err != nil {
			b.Fatalf("Failed to reset file content: %v", err)
		}
		
		err = parser.ParseAndReplace(jsonFile, replacements)
		if err != nil {
			b.Fatalf("ParseAndReplace failed: %v", err)
		}
	}
}

func BenchmarkParseAndReplace_TypeScript(b *testing.B) {
	tempDir := b.TempDir()
	tsFile := filepath.Join(tempDir, "benchmark.ts")
	
	content := `export const config = {
  debugMode: false,
  maxRetries: 3,
  serverUrl: "localhost",
  enableFeature: true,
  timeout: 5000
};`
	
	replacements := map[string]string{
		"debugMode":     "true",
		"maxRetries":    "10",
		"serverUrl":     "production.example.com",
		"enableFeature": "false",
	}
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		// Reset file content for each iteration
		err := os.WriteFile(tsFile, []byte(content), 0644)
		if err != nil {
			b.Fatalf("Failed to reset file content: %v", err)
		}
		
		err = parser.ParseAndReplace(tsFile, replacements)
		if err != nil {
			b.Fatalf("ParseAndReplace failed: %v", err)
		}
	}
}

func BenchmarkLoadSettings(b *testing.B) {
	tempDir := b.TempDir()
	settingsFile := filepath.Join(tempDir, "benchmark_settings.json")
	
	testSettings := settings.Settings{
		Boolean: []settings.BooleanSetting{
			{FilePath: "./file1.ts", Field: "field1", Value: "true"},
			{FilePath: "./file2.ts", Field: "field2", Value: "false"},
			{FilePath: "./file3.json", Field: "field3", Value: "true"},
		},
	}
	
	err := settings.SaveSettings(settingsFile, testSettings)
	if err != nil {
		b.Fatalf("Failed to save benchmark settings: %v", err)
	}
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		_, err := settings.LoadSettings(settingsFile)
		if err != nil {
			b.Fatalf("LoadSettings failed: %v", err)
		}
	}
}

func BenchmarkSaveSettings(b *testing.B) {
	tempDir := b.TempDir()
	
	testSettings := settings.Settings{
		Boolean: []settings.BooleanSetting{
			{FilePath: "./file1.ts", Field: "field1", Value: "true"},
			{FilePath: "./file2.ts", Field: "field2", Value: "false"},
			{FilePath: "./file3.json", Field: "field3", Value: "true"},
		},
	}
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		settingsFile := filepath.Join(tempDir, "benchmark_save_settings.json")
		err := settings.SaveSettings(settingsFile, testSettings)
		if err != nil {
			b.Fatalf("SaveSettings failed: %v", err)
		}
		// Clean up for next iteration
		os.Remove(settingsFile)
	}
}
