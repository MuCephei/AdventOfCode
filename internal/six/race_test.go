package six

import (
	"testing"
)

func TestRace(t *testing.T) {
	tests := []struct {
		time            int
		record          int
		expectedWinners int
	}{
		{time: 7, record: 9, expectedWinners: 4},
		{time: 15, record: 40, expectedWinners: 8},
		{time: 30, record: 200, expectedWinners: 9},
	}
	for _, test := range tests {
		race := NewRace(test.time, test.record)
		if winners := race.Winners(); winners != test.expectedWinners {
			t.Fatalf(`Incorrect winners for %d-%d: expected %d got %d`, test.time, test.record, test.expectedWinners, winners)
		}
	}
}
