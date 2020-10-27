package main

import (
	"context"
	"fmt"
	"github.com/vdksystem/heficed-sdk-go/heficed"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	cfg := heficed.Config{
		ClientId:     os.Getenv("HEFICED_CLIENT_ID"),
		ClientSecret: os.Getenv("HEFICED_CLIENT_SECRET"),
		TenantId:     os.Getenv("HEFICED_TENANT_ID"),
		Scopes:       []string{"protocompute"},
	}
	hfs, err := heficed.NewClient(cfg, ctx)
	if err != nil {
		log.Fatal(err)
	}

	protos := heficed.Protos{Client: hfs}

	//r := protos.ListInstances()
	//fmt.Println(r)
	ins := protos.GetInstance(290982)
	fmt.Println(ins)
}
