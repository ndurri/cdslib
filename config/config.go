package config

import (
	"os"
	"encoding/json"
)

func Load(file string, resp *interface{}) error {
	content, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	return json.Unmarshal(content, resp)
}