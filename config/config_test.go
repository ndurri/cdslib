package config

import (
	"testing"
	"strings"
)

func TestLoad(t *testing.T) {
	r := strings.NewReader(`{"name": "neil"}`)
	err := Load(r)
	if err != nil {
		t.Fatalf("Load() error %q", err)
	}
    if *config["name"] != "neil" {
    	t.Errorf("*config['name'] = %s, expected 'neil'", *config["name"])
    }
    if len(config) != 1 {
    	t.Errorf("len(config) = %v, expected 1", len(config))    	
    }
}

func TestGet(t *testing.T) {
	name := Get("name")
	if *name != "neil" {
		 t.Errorf("Get('name') = %v, expected 'neil'", *name)    	
	}
}