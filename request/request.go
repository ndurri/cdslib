package request

import (
	"github.com/ndurri/cdslib/config"
	"github.com/ndurri/cdslib/queue"
)

type MessageAttributes queue.MessageAttributes

type Config struct {
	RequestQueueName string
}

var config Config

func Init(content config.Content) error {
	if err := config.Unmarshal(content, &config); err != nil {
		return err
	}
	return queue.Init(content)
}

func Post(eori string, body string, params MessageAttributes) error {
	params["eori"] = eori
	return queue.Post(config.RequestQueueName, body, params)
}
