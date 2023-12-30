package seven

import (
	"slices"
	"strconv"
	"strings"
)

type Orchestrator struct {
	hands []*hand
}

func (o *Orchestrator) Load(lines []string) error {
	for _, line := range lines {
		cards, num, found := strings.Cut(line, " ")
		if !found {
			return nil
		}
		bid, err := strconv.Atoi(num)
		if err != nil {
			return err
		}
		o.hands = append(o.hands, NewHand(cards, bid))
	}
	return nil
}

func (o *Orchestrator) Answer() (string, error) {
	slices.SortFunc[[]*hand](o.hands, func (a, b *hand) int {
		return a.compare(b)
	})

	result := 0
	for i, hand := range o.hands {
		result += hand.bid * (i + 1)
	}
	return strconv.Itoa(result), nil
}
