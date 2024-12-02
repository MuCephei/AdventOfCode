package seven

type ranking int

const (
	highCard ranking = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

var cardMappings map[rune]int = map[rune]int{
	'A': 0, 'K': 1, 'Q': 2, 'T': 3, '9': 4,
	'8': 5, '7': 6, '6': 7, '5': 8, '4': 9, '3': 10, '2': 11, 'J': 12, 
}

func getRanking(values []int, jokers int) ranking {
	// Could this be neater? Yes.
	pair := false
	triple := false
	for _, value := range values {
		switch value {
		case 5:
			return fiveOfAKind
		case 4:
			if jokers == 1 {
				return fiveOfAKind
			}
			return fourOfAKind
		case 3:
			if pair {
				return fullHouse
			}
			triple = true
		case 2:
			if triple {
				return fullHouse
			}
			if pair {
				if jokers >= 1 {
					return fullHouse
				}
				return twoPair
			}
			pair = true
		}
	}
	if triple {
		if jokers == 1 {
			return fourOfAKind
		}
		if jokers == 2 {
			return fiveOfAKind
		}
		return threeOfAKind
	}
	if pair {
		if jokers == 1 {
			return threeOfAKind
		}
		if jokers == 2 {
			return fourOfAKind
		}
		if jokers == 3 {
			return fiveOfAKind
		}
		return onePair
	}
	if jokers == 1 {
		return onePair
	}
	if jokers == 2 {
		return threeOfAKind
	}
	if jokers == 3 {
		return fourOfAKind
	}
	if jokers == 4 || jokers == 5 {
		return fiveOfAKind
	}
	return highCard
}

type hand struct {
	cards string
	rank  ranking
	bid   int
}

func NewHand(cards string, bid int) *hand {
	values := make([]int, len(cardMappings))
	jokers := 0
	for _, card := range []rune(cards) {
		if card == 'J' {
			jokers += 1
		} else {
			values[cardMappings[card]] += 1
		}
	}

	h := &hand{
		cards: cards,
		rank:  getRanking(values, jokers),
		bid:   bid,
	}

	return h
}

func (h *hand) compare(o *hand) int {
	if h.rank > o.rank {
		return 1
	} else if h.rank < o.rank {
		return -1
	}
	for i, card := range h.cards {
		hValue := cardMappings[card]
		oValue := cardMappings[rune(o.cards[i])]
		if hValue < oValue {
			return 1
		} else if hValue > oValue {
			return -1
		}
	}
	return 0
}
