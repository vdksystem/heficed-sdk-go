package heficed

import (
	"bytes"
	"context"
	"fmt"
	"golang.org/x/oauth2/clientcredentials"
	"io/ioutil"
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

func (c *Client) request(method, path string, body []byte) (*http.Response, error) {
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

func (c *Client) mock_request(method, path string, body []byte) (*http.Response, error) {
	b := []byte(`{
    "id": 619153,
    "dateCreated": 1603756800,
    "dateDue": 1603756800,
    "datePaid": 1603813887,
    "credit": 0.0,
    "transactions": [
        {
            "id": 325171,
            "paymentMethod": "Credit Card",
            "date": 1603813887,
            "description": "Invoice Payment",
            "in": 97.11,
            "out": 0.0,
            "fees": 0.0,
            "ref": "ch_1HguSIDhDdRyJ3P1YC2At0Bs",
            "invoiceId": 619153
        }
    ],
    "items": [
        {
            "total": 79.0,
            "description": "Proto Compute - worker2.us-east-02.supresonic.nc (27/10/2020 - 27/11/2020)",
            "instanceId": 294677,
            "configuration": {
                "nice": {
                    "Product": "Proto Compute",
                    "Flavor": "FRA-D016",
                    "Location": "Frankfurt, Germany"
                },
                "pricing": {
                    "raw": {
                        "locationId": 0
                    },
                    "nice": {
                        "Location": 0
                    }
                },
                "raw": {
                    "product": "protocompute",
                    "flavorId": "FRA-D016",
                    "locationId": "de-fra1"
                }
            },
            "periodStart": 1603756800,
            "periodEnd": 1606435200
        },
        {
            "total": 0.6,
            "description": "Heficed Connect DE-FRA1 - WJMhFphH31EjLe6d (27/10/2020 - 27/11/2020)",
            "instanceId": 294676,
            "configuration": {
                "nice": {
                    "Product": "Connect",
                    "Location": "Frankfurt, Germany",
                    "Bandwidth": "1 TB"
                },
                "pricing": {
                    "raw": {
                        "bandwidth": 0.6,
                        "locationId": 0
                    },
                    "nice": {
                        "Bandwidth": 0.6,
                        "Location": 0
                    }
                },
                "raw": {
                    "product": "connect",
                    "location": "de-fra1",
                    "bandwidth": 1
                }
            },
            "periodStart": 1603756800,
            "periodEnd": 1606435200
        }
    ],
    "currency": "USD",
    "total": 97.11,
    "subtotal": 79.6,
    "tax": 17.51,
    "taxrate": 22.0,
    "leftToPay": 0.0,
    "status": "Paid",
    "billingDetails": {
        "firstName": "Sandro",
        "lastName": "Modarelli",
        "company": "Namecheap",
        "email": "sandro.modarelli@namecheap.com",
        "address": "via Goldoni 181",
        "address2": "",
        "state": "MO",
        "city": "Vignola",
        "country": "Italy",
        "postCode": "41058",
        "companyNumber": "",
        "vatNumber": ""
    },
    "asyncTaskIds": [],
    "usageBasedAvailable": null,
    "usageBasedRate": null
	}`)

	r := ioutil.NopCloser(bytes.NewReader([]byte(b)))
	resp := http.Response{
		Status:     "OK",
		StatusCode: 200,
		Body:       r,
	}
	return &resp, nil
}
