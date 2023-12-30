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
	'A': 0, 'K': 1, 'Q': 2, 'J': 3, 'T': 4, '9': 5,
	'8': 6, '7': 7, '6': 8, '5': 9, '4': 10, '3': 11, '2': 12,
}

func getRanking(values []int) ranking {
	pair := false
	triple := false
	for _, value := range values {
		switch value {
		case 5:
			return fiveOfAKind
		case 4:
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
				return twoPair
			}
			pair = true
		}
	}
	if triple {
		return threeOfAKind
	}
	if pair {
		return onePair
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
	for _, card := range []rune(cards) {
		values[cardMappings[card]] += 1
	}

	h := &hand{
		cards: cards,
		rank:  getRanking(values),
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
