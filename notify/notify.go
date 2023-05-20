package notify

import (
	"process/cdslib/config"
	"process/cdslib/queue"
)

type Config struct {
	MailSendQueue string
}

var notifyQueue *queue.Queue

func Init(content config.Content) error {
	var cfg Config
	if err := config.Unmarshal(content, &cfg); err != nil {
		return err
	}
	if err := queue.Init(content); err != nil {
		return err
	}
	notifyQueue = queue.NewQueue(cfg.MailSendQueue)
	return nil
}

func Post(to string, subject string, body string) error {
	params := queue.MessageAttributes{
		"to":      to,
		"subject": subject,
	}
	return notifyQueue.Post(body, params)
}
