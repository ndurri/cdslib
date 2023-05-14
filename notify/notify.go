package notify

import (
	"github.com/ndurri/cdslib/config"
	"github.com/ndurri/cdslib/queue"
)

type Config struct {
	EmailSendQueueName string
}

var cfg Config

func Init(content config.Content) error {
	if err := config.Unmarshal(content, &cfg); err != nil {
		return err
	}
	return queue.Init(content)
}

func Post(to string, subject string, body string) error {
	params := MessageAttributes{
		"to":      to,
		"subject": subject,
	}
	return queue.Post(cfg.EmailSendQueueName, body, params)
}
