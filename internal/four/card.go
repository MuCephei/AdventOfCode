package four

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
	return int64(winners)
}