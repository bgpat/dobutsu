package main

import (
	"log"
	"os"
	"strconv"

	"github.com/bgpat/dobutsu/client"
)

func main() {
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 4444
	}
	depth, err := strconv.Atoi(os.Getenv("DEPTH"))
	if err != nil {
		depth = 6
	}
	_, err = client.New(host, port, depth)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
}
