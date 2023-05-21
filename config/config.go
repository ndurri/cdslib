package config

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

type Content map[string]string

var content Content = Content{}

var NotFound error = errors.New("Config item not found.")

func Load(filename string) error {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, &content)
}

func Get(key string) string {
	value, prs := content[key]
	if !prs {
		log.Printf("FATAL: Request for %s.\n", key)
		log.Fatal(NotFound)
	}
	return value
}
