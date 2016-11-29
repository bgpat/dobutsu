package main

import (
	"log"

	"github.com/bgpat/dobutsu/client"
)

func main() {
	_, err := client.New("localhost", 4444)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
}
