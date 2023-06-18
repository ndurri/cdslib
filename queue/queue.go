package queue

import (
	"github.com/ndurri/cdslib/config"

	"github.com/ndurri/aws/sqs"
)

type MessageAttributes sqs.MessageAttributes
type Message sqs.Message

type Queue struct {
	Name string
}

var (
	queueURLPrefix *string = config.Get("QueueURLPrefix")
	put = sqs.Put
	get = sqs.Get
)

func NewQueue(name string) *Queue {
	return &Queue{Name: name}
}

func (q *Queue) URL() string {
	return *queueURLPrefix + q.Name
}

func (q *Queue) Post(body string, attributes MessageAttributes) (*string, error) {
	return put(q.URL(), body, sqs.MessageAttributes(attributes))
}

func (q *Queue) Get() (*Message, error) {
	for {
		message, err := get(q.URL())
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
