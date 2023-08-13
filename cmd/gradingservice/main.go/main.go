package main

import (
	"fmt"
	"github/wry-0313/go-distributed-system/registry"
)
func main() {
	host, port := "localhost", "6000"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)

	r := registry.Registration {
		ServiceName: registry.GradingService,
		ServiceURL: serviceAddress,
	}


}