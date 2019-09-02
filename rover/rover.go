package rover

import (
	"fmt"
	"mars_rover/rover/direction"
	"mars_rover/rover/position"
)

type rover struct {
	position  position.Position
	direction direction.Direction
}

// NewPosition.Position creates and returns new position
func NewPosition(x, y int) position.Position {
	return position.Position{X: x, Y: y}
}

// NewRover creates the rover with given position and direction
func NewRover(position position.Position, direction direction.Direction) (rover, error) {
	if !position.IsValid() {
		return rover{}, fmt.Errorf("invalid initial position %v", position)
	}
	return rover{position: position, direction: direction}, nil
}

func (r *rover) navigate(instruction rune) error {
	switch instruction {
	case 'M':
		r.move()
	case 'L':
		r.left()
	case 'R':
		r.right()
	default:
		return fmt.Errorf("Not a valid instruction %c", instruction)
	}
	return nil
}

func (r *rover) move() {
	r.position = r.direction.Move(r.position)
}

func (r *rover) right() {
	r.direction = r.direction.Right()
}

func (r *rover) left() {
	r.direction = r.direction.Left()
}

// Act parses and navigates the rover
func (r *rover) Act(instructions string) error {
	for _, instruction := range instructions {
		err := r.navigate(instruction)
		if err != nil {
			return err
		}
	}
	return nil
}
