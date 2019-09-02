package rover

import "fmt"

// Direction represents the rover direction
type Direction string

type position struct {
	x int
	y int
}

type rover struct {
	position  position
	direction Direction
}

// NewPosition creates and returns new position
func NewPosition(x, y int) position {
	return position{x: x, y: y}
}

// NewRover creates the rover with given position and direction
func NewRover(position position, direction Direction) rover {
	return rover{position: position, direction: direction}
}

// Navigate rover based on given single instruction
func (r *rover) Navigate(instruction rune) error {
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
	if r.direction == "N" {
		r.position.y++
	} else if r.direction == "E" {
		r.position.x++
	} else if r.direction == "W" {
		r.position.x--
	} else if r.direction == "S" {
		r.position.y--
	}
}

func (r *rover) right() {
	if r.direction == "N" {
		r.direction = "E"
	} else if r.direction == "E" {
		r.direction = "S"
	} else if r.direction == "W" {
		r.direction = "N"
	} else if r.direction == "S" {
		r.direction = "W"
	}
}

func (r *rover) left() {
	if r.direction == "N" {
		r.direction = "W"
	} else if r.direction == "E" {
		r.direction = "N"
	} else if r.direction == "W" {
		r.direction = "S"
	} else if r.direction == "S" {
		r.direction = "E"
	}
}
