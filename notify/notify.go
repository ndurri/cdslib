package notify

import (
	"github.com/ndurri/cdslib/config"

	"github.com/ndurri/cdslib/queue"
)

var notifyQueue *queue.Queue

func Post(to string, subject string, body string) (*string, error) {
	if notifyQueue == nil {
		notifyQueue = queue.NewQueue(config.Get("MailSendQueue"))
	}
	params := queue.MessageAttributes{
		"to":      to,
		"subject": subject,
	}
	return notifyQueue.Post(body, params)
}
