package three

type schematic struct {
	partNumbers []*partNumber
	symbols map[coordinate]rune
	partMap map[coordinate]*partNumber
}

func (s *schematic) AddLine(partNumbers []*partNumber, symbols map[coordinate]rune) {
	s.partNumbers = append(s.partNumbers, partNumbers...)
	for k, v := range symbols {
		s.symbols[k] = v
	}
	for _, partNumber := range partNumbers {
		for i := 0; i < partNumber.length; i++ {
			s.partMap[coordinate{x: partNumber.start.x + i, y: partNumber.start.y}] = partNumber
		}
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

func (s *schematic) GearTotal() int {
	result := 0
	for coord, symbol := range s.symbols {
		if symbol != '*' {
			continue
		}
		coordinates := surroundingCoordinates(coord, 1)
		nearbyPartNumbers := make(map[*partNumber]struct{})
		for _, point := range coordinates {
			if partNumber, ok := s.partMap[point]; ok {
				nearbyPartNumbers[partNumber] = struct{}{}
			}
		}
		if len(nearbyPartNumbers) != 2 {
			continue
		}
		product := 1
		for part := range nearbyPartNumbers {
			product *= part.value
		}
		result += product
	}
	return result
}

func (s *schematic) isPartNumber(partNumber *partNumber) bool {
	for _, coordinate := range surroundingCoordinates(partNumber.start, partNumber.length) {
		if _, ok := s.symbols[coordinate]; ok {
			return true
		}
	}
	return false
}

// surroundingCoordinates returns the surrounding coordinates of a point and a length in the x direction.
func surroundingCoordinates(start coordinate, length int) []coordinate {
	behind := coordinate{x: start.x-1, y: start.y}
	ahead := coordinate{x: start.x + length, y: start.y}
	around := []coordinate{behind, ahead}
	// start one beind and go to one ahead
	for i := -1; i <= length; i++ {
		x := start.x + i
		around = append(around, coordinate{x: x, y: start.y + 1})
		around = append(around, coordinate{x: x, y: start.y - 1})
	}
	return around
}

type partNumber struct {
	value int
	start coordinate
	length int
}

type coordinate struct {
	x int
	y int
}

