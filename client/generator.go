package client

import (
	"log"

	"github.com/bgpat/dobutsu/shogi"
)

func (c *Client) Generate(depth int) {
	if c.Queue == nil {
		c.Queue = make(map[string]*shogi.Board)
		c.Queue[c.Board.ToString()] = c.Board
	}
	q := make(map[string]*shogi.Board)
	for _, b := range c.Queue {
		b.GenerateNext()
		for m, n := range b.Next {
			s := n.ToString()
			if e, ok := q[s]; ok {
				b.Next[m] = e
				continue
			}
			q[s] = n
		}
	}
	log.Printf("generated %d (%d => %d)\n", depth, len(c.Queue), len(q))
	c.Queue = q
	if depth > 0 {
		c.Generate(depth - 1)
	}
}
