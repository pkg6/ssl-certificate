package helper

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

const WorkerPath = ".ssl-certificate"

func HomeDataFile(file string) string {
	home, _ := os.UserHomeDir()
	return path.Join(home, WorkerPath, "data", file)
}

func ReadFromFile(filePath string, v any) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(content, v)
}

func WriteToFile(filePath string, content []byte) error {
	// Get the directory part of the file path
	dir := filepath.Dir(filePath)
	// Check if the directory exists
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
	}
	// Write content to the file
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}
	return nil
}
