package request

import (
	"github.com/ndurri/cdslib/config"
	"github.com/ndurri/cdslib/queue"
)

type MessageAttributes queue.MessageAttributes

type Message struct {
	Id string
	Submitter string
	DocType string
	Body string
}

var requestQueue = queue.NewQueue(config.Get("RequestQueue"))

func (m *Message) Post() error {
	params := queue.MessageAttributes{
		"submitter": m.Submitter,
		"doctype": m.DocType,
	}
	id, err := requestQueue.Post(m.Body, queue.MessageAttributes(params))
	if err != nil {
		return err
	}
	m.Id = *id
	return nil
}
