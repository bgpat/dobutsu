package shogi

import (
	"strconv"
	"strings"
)

type Piece struct {
	Position Position
	Kind     string
	Player   int
}

var movements map[string][]Position = map[string][]Position{
	"c": {Position{Y: -1}},
	"e": {Position{X: -1, Y: -1}, Position{X: -1, Y: 1}, Position{X: 1, Y: -1}, Position{X: 1, Y: 1}},
	"l": {
		Position{X: -1, Y: -1}, Position{X: -1, Y: 1}, Position{X: 1, Y: -1}, Position{X: 1, Y: 1},
		Position{X: 0, Y: -1}, Position{X: 0, Y: 1}, Position{X: -1, Y: 0}, Position{X: 1, Y: 0},
	},
	"g": {Position{X: 0, Y: -1}, Position{X: 0, Y: 1}, Position{X: -1, Y: 0}, Position{X: 1, Y: 0}},
	"h": {
		Position{X: -1, Y: -1}, Position{X: 1, Y: -1},
		Position{X: 0, Y: -1}, Position{X: 0, Y: 1}, Position{X: -1, Y: 0}, Position{X: 1, Y: 0},
	},
}

func NewPiece(s string) *Piece {
	t := strings.Split(s, " ")
	if len(t) < 2 || len(t[1]) < 2 {
		return nil
	}
	if t[1] == "--" {
		return &Piece{
			Position: *NewPosition(t[0]),
		}
	}
	player, err := strconv.Atoi(string(t[1][1]))
	if err != nil {
		return nil
	}
	return &Piece{
		Position: *NewPosition(t[0]),
		Kind:     string(t[1][0]),
		Player:   player,
	}
}

func (p *Piece) NextMovements(b *Board) []Movement {
	var r []Movement
	for _, m := range movements[p.Kind] {
		if p.Player == 2 {
			m = m.Invert()
		}
		i := p.Position.Add(m)
		if b.Pieces[i] != nil && b.Pieces[i].Player != p.Player {
			r = append(r, Movement{
				From: p.Position,
				To:   i,
			})
		}
	}
	return r
}

func (p *Piece) Move(m *Movement) *Piece {
	return &Piece{
		Position: m.To,
		Kind:     p.Kind,
		Player:   p.Player,
	}
}

func (p *Piece) ToString() string {
	if p.Player == 0 {
		return p.Position.ToString() + " --"
	}
	return p.Position.ToString() + " " + p.Kind + strconv.Itoa(p.Player)
}

func (p *Piece) Log() string {
	if p.Player == 0 {
		return "  "
	}
	return p.Kind + strconv.Itoa(p.Player)
}
