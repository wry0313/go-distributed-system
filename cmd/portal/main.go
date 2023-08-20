package main

import (
	"context"
	"fmt"
	"github/wry-0313/go-distributed-system/log"
	"github/wry-0313/go-distributed-system/portal"
	"github/wry-0313/go-distributed-system/registry"
	"github/wry-0313/go-distributed-system/service"
	stlog "log"
)

func main() {
	err := portal.ImportTemplates()
	if err != nil {
		stlog.Fatal(err)
	}
	host, port := "localhost", "6000"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)

	r := registry.Registration{
		ServiceName: registry.PortalService,
		ServiceURL:  serviceAddress,
		RequiredServices: []registry.ServiceName{
			registry.LogService,
			registry.GradingService,
		},
		ServiceUpdateUrl: serviceAddress + "/services",
		HeartbeatURL: serviceAddress + "/heartbeat",
	}

	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		r,
		portal.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}
	if logProvider, err:= registry.GetProvider(registry.LogService); err == nil {
		log.SetClientLogger(logProvider, r.ServiceName)
	}
	<-ctx.Done()
	fmt.Println("Shutting down portal.")
}
