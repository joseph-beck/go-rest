package firestore

import (
	"encoding/json"
)

func MapStruct(s interface{}) (map[string]interface{}, error) {
	var m map[string]interface{}
	mar, err := json.Marshal(s)
	if err != nil {
		return nil, ErrMappingStruct
	}
	json.Unmarshal(mar, &m)

	return m, nil
}
