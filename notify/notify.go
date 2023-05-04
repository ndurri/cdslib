package notify

import (
	"github.com/ndurri/golib/service"
	"strings"
)

var ServiceURL string

func Post(to string, subject string, body string) error {
	params := service.NVP{
		"to": to,
		"subject": subject,
	}
	_, err := service.Post(ServiceURL, params, nil, nil, strings.NewReader(body))
	return err
}