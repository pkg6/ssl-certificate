package helper

import (
	"encoding/json"
	"os"
	"path"
	"path/filepath"
)

const WorkerPath = ".ssl-certificate"

func HomeDataFile(file string) string {
	home, _ := os.UserHomeDir()
	workerPath := path.Join(home, WorkerPath, "data", file)
	dir := filepath.Dir(workerPath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		_ = os.MkdirAll(dir, os.ModePerm)
	}
	return workerPath
}

func ReadFromFile(filePath string, v any) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(content, v)
}

func WriteToFile(filePath string, content []byte) error {
	return os.WriteFile(filePath, content, 0644)
}
