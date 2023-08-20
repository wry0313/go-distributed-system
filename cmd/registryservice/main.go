package main

import (
	"context"
	"fmt"
	"github/wry-0313/go-distributed-system/registry"
	"log"
	"net/http"
	"time"
)

func main() {
	registry.SetupRegistryService()
	http.Handle("/services", &registry.RegistryService{})

	var srv http.Server
	srv.Addr = registry.ServerPort

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Println("Registry service started. Press any key to stop.")
		var s string
		fmt.Scanln(&s)

		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdownCancel()

		err := srv.Shutdown(shutdownCtx)
		if err != nil {
			log.Println(err)
		}
		cancel()
	}()

	<-ctx.Done()
	fmt.Println("Shutting down registry service")
}
