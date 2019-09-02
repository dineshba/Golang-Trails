package main

import (
	"fmt"
	"mars_rover/rover"
)

func main() {
	position := rover.NewPosition(0, 0)
	newRover, err := rover.NewRover(position, rover.North{})
	if err != nil {
		panic(err)
	}
	instructions := "MMMRMMLMM"
	if err = newRover.Act(instructions); err != nil {
		panic(err)
	}
	fmt.Println(newRover)
}
