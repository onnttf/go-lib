package file

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

// LoadJSON read and unmarshal JSON content from a file into the provided target.
func LoadJSON(filePath string, target interface{}) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}
	if err := json.Unmarshal(content, target); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %v", err)
	}
	return nil
}

// LoadYAML read and unmarshal YAML content from a file into the provided target.
func LoadYAML(filePath string, target interface{}) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}
	if err := yaml.Unmarshal(content, target); err != nil {
		return fmt.Errorf("failed to unmarshal YAML: %v", err)
	}
	return nil
}
