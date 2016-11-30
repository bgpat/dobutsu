package main

import (
	"log"
	"os"
	"strconv"

	"github.com/bgpat/dobutsu/client"
)

func main() {
	host := os.Getenv("host")
	if host == "" {
		host = "localhost"
	}
	port, err := strconv.Atoi(os.Getenv("port"))
	if err != nil {
		port = 4444
	}
	_, err = client.New(host, port)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
}
