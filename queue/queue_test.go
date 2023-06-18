package queue

import (
	"github.com/ndurri/cdslib/config"
	"testing"
	"strings"
)

var qName = "queue"

func TestNew(t *testing.T) {
	q := NewQueue(&qName)
    if *q.Name != "queue" {
    	t.Errorf("q.Name = %s, expected 'queue'", *q.Name)
    }
}

func TestURL(t *testing.T) {
	ins := strings.NewReader(`{"QueueURLPrefix": "http://localhost:3000/"}`)
	err := config.Load(ins)
	if err != nil {
    	t.Error(err)		
	}
	q := NewQueue(&qName)
    if q.URL() != "http://localhost:3000/queue" {
    	t.Errorf("q.URL = %s, expected 'http://localhost:3000/queue'", q.URL())
    }	
}
