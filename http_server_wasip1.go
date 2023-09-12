package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/stealthrocket/net/wasip1"
)

func main() {
	server := http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello, World!")
		}),
	}

	l, err := wasip1.Listen("tcp", ":3030")
	if err != nil {
		log.Fatal("ERROR:", err)
	}
	defer l.Close()

	server.Serve(l)
}
