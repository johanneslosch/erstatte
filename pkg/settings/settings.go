package settings

import (
	"encoding/json"
	"os"
)

type BooleanSetting struct {
	FilePath string `json:"filePath"`
	Field    string `json:"field"`
	Value    string `json:"value"`
}

type Settings struct {
	Boolean []BooleanSetting `json:"boolean"`
}

func DefaultSettings() Settings {
	return Settings{
		Boolean: []BooleanSetting{
			{
				FilePath: "string",
				Field:    "string",
				Value:    "string",
			},
		},
	}
}

func LoadSettings(path string) (Settings, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		s := DefaultSettings()
		if err := SaveSettings(path, s); err != nil {
			return Settings{}, err
		}
		return s, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return Settings{}, err
	}

	var s Settings
	if err := json.Unmarshal(data, &s); err != nil {
		return Settings{}, err
	}
	return s, nil
}

func SaveSettings(path string, s Settings) error {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
