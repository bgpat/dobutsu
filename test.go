package main

import (
	"log"
	"runtime"
	"sync"

	"github.com/bgpat/dobutsu/shogi"
)

var cache [3]map[string]*shogi.Board
var mutex *sync.RWMutex
var wg *sync.WaitGroup

func main() {
	mutex = &sync.RWMutex{}
	wg = &sync.WaitGroup{}
	cache[1] = make(map[string]*shogi.Board)
	cache[2] = make(map[string]*shogi.Board)
	s := "A1 g2, B1 l2, C1 e2, A2 --, B2 c2, C2 --, A3 --, B3 c1, C3 --, A4 e1, B4 l1, C4 g1, "
	cache[1][s] = shogi.NewBoard(s, 1)
	queue := make([]*shogi.Board, 1)
	queue[0] = cache[1][s]
	phase := 10
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)
	for i := 0; i < phase; i++ {
		log.Printf("%d: %d\n", i, len(queue))
		q := make([]*shogi.Board, 0)
		l := len(queue)
		for j := 0; j < cpus; j++ {
			wg.Add(1)
			go rep(queue[l*j/cpus:l*(j+1)/cpus], &q)
		}
		wg.Wait()
		queue = q
	}
	log.Printf("%d: %d\n", phase, len(queue))
}

func rep(queue []*shogi.Board, q *[]*shogi.Board) {
	for _, b := range queue {
		b.GenerateNext()
		for _, c := range b.Next {
			if c.Result() > 0 {
				continue
			}
			s := c.ToString()
			mutex.Lock()
			if d, ok := cache[c.Player][s]; ok {
				d.Previous = append(d.Previous, b)
				mutex.Unlock()
				continue
			}
			cache[c.Player][s] = c
			*q = append(*q, c)
			mutex.Unlock()
		}
	}
	wg.Done()
}
