package main

import (
	"log"

	"github.com/bgpat/dobutsu/shogi"
)

var a [3]map[string]*shogi.Board
var n []*shogi.Board

func main() {
	a[1] = make(map[string]*shogi.Board)
	a[2] = make(map[string]*shogi.Board)
	n = make([]*shogi.Board, 1)
	s := "A1 g2, B1 l2, C1 e2, A2 --, B2 c2, C2 --, A3 --, B3 c1, C3 --, A4 e1, B4 l1, C4 g1, "
	a[1][s] = shogi.NewBoard(s, 1)
	n[0] = a[1][s]
	m := 8
	for i := 0; i < m; i++ {
		log.Printf("%d: %d\n", i, len(n))
		rep(i)
	}
	log.Printf("%d: %d\n", m, len(n))
}

func rep(i int) {
	m := make([]*shogi.Board, 0)
	for _, b := range n {
		b.GenerateNext()
		for _, c := range b.Next {
			if c.Result() > 0 {
				continue
			}
			s := c.ToString()
			if d, ok := a[c.Player][s]; ok {
				d.Previous = append(d.Previous, b)
			} else {
				a[c.Player][s] = c
				m = append(m, c)
			}
		}
	}
	n = m
}
