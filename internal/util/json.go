package util

import "encoding/json"

func PettyJson(v any) string {
	str, _ := json.MarshalIndent(v, "", "  ")
	return string(str)
}
