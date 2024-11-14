package data

import (
	"encoding/json"
	"os"
)

type LocalData[T any] struct {
	FileName string
}

func NewLocalData[T any](fileName string) *LocalData[T] {
	return &LocalData[T]{fileName}
}

func (t *LocalData[T]) Save(d T) error {
	data, err := json.Marshal(d)
	if err != nil {
		return err
	}
	return os.WriteFile(t.FileName, data, 0644)
}

func (t *LocalData[T]) Load() (T, error) {
	var data T
	content, err := os.ReadFile(t.FileName)
	if err != nil {
		return data, err
	}
	if err := json.Unmarshal(content, &data); err != nil {
		return data, err
	}
	return data, nil
}
