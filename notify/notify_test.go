package notify

import (
	"testing"
	"github.com/ndurri/cdslib/config"
	"strings"
)

const configData = `{
	"QueueURLPrefix": "http://localhost:3000/",
	"MailSendQueue": "MailSendQueue"
}`

func TestNotify(t *testing.T) {
	ins := strings.NewReader(configData)
	err := config.Load(ins)
	if err != nil {
    	t.Error(err)		
	}
	if notifyQueue.URL() != "http://localhost:3000/MailSendQueue" {
		t.Errorf("notifyQueue.URL() = %s expected 'http://localhost:3000/MailSendQueue'", notifyQueue.URL())
	}
}