package direction

import (
	"mars_rover/rover/position"
)

// Direction represents the rover direction
type Direction interface {
	String() string
	Left() Direction
	Right() Direction
	Move(p position.Position) position.Position
}

type North struct{}

func (North) Move(p position.Position) position.Position {
	return position.Position{X: p.X, Y: p.Y + 1}
}

func (North) Left() Direction {
	return West{}
}
func (North) Right() Direction {
	return East{}
}

func (North) String() string {
	return "N"
}

type West struct{}

func (West) Move(p position.Position) position.Position {
	return position.Position{X: p.X - 1, Y: p.Y}
}

func (West) String() string {
	return "W"
}

func (West) Left() Direction {
	return South{}
}

func (West) Right() Direction {
	return North{}
}

type East struct{}

func (East) Move(p position.Position) position.Position {
	return position.Position{X: p.X + 1, Y: p.Y}
}

func (East) String() string {
	return "E"
}

func (East) Left() Direction {
	return North{}
}

func (East) Right() Direction {
	return South{}
}

type South struct{}

func (South) Move(p position.Position) position.Position {
	return position.Position{X: p.X, Y: p.Y - 1}
}

func (South) String() string {
	return "S"
}

func (South) Left() Direction {
	return East{}
}

func (South) Right() Direction {
	return West{}
}
