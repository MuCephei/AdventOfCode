package four

import (
	"strconv"
	"strings"
)

type Orchestrator struct {
	cards []*card
	copiesOfCards []int64
}

func (o *Orchestrator) Load(lines []string) error {
	o.cards = make([]*card, 0)
	for _, line := range lines {
		o.cards = append(o.cards, NewCard(parseLine(line)))
		o.copiesOfCards = append(o.copiesOfCards, 1)
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
	for i, card := range o.cards {
		points := card.Points()
		copies := o.copiesOfCards[i]
		for j := int64(1); j <= points; j++ {
			o.copiesOfCards[int64(i) + j] += copies
		}
	}
	totalCards :=int64(0)
	for _, copies := range o.copiesOfCards {
		totalCards += copies
	}
	return strconv.FormatInt(totalCards, 10), nil
}
