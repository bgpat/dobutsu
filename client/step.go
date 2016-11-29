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
			log.Printf(" @%s %s\n", m.ToString(), b.ToString())
			if res == b.ToString()+"\n" {
				c.Board = b
				c.Phase = "update"
				return nil
			}
		}
		return errors.New("fault: " + res)
	case "update":
		c.Board.GenerateNext()
		if c.Turn == c.Player {
			c.Phase = "move"
		} else {
			c.Phase = "turn"
		}
	case "move":
		for m, _ := range c.Board.Next {
			c.Command(m.ToString())
			c.Phase = "turn"
			return nil
		}
		c.Phase = ""
	case "initialize":
		res := c.Command("initboard")
		c.Board = shogi.NewBoard(res, 1)
		c.Phase = "load"
	case "load":
		res := c.Command("board")
		c.Board = c.Board.Load(res)
		c.Phase = "turn"
	}
	return nil
}
