package main

import (
	"log"

	"github.com/bgpat/dobutsu/shogi"
)

var cache [3]map[string]*shogi.Board

func main() {
	cache[1] = make(map[string]*shogi.Board)
	cache[2] = make(map[string]*shogi.Board)
	s := "A1 g2, B1 l2, C1 e2, A2 --, B2 c2, C2 --, A3 --, B3 c1, C3 --, A4 e1, B4 l1, C4 g1, "
	cache[1][s] = shogi.NewBoard(s, 1)
	queue := make([]*shogi.Board, 1)
	queue[0] = cache[1][s]
	m := 8
	for i := 0; i < m; i++ {
		log.Printf("%d: %d\n", i, len(queue))
		q := make([]*shogi.Board, 0)
		for j := 0; j < len(queue); j += 100 {
			q = append(q, rep(queue[j:j+100])...)
		}
		queue = q
	}
	log.Printf("%d: %d\n", m, len(queue))
}

func rep(queue []*shogi.Board) []*shogi.Board {
	q := make([]*shogi.Board, 0)
	for _, b := range queue {
		b.GenerateNext()
		for _, c := range b.Next {
			if c.Result() > 0 {
				continue
			}
			s := c.ToString()
			if d, ok := cache[c.Player][s]; ok {
				d.Previous = append(d.Previous, b)
				continue
			}
			cache[c.Player][s] = c
			q = append(q, c)
		}
	}
	return q
}
