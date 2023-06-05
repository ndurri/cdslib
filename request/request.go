package request

import (
	"github.com/ndurri/cdslib/config"
	"github.com/ndurri/cdslib/queue"
)

type MessageAttributes queue.MessageAttributes

var requestQueue *queue.Queue

func Post(doctype string, submitter string, body string, params MessageAttributes) (*string, error) {
	if requestQueue == nil {
		requestQueue = queue.NewQueue(config.Get("RequestQueue"))
	}
	params["submitter"] = submitter
	params["doctype"] = doctype
	return requestQueue.Post(body, queue.MessageAttributes(params))
}
