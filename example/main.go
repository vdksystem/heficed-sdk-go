package main

import (
	"fmt"
	"github.com/vdksystem/heficed-sdk-go/heficed"
	"log"
	"os"
)

type createInstanceResponse struct {
	Items []item `json:"items"`
}

type item struct {
	InstanceId int `json:"instanceId"`
}

func main() {
	cfg := heficed.Config{
		ClientId:     os.Getenv("HEFICED_CLIENT_ID"),
		ClientSecret: os.Getenv("HEFICED_CLIENT_SECRET"),
		TenantId:     os.Getenv("HEFICED_TENANT_ID"),
		Scopes:       []string{"protocompute"},
	}
	hfs, err := heficed.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	protos := heficed.Protos{Client: hfs}

	//r := protos.ListInstances()
	//fmt.Println(r)
	//ins := protos.GetInstance(290982)
	//fmt.Println(ins)

	ps, err := protos.CreateInstance(heficed.NewInstance{
		LocationId:    "de-fra1",
		FlavorId:      "FRA-D016",
		Hostname:      "worker3.eu-central-01.supresonic.nc",
		BillingTypeId: 1,
	})

	fmt.Println(ps)
}
