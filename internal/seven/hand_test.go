package seven

import (
	"testing"
)

func TestRanking(t *testing.T) {
	tests := []struct {
		hand            string
		expectedRanking ranking
	}{
		{hand: "AAAQA", expectedRanking: fourOfAKind},
		{hand: "32T3K", expectedRanking: onePair},
		{hand: "T55J5", expectedRanking: threeOfAKind},
		{hand: "KK677", expectedRanking: twoPair},
		{hand: "KTJJT", expectedRanking: twoPair},
		{hand: "QQQJA", expectedRanking: threeOfAKind},
		{hand: "A2T3K", expectedRanking: highCard},
		{hand: "QAQQA", expectedRanking: fullHouse},
	}
	for _, test := range tests {
		hand := NewHand(test.hand, 1)
		if hand.rank != test.expectedRanking {
			t.Fatalf(`Incorrect rank for %s: expected %d got %d`, test.hand, test.expectedRanking, hand.rank)
		}
	}
}
