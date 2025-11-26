package tupi

import "encoding/json"

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
	return data, nil

}
