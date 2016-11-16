package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/bgpat/dobutsu/shogi"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	b := shogi.NewBoard("A1 g2, B1 l2, C1 e2, A2 --, B2 c2, C2 --, A3 --, B3 c1, C3 --, A4 e1, B4 l1, C4 g1, ", 1)
	if s.Scan() {
		b = b.Load(s.Text())
		fmt.Println(b.Log())
	}
}
