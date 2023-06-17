package config

import (
	"testing"
	"strings"
)

func TestLoad(t *testing.T) {
	r := strings.NewReader(`{"name": "neil", "size": "7"}`)
	err := Load(r)
	if err != nil {
		t.Fatalf("Load() error %q", err)
	}
    if *config["name"] != "neil" {
    	t.Errorf("*config['name'] = %s, want 'neil'", *config["name"])
    }
    if *config["size"] != "7" {
    	t.Errorf("*config['size'] = %s, want '7'", *config["size"])
    }
    if len(config) != 2 {
    	t.Errorf("len(config) = %v, want 2", len(config))    	
    }
}

func TestGet(t *testing.T) {
	name := Get("name")
	if *name != "neil" {
		 t.Errorf("Get('name') = %v, want 'neil'", *name)    	
	}
}