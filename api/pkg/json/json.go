package json

import (
	"encoding/json"
	"io"
)

func ReadJSON(r io.ReadCloser, dst any) error {
	dec := json.NewDecoder(r)
	err := dec.Decode(dst)
	if err != nil {
		return err
	}

	return nil
}

func WriteJSON(data any) ([]byte, error) {
	return json.Marshal(data)
}
