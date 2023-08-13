package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := "3000"
	var srv http.Server
	srv.Addr = ":" + port

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1> Hello World</h1>")
	})

	fmt.Printf("server started at http://localhost:%v\n", port)
	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}
