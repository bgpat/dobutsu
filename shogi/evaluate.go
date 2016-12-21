package shogi

type Evaluation struct {
	Result int
	Loop   bool
	Point  int
	Depth  int
}

func (a *Board) Less(b *Board) bool {
	if a.Evaluation.Result > 0 {
		return a.Evaluation.Result == a.Player
	}
	if b.Evaluation.Result > 0 {
		return b.Evaluation.Result != b.Player
	}
	if a.Evaluation.Loop && !b.Evaluation.Loop {
		return a.Evaluation.Point >= 0
	}
	if !a.Evaluation.Loop && b.Evaluation.Loop {
		return b.Evaluation.Point < 0
	}
	return a.Evaluation.Point < b.Evaluation.Point
}

func (b *Board) Evaluate(depth int) {
	result := b.Result()
	loop := b.IsLoop()
	if result == 0 && !loop && b.Next != nil {
		var max *Board
		for _, n := range b.Next {
			if n.Evaluation == nil || n.Evaluation.Depth < depth {
				//if n.Evaluation == nil {
				n.Evaluate(depth - 1)
			}
			if max == nil || max.Less(n) {
				max = n
			}
		}
		b.Evaluation = &Evaluation{
			Result: max.Evaluation.Result,
			Loop:   max.Evaluation.Loop,
			Point:  -max.Evaluation.Point,
			Depth:  max.Evaluation.Depth + 1,
		}
		return
	}
	b.Evaluation = &Evaluation{
		Result: result,
		Loop:   loop,
		Point:  b.EvalPoint(),
		Depth:  0,
	}
}

func (b *Board) EvalPoint() int {
	point := 0
	for _, p := range b.Pieces {
		if p.Player == b.Player {
			point--
		} else if p.Player != 0 {
			point++
		}
	}
	return point
}

func (b *Board) IsLoop() bool {
	if b.Count == nil {
		return false
	}
	return b.Count[b.ToString()] > 0
}
