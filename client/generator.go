package client

import (
	"log"
	"sync"

	"github.com/bgpat/dobutsu/shogi"
)

func (c *Client) Generate(depth int) {
	if c.Queue == nil {
		c.Queue = make(map[string]*shogi.Board)
		c.Queue[c.Board.ToString()] = c.Board
	}
	q := make(map[string]*shogi.Board)
	mu := new(sync.RWMutex)
	wg := new(sync.WaitGroup)
	for _, b := range c.Queue {
		wg.Add(1)
		go func(b *shogi.Board) {
			b.GenerateNext()
			for m, n := range b.Next {
				s := n.ToString()
				mu.Lock()
				if e, ok := q[s]; ok {
					b.Next[m] = e
					mu.Unlock()
					continue
				}
				q[s] = n
				mu.Unlock()
			}
			wg.Done()
		}(b)
	}
	wg.Wait()
	log.Printf("generated %d (%d => %d)\n", depth, len(c.Queue), len(q))
	c.Queue = q
	depth--
	if depth > 0 {
		c.Generate(depth)
	}
}
