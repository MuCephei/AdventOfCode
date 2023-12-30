package six

import (
	"errors"
	"strconv"
	"strings"
)

type Orchestrator struct {
	races []*race
}

func (o *Orchestrator) Load(lines []string) error {
	times := make([]int, 0)
	records := make([]int, 0)
	_, timeline, found := strings.Cut(lines[0], ":")
	if !found {
		return errors.New("Could not find times")
	}
	for _, t := range strings.Fields(timeline) {
		time, err := strconv.Atoi(t)
		if err != nil {
			return err
		}
		times = append(times, time)
	}

	_, recordline, found := strings.Cut(lines[1], ":")
	if !found {
		return errors.New("Could not find times")
	}
	for _, r := range strings.Fields(recordline) {
		record, err := strconv.Atoi(r)
		if err != nil {
			return err
		}
		records = append(records, record)
	}
	o.races = make([]*race, 0)

	for i, time := range times {
		o.races = append(o.races, NewRace(time, records[i]))
	}
	return nil
}

func (o *Orchestrator) Answer() (string, error) {
	result := 1
	for _, race := range o.races {
		result *= race.Winners()
	}
	return strconv.Itoa(result), nil
}
