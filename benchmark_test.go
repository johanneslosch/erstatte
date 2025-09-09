package main

import (
	"os"
	"path/filepath"
	"testing"
)

func BenchmarkParseAndReplace_JSON(b *testing.B) {
	tempDir := b.TempDir()
	jsonFile := filepath.Join(tempDir, "benchmark.json")
	
	// Create a larger JSON file for benchmarking
	content := `{
  "debugMode": false,
  "maxRetries": 3,
  "serverUrl": "localhost",
  "enableFeature": true,
  "timeout": 5000,
  "apiKey": "secret-key",
  "database": {
    "host": "localhost",
    "port": 5432,
    "name": "testdb"
  },
  "logging": {
    "level": "info",
    "enableConsole": true,
    "enableFile": false
  }
}`
	
	err := os.WriteFile(jsonFile, []byte(content), 0644)
	if err != nil {
		b.Fatalf("Failed to create benchmark JSON file: %v", err)
	}
	
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
		
		err = ParseAndReplace(jsonFile, replacements)
		if err != nil {
			b.Fatalf("ParseAndReplace failed: %v", err)
		}
	}
}

func BenchmarkParseAndReplace_TypeScript(b *testing.B) {
	tempDir := b.TempDir()
	tsFile := filepath.Join(tempDir, "benchmark.ts")
	
	// Create a larger TypeScript file for benchmarking
	content := `export const locationWorkerSettings = {
  locationWorkerlogs: false,
  debugLocationWorker: false,
  debugPermissions: false,
  debugInitialization: false,
  enableGPS: true,
  trackingInterval: 5000,
  maxRetries: 3,
  timeout: 10000,
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
    enableMigrations: true,
    backupInterval: 3600,
  },
  logging: {
    databaseProviderLogs: false,
    todoCreationLogs: false,
    todoGetLogs: false,
    todoUpdateLogs: false,
    todoDeleteLogs: false,
    todoNotificationLogs: false,
    tagCreationLogs: false,
    tagGetLogs: false,
    tagUpdateLogs: false,
    tagDeleteLogs: false,
    todoTagCreationLogs: false,
    todoTagGetLogs: false,
    todoTagRemovalLogs: false,
    statsLogs: false,
    syncLogs: false,
    verboseMode: false,
  },
};`
	
	err := os.WriteFile(tsFile, []byte(content), 0644)
	if err != nil {
		b.Fatalf("Failed to create benchmark TypeScript file: %v", err)
	}
	
	replacements := map[string]string{
		"resetDatabase":      "true",
		"resetAsyncStorage":  "true",
		"debugLocationWorker": "true",
		"environment":        "production",
		"verboseMode":        "true",
		"enableGPS":          "false",
	}
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		// Reset file content for each iteration
		err := os.WriteFile(tsFile, []byte(content), 0644)
		if err != nil {
			b.Fatalf("Failed to reset file content: %v", err)
		}
		
		err = ParseAndReplace(tsFile, replacements)
		if err != nil {
			b.Fatalf("ParseAndReplace failed: %v", err)
		}
	}
}

func BenchmarkLoadSettings(b *testing.B) {
	tempDir := b.TempDir()
	settingsFile := filepath.Join(tempDir, "benchmark_settings.json")
	
	// Create a settings file with multiple entries
	settings := Settings{
		Boolean: []BooleanSetting{
			{FilePath: "./file1.ts", Field: "field1", Value: "true"},
			{FilePath: "./file2.ts", Field: "field2", Value: "false"},
			{FilePath: "./file3.json", Field: "field3", Value: "true"},
			{FilePath: "./file4.properties", Field: "field4", Value: "false"},
			{FilePath: "./file5.ts", Field: "field5", Value: "true"},
		},
	}
	
	err := SaveSettings(settingsFile, settings)
	if err != nil {
		b.Fatalf("Failed to save benchmark settings: %v", err)
	}
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		_, err := LoadSettings(settingsFile)
		if err != nil {
			b.Fatalf("LoadSettings failed: %v", err)
		}
	}
}

func BenchmarkSaveSettings(b *testing.B) {
	tempDir := b.TempDir()
	
	// Create a settings object with multiple entries
	settings := Settings{
		Boolean: []BooleanSetting{
			{FilePath: "./file1.ts", Field: "field1", Value: "true"},
			{FilePath: "./file2.ts", Field: "field2", Value: "false"},
			{FilePath: "./file3.json", Field: "field3", Value: "true"},
			{FilePath: "./file4.properties", Field: "field4", Value: "false"},
			{FilePath: "./file5.ts", Field: "field5", Value: "true"},
		},
	}
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		settingsFile := filepath.Join(tempDir, "benchmark_save_settings.json")
		err := SaveSettings(settingsFile, settings)
		if err != nil {
			b.Fatalf("SaveSettings failed: %v", err)
		}
		// Clean up for next iteration
		os.Remove(settingsFile)
	}
}
