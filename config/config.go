package config

import (
	"encoding/json"
	"os"
)

type Content []byte

func Load(filename string) (Content, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func Unmarshal(content Content, dest interface{}) error {
	return json.Unmarshal(content, dest)
}
