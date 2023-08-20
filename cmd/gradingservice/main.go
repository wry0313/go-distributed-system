package main

import (
	"context"
	"fmt"
	"github/wry-0313/go-distributed-system/grades"
	"github/wry-0313/go-distributed-system/log"
	"github/wry-0313/go-distributed-system/registry"
	"github/wry-0313/go-distributed-system/service"
	stlog "log"
)
func main() {
	host, port := "localhost", "5001"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)

	r := registry.Registration {
		ServiceName: registry.GradingService,
		ServiceURL: serviceAddress,
		RequiredServices: []registry.ServiceName{registry.LogService},
		ServiceUpdateUrl: serviceAddress + "/services",
		HeartbeatURL: serviceAddress + "/heartbeat",
	}

  ctx, err := service.Start(context.Background(),
    host, port, r, grades.RegisterHandlers)
  if err != nil {
    stlog.Fatal(err)
  }
  if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
	fmt.Printf("Logging service found at: %s\n", logProvider)
	log.SetClientLogger(logProvider, r.ServiceName)
  }
  <-ctx.Done()
  fmt.Println("shutting down grading service")
}
