package request

import (
	"github.com/ndurri/cdslib/config"
	"github.com/ndurri/cdslib/queue"
)

type MessageAttributes queue.MessageAttributes

type Config struct {
	RequestQueueName string
}

var cfg Config

func Init(content config.Content) error {
	if err := config.Unmarshal(content, &cfg); err != nil {
		return err
	}
	return queue.Init(content)
}

func Post(doctype string, eori string, body string, params MessageAttributes) error {
	params["eori"] = eori
	params["doctype"] = doctype
	return queue.Post(cfg.RequestQueueName, body, params)
}
