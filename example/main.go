package main

import (
	"fmt"
	"github.com/vdksystem/heficed-sdk-go/protocompute"
	"log"
	"os"
)

func main() {
	protos, err := protocompute.New(
		os.Getenv("HEFICED_TENANT_ID"),
		os.Getenv("HEFICED_CLIENT_ID"),
		os.Getenv("HEFICED_CLIENT_SECRET"),
	)
	if err != nil {
		log.Fatal(err)
	}

	//r := protos.ListInstances()
	ins := protos.GetInstance(290982)
	fmt.Println(ins)
}
