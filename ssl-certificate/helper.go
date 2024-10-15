package main

import (
	"encoding/json"
	"os"
)

func Load(file string) error {
	fileByte, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(fileByte, &cfg); err != nil {
		return err
	}
	return err
}
