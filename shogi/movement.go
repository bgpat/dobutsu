package shogi

type Movement struct {
	From Position
	To   Position
}

func NewMovement(from, to Position) *Movement {
	return &Movement{
		From: from,
		To:   to,
	}
}

func (m *Movement) ToString() string {
	return "mv " + m.From.ToString() + " " + m.To.ToString()
}
