package helper

import "encoding/json"

func JsonUnmarshal(data any, v any) error {
	dtByte, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(dtByte, v)
}
