package notify

import (
	"github.com/ndurri/cdslib/config"

	"github.com/ndurri/cdslib/queue"
)

type Message struct {
	To string
	Subject string
	Body string
}

var notifyQueue = queue.NewQueue(config.Get("MailSendQueue"))

func (m *Message) Post() (*string, error) {
	params := queue.MessageAttributes{
		"to":      m.To,
		"subject": m.Subject,
	}
	return notifyQueue.Post(m.Body, params)
}
