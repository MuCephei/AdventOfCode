package four

import (
	"math"
)

type card struct {
	winningNumbers map[string]struct{}
	numbers map[string]struct{}
}

func NewCard(winningNumbers, numbers map[string]struct{}) *card {
	return &card{
		winningNumbers: winningNumbers,
		numbers: numbers,
	}
}

func (c *card) Points() int64 {
	winners := 0
	for winner := range c.winningNumbers {
		if _, ok := c.numbers[winner]; ok {
			winners += 1
		}
	}
	if winners == 0 {
		return 0
	}
	return int64(math.Pow(float64(2), float64(winners - 1)))
}