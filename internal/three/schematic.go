package three

type schematic struct {
	partNumbers []*partNumber
	symbols map[coordinate]rune
}

func (s *schematic) AddLine(partNumbers []*partNumber, symbols map[coordinate]rune) {
	s.partNumbers = append(s.partNumbers, partNumbers...)
	for k, v := range symbols {
		s.symbols[k] = v
	}
}

func (s *schematic) Total() int {
	result := 0
	for _, partNumber := range s.partNumbers {
		if s.isPartNumber(partNumber) {
			result += partNumber.value
		}
	}
	return result
}

func (s *schematic) isPartNumber(partNumber *partNumber) bool {
	for _, coordinate := range partNumber.surroundingCoordinates() {
		if _, ok := s.symbols[coordinate]; ok {
			return true
		}
	}
	return false
}

type partNumber struct {
	value int
	start coordinate
	length int
}

// surroundingCoordinates returns the surrounding coordinates of a number.
func (pn *partNumber) surroundingCoordinates () []coordinate {
	behind := coordinate{x: pn.start.x-1, y: pn.start.y}
	ahead := coordinate{x: pn.start.x + pn.length, y: pn.start.y}
	around := []coordinate{behind, ahead}
	// start one beind and go to one ahead
	for i := -1; i <= pn.length; i++ {
		x := pn.start.x + i
		around = append(around, coordinate{x: x, y: pn.start.y + 1})
		around = append(around, coordinate{x: x, y: pn.start.y - 1})
	}
	return around
}

type coordinate struct {
	x int
	y int
}

