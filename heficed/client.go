package heficed

import (
	"context"
	"fmt"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
)

const (
	TokenUrl string = "https://iam-proxy.heficed.com/oauth2/token"
	APIUrl   string = "https://api.heficed.com"
)

var (
	tenantId string
)

type Config struct {
	ClientId     string
	ClientSecret string
	TenantId     string
	Scopes       []string
}

type Client struct {
	*http.Client
}

func NewClient(cfg Config) (*Client, error) {
	conf := &clientcredentials.Config{
		ClientID:       cfg.ClientId,
		ClientSecret:   cfg.ClientSecret,
		TokenURL:       TokenUrl,
		Scopes:         cfg.Scopes,
		EndpointParams: nil,
		AuthStyle:      0,
	}

	tenantId = cfg.TenantId

	return &Client{conf.Client(context.Background())}, nil
}

func (c *Client) request(method, path, body string) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s/%s", APIUrl, tenantId, path)
	fmt.Println(url)
	if method == "GET" {
		resp, err := c.Get(url)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}

	//TODO: implement POST

	return nil, nil
}
