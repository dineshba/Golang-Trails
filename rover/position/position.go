package position

type Position struct {
	X int
	Y int
}

func (p Position) IsValid() bool {
	return p.X >= 0 && p.Y >= 0
}
