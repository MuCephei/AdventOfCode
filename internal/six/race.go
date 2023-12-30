package six

type race struct {
	time int
	record int
}

func NewRace(time, record int) *race {
	return &race{
		time: time,
		record: record,
	}
}

func (r *race) Winners() int {
	// Can you do this with a single equation? Yes.
	// Am I going to bother figuring it out yet? No.
	total := 0
	for speed := 0; speed < r.time; speed++ {
		if (r.time - speed) * speed > r.record {
			total += 1
		}
	}
	return total
}