package queue

import (
	"github.com/ndurri/cdslib/config"

	"github.com/ndurri/aws/sqs"
)

type MessageAttributes sqs.MessageAttributes
type Message sqs.Message

type Queue struct {
	Name string
	URL  string
}

func NewQueue(name string) *Queue {
	return &Queue{Name: name, URL: config.Get("QueueURLPrefix") + name}
}

func (q *Queue) Post(body string, attributes MessageAttributes) error {
	return sqs.Put(q.URL, body, sqs.MessageAttributes(attributes))
}

func (q *Queue) Get() (*Message, error) {
	for {
		message, err := sqs.Get(q.URL)
		if err != nil {
			return nil, err
		}
		if message != nil {
			return (*Message)(message), nil
		}
	}
}

func (m *Message) Delete() error {
	return (*sqs.Message)(m).Delete()
}
