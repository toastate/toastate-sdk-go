package apiclient

import (
	"bytes"
	"encoding/json"
)

func marshalBody(body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	b := new(bytes.Buffer)
	enc := json.NewEncoder(b)
	enc.SetEscapeHTML(false)
	err := enc.Encode(body)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
