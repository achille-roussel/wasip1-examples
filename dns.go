package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addrs, err := net.LookupHost("github.com")
	if err != nil {
		log.Fatal("ERROR:", err)
	}
	for _, addr := range addrs {
		fmt.Println(addr)
	}
}
