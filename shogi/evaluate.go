package shogi

type Evaluation struct {
	Result int
	Loop   bool
	Point  int
	Depth  int
}

func (a *Board) Less(b *Board) bool {
	if b.Evaluation.Result > 0 {
		return b.Evaluation.Result != b.Player
	}
	if a.Evaluation.Result > 0 {
		return a.Evaluation.Result == a.Player
	}
	if a.Evaluation.Loop && !b.Evaluation.Loop {
		return true
	}
	return a.Evaluation.Point < b.Evaluation.Point
}

func (b *Board) Evaluate() {
	result := b.Result()
	loop := b.IsLoop()
	if result == 0 && !loop && len(b.Next) > 0 {
		var max *Board
		for _, n := range b.Next {
			if n.Evaluation == nil {
				n.Evaluate()
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
