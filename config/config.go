package config

import (
	"encoding/json"
	"io"
)

type Config map[string]*string

var config Config = Config{}

func Load(ins io.Reader) error {
	bytes, err := io.ReadAll(ins)
	if err != nil {
		return err
	}
	var loaded map[string]string
	err = json.Unmarshal(bytes, &loaded)
	if err != nil {
		return err
	}
	for key, value := range loaded {
		item := Get(key)
		*item = value
	}
	return nil
}

func Get(key string) *string {
	item, prs := config[key]
	if !prs {
		var newitem string = ""
		config[key] = &newitem
		item = &newitem
	}
	return item
}

