package four

import (
	"strconv"
	"strings"
)

type Orchestrator struct {
	cards []*card
}

func (o *Orchestrator) Load(lines []string) error {
	o.cards = make([]*card, 0)
	for _, line := range lines {
		o.cards = append(o.cards, NewCard(parseLine(line)))
	}
	return nil
}

func parseLine(line string) (map[string]struct{}, map[string]struct{}) {
	_, contents, found := strings.Cut(line, ":")
	if !found {
		return nil, nil
	}
	winningNumbers, otherNumbers, found := strings.Cut(contents, "|")
	if !found {
		return nil, nil
	}
	winners := numberMap(winningNumbers)
	numbers := numberMap(otherNumbers)
	return winners, numbers
}

func numberMap(numbers string) map[string]struct{} {
	mapping := make(map[string]struct{}, 0)
	for _, num := range strings.Split(numbers, " ") {
		if num != "" {
			mapping[num] = struct{}{}
		}
	}
	return mapping
}

func (o *Orchestrator) Answer() (string, error) {
	var result int64 = 0
	for _, card := range o.cards {
		result += card.Points()
	}
	return strconv.FormatInt(result, 10), nil
}
