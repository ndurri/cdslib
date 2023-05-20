package request

import (
	"process/cdslib/config"
	"process/cdslib/queue"
)

type MessageAttributes queue.MessageAttributes

type Config struct {
	RequestQueue string
}

var requestQueue *queue.Queue

func Init(content config.Content) error {
	var cfg Config
	if err := config.Unmarshal(content, &cfg); err != nil {
		return err
	}
	if err := queue.Init(content); err != nil {
		return err
	}
	requestQueue = queue.NewQueue(cfg.RequestQueue)
	return nil
}

func Post(doctype string, eori string, body string, params MessageAttributes) error {
	params["eori"] = eori
	params["doctype"] = doctype
	return requestQueue.Post(body, queue.MessageAttributes(params))
}
