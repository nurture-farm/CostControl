package httpclient

import (
	"github.com/nurture-farm/costcontrol/util"
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	Client = &http.Client{
		Timeout: util.Config.Timeout,
	}
}
