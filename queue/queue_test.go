package queue

import (
	"github.com/ndurri/cdslib/config"
	"testing"
	"strings"
	"github.com/ndurri/aws/sqs"
)

func fakePut(url string, body string, attributes sqs.MessageAttributes) (*string, error) {
	var messageId = "messageId"
	return &messageId, nil
}

func fakeGet(url string) (*sqs.Message, error) {
	return &sqs.Message{}, nil
}

func TestNew(t *testing.T) {
	q := NewQueue("queue")
    if q.Name != "queue" {
    	t.Errorf("q.Name = %s, expected 'queue'", q.Name)
    }
}

func TestURL(t *testing.T) {
	ins := strings.NewReader(`{"QueueURLPrefix": "http://localhost:3000/"}`)
	err := config.Load(ins)
	if err != nil {
    	t.Error(err)		
	}
	q := NewQueue("queue")
    if q.URL() != "http://localhost:3000/queue" {
    	t.Errorf("q.URL = %s, expected 'http://localhost:3000/queue'", q.URL())
    }	
}

func TestPost(t *testing.T) {
	ins := strings.NewReader(`{"QueueURLPrefix": "http://localhost:3000/"}`)
	err := config.Load(ins)
	if err != nil {
    	t.Error(err)		
	}
	q := NewQueue("queue")
	put = fakePut
	ma := MessageAttributes{
		"attr1": "value1",
	}
	messageId, err := q.Post("body", ma)
	if err != nil {
		t.Fatal(err)		
	}
	if *messageId != "messageId" {
    	t.Errorf("messageId = %s, expected 'messageId'", *messageId)		
	}
}

func TestGet(t *testing.T) {
	ins := strings.NewReader(`{"QueueURLPrefix": "http://localhost:3000/"}`)
	err := config.Load(ins)
	if err != nil {
    	t.Error(err)		
	}
	q := NewQueue("queue")
	get = fakeGet
	_, err = q.Get()
	if err != nil {
		t.Fatal(err)		
	}
}
