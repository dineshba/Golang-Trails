package rover

import "fmt"

type Instruction int

const (
	Move Instruction = iota
	Left
	Right
)

func ParseInstructions(instructions string) ([]Instruction, error) {
	var acc []Instruction
	for _, instruction := range instructions {
		switch instruction {
		case 'M':
			acc = append(acc, Move)
		case 'L':
			acc = append(acc, Left)
		case 'R':
			acc = append(acc, Right)
		default:
			return nil, fmt.Errorf("not a valid instruction %c", instruction)
		}
	}
	return acc, nil
}
