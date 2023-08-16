package main

import (
	"fmt"
	"github/wry-0313/go-distributed-system/grades"
	"github/wry-0313/go-distributed-system/registry"
	"github/wry-0313/go-distributed-system/service"
  stlog "log"
  "context"
)
func main() {
	host, port := "localhost", "6000"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)

	r := registry.Registration {
		ServiceName: registry.GradingService,
		ServiceURL: serviceAddress,
	}

  ctx, err := service.Start(context.Background(),
    host, port, r, grades.RegisterHandlers)
  if err != nil {
    stlog.Fatal(err)
  }
  <-ctx.Done()
  fmt.Println("shutting down grading service")
}
