package five

import (
	"errors"
	"strconv"
	"strings"
)

const delimiter string = "|"

type Orchestrator struct {
	seeds      []*seed
	converters map[string]*converter
	inverters  map[string]*converter
}

type seed struct {
	start  int64
	length int64
}

func (s *seed) in(x int64) bool {
	return x >= s.start && x < s.start+s.length
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
	o.inverters = make(map[string]*converter)
	for _, converter := range o.converters {
		o.inverters[converter.to] = converter
	}

	return nil
}

func parseSeeds(line string) ([]*seed, error) {
	seeds := make([]*seed, 0)
	_, seedList, found := strings.Cut(line, "seeds:")
	if !found {
		return nil, errors.New("Could not parse initial seeds")
	}
	seedInfo := strings.Split(strings.TrimSpace(seedList), " ")
	for i := 0; i < len(seedInfo); i += 2 {
		start, err := strconv.ParseInt(seedInfo[i], 10, 0)
		if err != nil {
			return nil, err
		}
		length, err := strconv.ParseInt(seedInfo[i+1], 10, 0)
		if err != nil {
			return nil, err
		}
		seeds = append(seeds, &seed{start: start, length: length})
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
		to:   to,
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

func (o *Orchestrator) Invert(x int64, category string) (int64, string) {
	inverter := o.inverters[category]
	return inverter.Invert(x), inverter.from
}

func (o *Orchestrator) Answer() (string, error) {
	x := int64(0)
	for true {
		y, category := o.Invert(x, "location")
		for category != "seed" {
			y, category = o.Invert(y, category)
		}
		for _, seed := range o.seeds {
			if seed.in(y) {
				return strconv.FormatInt(x, 10), nil
			}
		}
		x++
	}
	return "", errors.New("Reached end of infinte loop")
}
