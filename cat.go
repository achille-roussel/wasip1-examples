package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	flag.Parse()

	for _, arg := range flag.Args() {
		if err := cat(arg); err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		}
	}
}

func cat(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(os.Stdout, f)
	return err
}
