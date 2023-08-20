package main

import (
	"context"
	"fmt"
	"github/wry-0313/go-distributed-system/log"
	"github/wry-0313/go-distributed-system/registry"
	"github/wry-0313/go-distributed-system/service"
	stlog "log"
)

func main() {
	log.New("./distributed.log")
	host, port := "localhost", "4000"
	serverAddress := fmt.Sprintf("http://%s:%s", host, port)
	r := registry.Registration{
		ServiceName: registry.LogService, 
		ServiceURL:  serverAddress,
		RequiredServices: make([]registry.ServiceName, 0),
		ServiceUpdateUrl: serverAddress + "/services",
	}
	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		r,
		log.RegisterHandlers,
	)
	if err != nil {
		fmt.Println("Error starting log service")
		stlog.Fatal(err)
	}

	<-ctx.Done()
	fmt.Println("Shutting down log service")
}
