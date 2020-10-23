package protocompute

import (
	"fmt"
	"github.com/vdksystem/heficed-sdk-go/client"
	"net/http"
)

const APIPath = "protocompute"

type protos struct {
	client *client.Client
	path   string
}

func New(tenantId, clientId, clientSecret string) (*protos, error) {

	cfg := client.Config{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		Scopes:       []string{"protocompute"},
	}
	c, err := client.New(cfg)
	if err != nil {
		return &protos{}, err
	}
	p := protos{
		client: c,
		path:   fmt.Sprintf("https://api.heficed.com/%s/%s", tenantId, APIPath),
	}
	return &p, nil
}

func (p *protos) getRequest(path string) *http.Response {
	url := p.path + path
	resp, err := p.client.Get(url)
	if err != nil {
		return nil
	}

	return resp
}
