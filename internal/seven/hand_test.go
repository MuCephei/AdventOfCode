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
		{hand: "T55A5", expectedRanking: threeOfAKind},
		{hand: "T55J5", expectedRanking: fourOfAKind},
		{hand: "KK677", expectedRanking: twoPair},
		{hand: "KT33T", expectedRanking: twoPair},
		{hand: "KTJJT", expectedRanking: fourOfAKind},
		{hand: "QQQ5A", expectedRanking: threeOfAKind},
		{hand: "QQQJA", expectedRanking: fourOfAKind},
		{hand: "A2T3K", expectedRanking: highCard},
		{hand: "A2TJK", expectedRanking: onePair},
		{hand: "QAQQA", expectedRanking: fullHouse},
		{hand: "QJJQ2", expectedRanking: fourOfAKind},
		{hand: "QJJQJ", expectedRanking: fiveOfAKind},
		{hand: "JJJJJ", expectedRanking: fiveOfAKind},
	}
	for _, test := range tests {
		hand := NewHand(test.hand, 1)
		if hand.rank != test.expectedRanking {
			t.Fatalf(`Incorrect rank for %s: expected %d got %d`, test.hand, test.expectedRanking, hand.rank)
		}
	}
}
