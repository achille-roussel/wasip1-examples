package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	r, err := http.Get("https://eo3qh2ncelpc9q0.m.pipedream.net")
	if err != nil {
		log.Fatal("ERROR:", err)
	}
	defer r.Body.Close()

	if _, err := io.Copy(os.Stdout, r.Body); err != nil {
		log.Fatal("ERROR:", err)
	}
}
