package client

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/bgpat/dobutsu/shogi"
)

func (c *Client) Step() error {
	switch c.Phase {
	case "connected":
		res := c.Read()
		player, err := strconv.Atoi(string(res[14]))
		if err != nil {
			return err
		}
		c.Player = player
		c.Phase = "initialize"
	case "turn":
		res := c.Command("turn")
		player, err := strconv.Atoi(string(res[6]))
		if err != nil {
			break
		}
		turn := c.Turn
		c.Turn = player
		c.Board.Player = c.Turn
		if turn == c.Turn {
			time.Sleep(time.Second)
			break
		}
		if turn == 0 {
			c.Phase = "update"
			break
		}
		c.Phase = "board"
	case "board":
		res := c.Command("board")
		for m, b := range c.Board.Next {
			s := b.ToString()
			if res == s+"\n" {
				c.Board = b
				log.Println(m.ToString() + "\n" + b.Log())
				c.Count[s]++
				c.Phase = "update"
				return nil
			}
		}
		return errors.New("fault: " + res)
	case "update":
		if c.Board.Result() > 0 {
			c.Phase = "result"
			break
		}
		c.Board.GenerateNext()
		if c.Turn == c.Player {
			c.Phase = "move"
		} else {
			c.Phase = "turn"
		}
	case "move":
		var board *shogi.Board
		var movement shogi.Movement
		c.Queue = nil
		c.Generate(c.Depth)
		for m, b := range c.Board.Next {
			if b.Evaluation == nil {
				b.Evaluate()
			}
			if board == nil || board.Less(b) {
				board = b
				movement = m
			}
		}
		if board == nil {
			c.Phase = ""
			break
		}
		c.Command(movement.ToString())
		c.Queue = nil
		c.Phase = "turn"
	case "initialize":
		res := c.Command("initboard")
		c.Board = shogi.NewBoard(res, 1)
		c.Board.Count = c.Count
		c.Phase = "load"
	case "load":
		res := c.Command("board")
		c.Board = c.Board.Load(res)
		c.Phase = "turn"
	case "result":
		log.Printf("Result: %d\n", c.Board.Result())
		c.Phase = ""
	}
	return nil
}
