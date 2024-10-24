package helper

import (
	"encoding/json"
	"os"
)

func JsonUnmarshal(data any, v any) error {
	dtByte, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(dtByte, v)
}

func JsonFileUnmarshal(file string, v any) error {
	fileByte, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	return json.Unmarshal(fileByte, v)
}

func JsonMarshal(v any) string {
	respByte, err := json.Marshal(v)
	if err != nil {
		respByte = []byte("")
	}
	return string(respByte)
}
