package main

import (
	"log"
	"strconv"

	"github.com/bgpat/dobutsu/client"
	"github.com/bgpat/dobutsu/shogi"
)

func main() {
	var c client.Client
	res, err := c.Connect("localhost", 4444)
	if err != nil {
		log.Fatalf("failed to connect: %+v\n", err)
	}
	log.Print(res)
	player, _ := strconv.Atoi(string(res[14]))
	log.Printf("I am player%d.\n", player)
	res = c.Command("turn")
	player, _ = strconv.Atoi(string(res[6]))
	log.Printf("Player%d's turn.\n", player)
	res = c.Command("initboard")
	b := shogi.NewBoard(res, player)
	res = c.Command("board")
	b = b.Load(res)
	log.Printf("%s\n%s\n================================\n", b.ToString(), b.Log())
	b.GenerateNext()
	for i, e := range b.Next {
		log.Printf("%s: %s\n%s\n\n", i.ToString(), e.ToString(), e.Log())
		//log.Printf("%s: %s\n", i.ToString(), e.ToString())
	}
}
