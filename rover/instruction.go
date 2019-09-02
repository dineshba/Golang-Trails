package rover

import "fmt"

type Instruction interface {
	Action(r *rover)
}

type Move struct{}

func (Move) Action(r *rover) {
	r.move()
}

type Left struct{}

func (Left) Action(r *rover) {
	r.left()
}

type Right struct{}

func (Right) Action(r *rover) {
	r.right()
}

func ParseInstructions(instructions string) ([]Instruction, error) {
	var acc []Instruction
	for _, instruction := range instructions {
		switch instruction {
		case 'M':
			acc = append(acc, Move{})
		case 'L':
			acc = append(acc, Left{})
		case 'R':
			acc = append(acc, Right{})
		default:
			return nil, fmt.Errorf("not a valid instruction %c", instruction)
		}
	}
	return acc, nil
}
