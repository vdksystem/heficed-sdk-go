package main

import (
	"context"
	"fmt"
	"github.com/vdksystem/heficed-sdk-go/client"
	"github.com/vdksystem/heficed-sdk-go/protocompute"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	cfg := client.Config{
		ClientId:     os.Getenv("HEFICED_CLIENT_ID"),
		ClientSecret: os.Getenv("HEFICED_CLIENT_SECRET"),
		TenantId:     os.Getenv("HEFICED_TENANT_ID"),
		Scopes:       []string{"protocompute"},
	}
	hfs, err := client.New(cfg, ctx)
	if err != nil {
		log.Fatal(err)
	}

	protos := protocompute.Protos{Client: hfs}

	r := protos.ListInstances()
	fmt.Println(r)
	ins := protos.GetInstance(290982)
	fmt.Println(ins)
}
