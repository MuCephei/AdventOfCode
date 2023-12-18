package five

import (
	"errors"
	"strconv"
	"strings"
)

const delimiter string = "|"

type Orchestrator struct {
	seeds      []int64
	converters map[string]*converter
}

func (o *Orchestrator) Load(lines []string) error {
	seeds, err := parseSeeds(lines[0])
	if err != nil {
		return err
	}
	o.seeds = seeds
	converters, err := parseConverters(lines[2:])
	if err != nil {
		return err
	}
	o.converters = converters

	return nil
}

func parseSeeds(line string) ([]int64, error) {
	seeds := make([]int64, 0)
	_, seedList, found := strings.Cut(line, "seeds:")
	if !found {
		return nil, errors.New("Could not parse initial seeds")
	}
	for _, seed := range strings.Split(strings.TrimSpace(seedList), " ") {
		s, err := strconv.ParseInt(seed, 10, 0)
		if err != nil {
			return nil, err
		}
		seeds = append(seeds, s)
	}
	return seeds, nil
}

func parseConverters(lines []string) (map[string]*converter, error) {
	converters := make(map[string]*converter)
	wholeThing := strings.Join(lines, delimiter)
	for _, section := range strings.Split(wholeThing, delimiter+delimiter) {
		converter, err := parseConverter(section)
		if err != nil {
			return nil, err
		}
		converters[converter.from] = converter
	}

	return converters, nil
}

func parseConverter(section string) (*converter, error) {
	mapping, conversions, found := strings.Cut(strings.TrimSpace(section), " map:|")
	if !found {
		return nil, errors.New("Could not parse section: " + section)
	}
	from, to, found := strings.Cut(mapping, "-to-")
	if !found {
		return nil, errors.New("Could not parse mapping")
	}
	converter := &converter{
		from: from,
		to: to,
		subconverters: make([]*subConverter, 0),
	}
	for _, conversion := range strings.Split(conversions, delimiter) {
		conversionParts := strings.Split(conversion, " ")
		destination, err := strconv.ParseInt(conversionParts[0], 10, 0)
		if err != nil {
			return nil, err
		}
		source, err := strconv.ParseInt(conversionParts[1], 10, 0)
		if err != nil {
			return nil, err
		}
		length, err := strconv.ParseInt(conversionParts[2], 10, 0)
		if err != nil {
			return nil, err
		}
		converter.AddSubConverter(destination, source, length)
	}
	return converter, nil
}

func (o *Orchestrator) Convert(x int64, category string) (int64, string) {
	converter := o.converters[category]
	return converter.Convert(x), converter.to
}

func (o *Orchestrator) Answer() (string, error) {
	best := int64(0)
	for _, seed := range o.seeds {
		x, category := o.Convert(seed, "seed")
		for category != "location" {
			x, category = o.Convert(x, category)
		}
		if best == 0 || best > x {
			best = x
		}
	}
	return strconv.FormatInt(best, 10), nil
}
