package main

import (
	"fmt"
	"os"
)

func main() {
	rover := rover{position: position{x: 0, y: 0}, direction: direction("N")}
	instructions := "MMMRMMLMM"
	for _, instruction := range instructions {
		err := rover.navigate(instruction)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	fmt.Println(rover)
}

type direction string

type position struct {
	x int
	y int
}

type rover struct {
	position  position
	direction direction
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
		return fmt.Errorf("not a valid instruction %c", instruction)
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
