package direction

// Direction represents the rover direction
type Direction string

// IsValid validates the direction
func (direction Direction) IsValid() bool {
	return direction == "N" || direction == "E" || direction == "W" || direction == "S"
}
