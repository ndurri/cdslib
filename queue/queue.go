package queue

import (
	"github.com/ndurri/aws/sqs"
	"github.com/ndurri/cdslib/config"
)

type MessageAttributes sqs.MessageAttributes
type Message sqs.Message
type Queue struct {
	Name string
	URL  string
}

type Config struct {
	QueueURLPrefix string
}

var config Config

func Init(content config.Content) error {
	if err := config.Unmarshal(content, &config); err != nil {
		return err
	}
	return sqs.Init()
}

func NewQueue(name string) Queue {
	return Queue{Name: name, URL: config.QueueURLPrefix + name}
}

func (q *Queue) Post(body string, attributes MessageAttributes) error {
	return sqs.Put(q.URL, body, sqs.MessageAttributes(attributes))
}

func (q *Queue) Get() (*Message, error) {
	return sqs.Get(q.URL)
}