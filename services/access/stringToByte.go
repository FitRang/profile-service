package access

import "encoding/json"

func stringToByte(s string) []byte {
	return []byte(s)
}

func jsonToBytes(v any) ([]byte, error) {
	return json.Marshal(v)
}
