package request

import (
	"github.com/ndurri/cdslib/config"
	"github.com/ndurri/cdslib/queue"
)

type MessageAttributes queue.MessageAttributes

var requestQueue *queue.Queue

func Post(doctype string, eori string, body string, params MessageAttributes) error {
	if requestQueue == nil {
		requestQueue = queue.NewQueue(config.Get("RequestQueue"))
	}
	params["eori"] = eori
	params["doctype"] = doctype
	return requestQueue.Post(body, queue.MessageAttributes(params))
}
