package shogi

import (
	"strconv"
)

type Position struct {
	X int
	Y int
}

func NewPosition(s string) *Position {
	if len(s) < 2 {
		return nil
	}
	y, err := strconv.Atoi(string(s[1]))
	if err != nil {
		return nil
	}
	return &Position{
		X: int(s[0]) - 'A',
		Y: y - 1,
	}
}

func (p *Position) Add(q Position) Position {
	return Position{
		X: p.X + q.X,
		Y: p.Y + q.Y,
	}
}

func (p *Position) Invert() Position {
	return Position{
		X: -p.X,
		Y: -p.Y,
	}
}

func (p *Position) ToString() string {
	return string(byte(p.X+'A')) + strconv.Itoa(p.Y+1)
}
