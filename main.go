package main

import (
	"log"
	"os"
	"os/signal"
	"runtime/pprof"
	"strconv"

	"github.com/bgpat/dobutsu/client"
)

func main() {
	f, err := os.Create("main.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			log.Printf("captured %v, stopping profiler and exiting...", sig)
			pprof.StopCPUProfile()
			os.Exit(1)
		}
	}()

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
