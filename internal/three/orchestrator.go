package three

import (
	"strconv"
)

type Orchestrator struct {
	schematic schematic
}

func (o *Orchestrator) Load(lines []string) error {
	o.schematic.partNumbers = make([]*partNumber, 0)
	o.schematic.symbols = make(map[coordinate]rune)
	for n, line := range lines {
		partNumbers, symbols, err := parseLine(n, line)
		if err != nil {
			return nil
		}
		o.schematic.AddLine(partNumbers, symbols)
	}
	return nil
}

func parseLine(n int, line string) ([]*partNumber, map[coordinate]rune, error) {
	currentNumber := ""
	partNumbers := make([]*partNumber, 0)
	symbols := make(map[coordinate]rune)
	for i, char := range line {
		if isDigit(char) {
			currentNumber += string(char)
			continue
		} 
		if currentNumber != "" {
			partNumber, err := makeNumber(n, i, currentNumber)
			if err != nil {
				return nil, nil, err
			}
			partNumbers = append(partNumbers, partNumber)
			currentNumber = ""
		}
		if char != '.' {
			symbols[coordinate{x: i, y: n}] = char
		}
	}

	if currentNumber != "" {
		partNumber, err := makeNumber(n, len(line), currentNumber)
		if err != nil {
			return nil, nil, err
		}
		partNumbers = append(partNumbers, partNumber)
		currentNumber = ""
	}
	return partNumbers, symbols, nil
}

func makeNumber(n, i int, value string) (*partNumber, error) {
	numberValue, err := strconv.Atoi(value)
	if err != nil {
		return nil, err
	}
	return &partNumber{
		value: numberValue,
		start: coordinate{
			x: i - len(value),
			y: n,
		},
		length: len(value),
	}, nil
}

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

func (o *Orchestrator) Answer() (string, error) {
	return strconv.Itoa(o.schematic.Total()), nil
}
