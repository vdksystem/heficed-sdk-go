package client

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
)

const (
	TokenUrl string = "https://iam-proxy.heficed.com/oauth2/token"
	APIUrl   string = "https://api.heficed.com"
)

type Config struct {
	TenantId     string
	ClientId     string
	ClientSecret string
	Context      context.Context
}

func (c *Config) Client() (*http.Client, error) {
	conf := &clientcredentials.Config{
		ClientID:       c.ClientId,
		ClientSecret:   c.ClientSecret,
		TokenURL:       TokenUrl,
		Scopes:         []string{},
		EndpointParams: nil,
		AuthStyle:      0,
	}

	return conf.Client(c.Context), nil
}

func (c *Config) Request(method, path, body string) (interface{}, error) {
	url := fmt.Sprintf("%s/%s/%s", APIUrl, c.TenantId, path)
	var resp *http.Response
	var err error
	if method == "GET" {
		client, err := c.Client()
		if err != nil {
			return nil, err
		}
		resp, err = client.Get(url)
		if err != nil {
			return nil, err
		}
	}
	var responseData map[string]json.RawMessage
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		return nil, err
	}

	//TODO: implement POST

	return responseData, nil
}
