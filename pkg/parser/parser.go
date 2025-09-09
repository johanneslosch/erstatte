package parser

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// ParseAndReplace replaces values in .json, .properties, and .ts files based on the replacements map.
func ParseAndReplace(filePath string, replacements map[string]string) error {
	ext := strings.ToLower(filepath.Ext(filePath))
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	data := string(content)

	switch ext {
	case ".json":
		var obj map[string]interface{}
		if err := json.Unmarshal(content, &obj); err != nil {
			return err
		}
		for k, v := range replacements {
			obj[k] = v
		}
		newContent, err := json.MarshalIndent(obj, "", "  ")
		if err != nil {
			return err
		}
		return os.WriteFile(filePath, newContent, 0644)

	case ".properties":
		lines := strings.Split(data, "\n")
		for i, line := range lines {
			if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
				continue
			}
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				if val, ok := replacements[key]; ok {
					lines[i] = key + "=" + val
				}
			}
		}
		return os.WriteFile(filePath, []byte(strings.Join(lines, "\n")), 0644)

	case ".ts":
		for k, v := range replacements {
			pattern := regexp.QuoteMeta(k) + `\s*:\s*([^,\n]+)`
			re := regexp.MustCompile(pattern)
			// Check if value is boolean or number, otherwise set as string
			if !(v == "true" || v == "false" || regexp.MustCompile(`^\d+$`).MatchString(v)) {
				v = `"` + v + `"`
			}
			data = re.ReplaceAllString(data, k+": "+v)
		}
		return os.WriteFile(filePath, []byte(data), 0644)

	default:
		return errors.New("unsupported file type")
	}
}
