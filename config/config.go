package config

import (
	"encoding/json"
	"io"
)

type ConfigItem string
type Config map[string]*ConfigItem

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
		*item = ConfigItem(value)
	}
	return nil
}

func Get(key string) *ConfigItem {
	item, prs := config[key]
	if !prs {
		var newitem ConfigItem = ""
		config[key] = &newitem
		item = &newitem
	}
	return item
}

