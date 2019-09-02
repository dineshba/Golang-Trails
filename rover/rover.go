package rover

import (
	"fmt"
)

type rover struct {
	position  Position
	direction Direction
}

// NewPosition.Position creates and returns new position
func NewPosition(x, y int) Position {
	return Position{X: x, Y: y}
}

// NewRover creates the rover with given position and direction
func NewRover(position Position, direction Direction) (rover, error) {
	if !position.IsValid() {
		return rover{}, fmt.Errorf("invalid initial position %v", position)
	}
	return rover{position: position, direction: direction}, nil
}

func (r rover) navigate(instruction Instruction) rover {
	return instruction.Action(r)
}

func (r rover) move() rover {
	r.position = r.direction.Move(r.position)
	return r
}

func (r rover) right() rover {
	r.direction = r.direction.Right()
	return r
}

func (r rover) left() rover {
	r.direction = r.direction.Left()
	return r
}

// Act parses and navigates the rover
func (r rover) Act(instructions string) (rover, error) {
	parseInstructions, err := ParseInstructions(instructions)
	if err != nil {
		return rover{}, err
	}
	rover := r
	for _, instruction := range parseInstructions {
		rover = rover.navigate(instruction)
	}
	return rover, nil
}

func (r rover) String() string {
	return fmt.Sprintf("Postion: %v, Direction: %v", r.position, r.direction)
}
