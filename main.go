package main

import (
	"fmt"
	"mars_rover/rover"
	"mars_rover/rover/direction"
)

func main() {
	position := rover.NewPosition(0, 0)
	newRover, err := rover.NewRover(position, direction.Direction("N"))
	if err != nil {
		panic(err)
	}
	instructions := "MMMRMMLMM"
	if err = newRover.Act(instructions); err != nil {
		panic(err)
	}
	fmt.Println(newRover)
}
