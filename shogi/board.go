package shogi

import (
	"strings"
)

type Board struct {
	Next       map[Movement]*Board
	Player     int
	Pieces     map[Position]*Piece
	Width      int
	Height     int
	Count      map[string]int
	Evaluation *Evaluation
}

func NewBoard(s string, player int) *Board {
	b := &Board{
		Player: player,
		Width:  0,
		Height: 0,
	}
	b = b.Load(s)
	for _, p := range b.Pieces {
		if b.Width < p.Position.X {
			b.Width = p.Position.X
		}
		if b.Height < p.Position.Y {
			b.Height = p.Position.Y
		}
	}
	b.Width++
	b.Height++
	return b
}

func (b *Board) Load(s string) *Board {
	t := strings.Split(s, ", ")
	pieces := make(map[Position]*Piece)
	for _, s := range t {
		p := NewPiece(s)
		if p == nil {
			continue
		}
		pieces[p.Position] = p
	}
	return &Board{
		Player: b.Player,
		Pieces: pieces,
		Width:  b.Width,
		Height: b.Height,
	}
}

func (b *Board) Clone() *Board {
	pieces := make(map[Position]*Piece)
	for p, q := range b.Pieces {
		pieces[p] = q
	}
	return &Board{
		Next:   make(map[Movement]*Board),
		Player: b.Player,
		Pieces: pieces,
		Width:  b.Width,
		Height: b.Height,
	}
}

func (b *Board) NextTurn() *Board {
	c := b.Clone()
	c.Player = b.Player ^ 3
	return c
}

func (b *Board) GetHand(player int) []*Piece {
	var p []*Piece
	i := Position{
		X: b.Width + player - 1,
		Y: 0,
	}
	for b.Pieces[i] != nil {
		p = append(p, b.Pieces[i])
		i.Y++
	}
	return p
}

func (b *Board) Catch(p *Piece, player int) {
	i := Position{
		X: b.Width + player - 1,
		Y: 0,
	}
	for _, ok := b.Pieces[i]; ok; _, ok = b.Pieces[i] {
		i.Y++
	}
	b.Pieces[i] = &Piece{
		Position: i,
		Kind:     p.Kind,
		Player:   player,
	}
	if p.Kind == "h" {
		b.Pieces[i].Kind = "c"
	}
}

func (b *Board) Move(m *Movement) *Board {
	c := b.NextTurn()
	if c.Pieces[m.To].Player != 0 {
		c.Catch(c.Pieces[m.To], b.Player)
	}
	c.Pieces[m.To] = c.Pieces[m.From].Move(m)
	if c.Pieces[m.To].Kind == "c" && m.To.IsEdge(b) {
		c.Pieces[m.To].Kind = "h"
	}
	c.Pieces[m.From] = &Piece{Position: m.From}
	return c
}

func (b *Board) Put(m *Movement) *Board {
	c := b.NextTurn()
	c.Pieces[m.To] = &Piece{
		Position: m.To,
		Kind:     c.Pieces[m.From].Kind,
		Player:   b.Player,
	}
	i := Position{
		X: m.From.X,
		Y: m.From.Y,
	}
	j := Position{
		X: m.From.X,
		Y: m.From.Y + 1,
	}
	for c.Pieces[i] != nil {
		if _, ok := c.Pieces[j]; ok {
			c.Pieces[i] = &Piece{
				Position: i,
				Kind:     c.Pieces[j].Kind,
				Player:   b.Player,
			}
			i = j
			j.Y++
		} else {
			delete(c.Pieces, i)
		}
	}
	return c
}

func (b *Board) GenerateNext() {
	if len(b.Next) > 0 {
		return
	}
	b.Next = make(map[Movement]*Board)
	for _, p := range b.Pieces {
		if p.Player == b.Player {
			if p.Position.X < b.Width {
				for _, m := range p.NextMovements(b) {
					b.Next[m] = b.Move(&m)
				}
			} else {
				for _, q := range b.Pieces {
					if q.Player > 0 {
						continue
					}
					m := NewMovement(p.Position, q.Position)
					b.Next[*m] = b.Put(m)
				}
			}
		}
	}
}

func (b *Board) Result() int {
	another := 1
	if b.Player == 1 {
		another = 2
	}
	c := b.Catched(b.Player)
	if c > 0 {
		return c
	}
	c = b.Catched(another)
	if c > 0 {
		return c
	}
	return b.Tried()
}

func (b *Board) Catched(player int) int {
	i := Position{
		X: b.Width + player - 1,
		Y: 0,
	}
	for p, ok := b.Pieces[i]; ok; p, ok = b.Pieces[i] {
		if p.Kind == "l" || p.Kind == "L" {
			return player
		}
		i.Y++
	}
	return 0
}

func (b *Board) Tried() int {
	i := Position{
		X: 0,
		Y: 0,
	}
	if b.Player == 2 {
		i.Y = b.Height - 1
	}
	for i.X < b.Width {
		p := b.Pieces[i]
		if p.Player == b.Player && (p.Kind == "l" || p.Kind == "L") {
			return b.Player
		}
		i.X++
	}
	return 0
}

func (b *Board) ToString() string {
	s := ""
	i := Position{
		X: 0,
		Y: 0,
	}
	for i.Y < b.Height {
		for i.X < b.Width {
			if b.Pieces[i] != nil {
				s += b.Pieces[i].ToString() + ", "
			}
			i.X++
		}
		i.X = 0
		i.Y++
	}
	for _, p := range b.GetHand(1) {
		s += p.ToString() + ", "
	}
	for _, p := range b.GetHand(2) {
		s += p.ToString() + ", "
	}
	return s
}

func (b *Board) Log() string {
	t := make([][]string, b.Height)
	for y := 0; y < b.Height; y++ {
		t[y] = make([]string, b.Width)
		for x := 0; x < b.Width; x++ {
			i := Position{
				X: x,
				Y: y,
			}
			t[y][x] = b.Pieces[i].Log()
		}
	}
	s := make([]string, b.Height)
	for i, r := range t {
		s[i] = "| " + strings.Join(r, " | ") + " |"
	}
	v := "\n" + strings.Repeat("+----", b.Width) + "+\n"
	var h1 []string
	for _, p := range b.GetHand(1) {
		h1 = append(h1, p.Log())
	}
	var h2 []string
	for _, p := range b.GetHand(2) {
		h2 = append(h2, p.Log())
	}
	return strings.Join(h2, ", ") + v + strings.Join(s, v) + v + strings.Join(h1, ", ")
}
