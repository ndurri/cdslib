package request

import (
	"testing"
	"github.com/ndurri/cdslib/config"
	"strings"
)

const configData = `{
	"QueueURLPrefix": "http://localhost:3000/",
	"RequestQueue": "RequestQueue"
}`

func TestNotify(t *testing.T) {
	ins := strings.NewReader(configData)
	err := config.Load(ins)
	if err != nil {
    	t.Error(err)		
	}
	if requestQueue.URL() != "http://localhost:3000/RequestQueue" {
		t.Errorf("notifyQueue.URL() = %s expected 'http://localhost:3000/RequestQueue'", requestQueue.URL())
	}
}