package main

import (
	"fmt"

	"github.com/johanneslosch/erstatte/pkg/parser"
	"github.com/johanneslosch/erstatte/pkg/settings"
)

// version is set via ldflags during build
var version = "dev"

func main() {
	settingsData, err := settings.LoadSettings("settings.json")
	if err != nil {
		panic(err)
	}

	// Process each BooleanSetting with the parser
	for _, s := range settingsData.Boolean {
		replacements := map[string]string{s.Field: s.Value}
		err := parser.ParseAndReplace(s.FilePath, replacements)
		if err != nil {
			fmt.Printf("Error replacing in %s: %v\n", s.FilePath, err)
		} else {
			fmt.Printf("%s: %s = %s replaced\n", s.FilePath, s.Field, s.Value)
		}
	}
}
