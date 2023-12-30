package five

import "sort"

type converter struct {
	from                 string
	to                   string
	subconverters        []*subConverter
	inverseSubconverters []*subConverter
}

type subConverter struct {
	destination int64
	source      int64
	length      int64
}

func (c *converter) AddSubConverter(destination, source, length int64) {
	sub := &subConverter{
		destination: destination,
		source:      source,
		length:      length,
	}
	c.subconverters = append(c.subconverters, sub)
	c.inverseSubconverters = append(c.inverseSubconverters, sub)
	sort.Slice(c.subconverters, func(i, j int) bool {
		return c.subconverters[i].source > c.subconverters[j].source
	})
	sort.Slice(c.inverseSubconverters, func(i, j int) bool {
		return c.inverseSubconverters[i].destination > c.inverseSubconverters[j].destination
	})
}

func (c *converter) Convert(x int64) int64 {
	// After profiling, this is where a lot of the time is spent, it makes sense to optimize this.
	length := len(c.subconverters)
	i := sort.Search(length-1, func(i int) bool {
		return c.subconverters[i].source <= x
	})
	subconverter := c.subconverters[i]
	if i > length || x < subconverter.source || x >= subconverter.source+subconverter.length {
		return x
	}
	value := subconverter.destination + x - subconverter.source
	return value
}

func (c *converter) Invert(x int64) int64 {
	length := len(c.inverseSubconverters)
	i := sort.Search(length-1, func(i int) bool {
		return c.inverseSubconverters[i].destination <= x
	})
	subconverter := c.inverseSubconverters[i]
	if i > length || x < subconverter.destination || x >= subconverter.destination+subconverter.length {
		return x
	}
	value := subconverter.source + x - subconverter.destination
	return value
}
