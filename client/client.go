package client

import (
	"context"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
)

const TokenUrl string = "https://iam-proxy.heficed.com/oauth2/token"

type Config struct {
	ClientId     string
	ClientSecret string
	Scopes       []string
}

type Client struct {
	*http.Client
}

func New(cfg Config) (*Client, error) {
	ctx := context.Background()
	conf := &clientcredentials.Config{
		ClientID:       cfg.ClientId,
		ClientSecret:   cfg.ClientSecret,
		TokenURL:       TokenUrl,
		Scopes:         cfg.Scopes,
		EndpointParams: nil,
		AuthStyle:      0,
	}

	c := Client{conf.Client(ctx)}

	return &c, nil
}
