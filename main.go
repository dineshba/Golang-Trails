package main

import (
	"fmt"
	"mars_rover/rover"
	"os"
)

func main() {
	position := rover.NewPosition(0, 0)
	newRover, err := rover.NewRover(position, rover.Direction("N"))
	if err != nil {
		panic(err)
	}
	instructions := "MMMRMMLMM"
	for _, instruction := range instructions {
		err = newRover.Navigate(instruction)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	fmt.Println(newRover)
}
