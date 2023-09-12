package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: ":3030",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello, World!")
		}),
	}

	server.ListenAndServe()
}
