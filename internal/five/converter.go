package five

type converter struct {
	from          string
	to            string
	subconverters []*subConverter
}

type subConverter struct {
	destination int64
	source      int64
	length      int64
}

func (c *converter) AddSubConverter(destination, source, length int64) {
	c.subconverters = append(c.subconverters, &subConverter{
		destination: destination,
		source:      source,
		length:      length,
	})
}

func (c *converter) Convert(x int64) int64 {
	// A binary search on a sorted list of subconverters would probably be faster.
	// I don't think it's worth it here.
	for _, subconv := range c.subconverters {
		delta := x - subconv.source
		if subconv.source <= x && delta <= subconv.length {
			return subconv.destination + delta
		}
	}
	return x
}
