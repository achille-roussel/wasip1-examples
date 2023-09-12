package main

import (
	"fmt"
	"log"
	"net"

	"github.com/stealthrocket/net/wasip1"
)

func init() {
	net.DefaultResolver.Dial = wasip1.DialContext
}

func main() {
	addrs, err := net.LookupHost("github.com")
	if err != nil {
		log.Fatal("ERROR:", err)
	}
	for _, addr := range addrs {
		fmt.Println(addr)
	}
}
