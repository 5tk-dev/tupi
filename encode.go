package tupi

import (
	"encoding/json"
	"strings"
)

func EncodeStruct(v any) (any, error) {
	j, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	data := map[string]any{}
	err = json.Unmarshal(j, &data)
	if err != nil {
		return nil, err
	}
	newData := map[string]any{}
	for k, v := range data {
		newData[strings.ToLower(k)] = v
	}
	return newData, nil
}
