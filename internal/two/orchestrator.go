package two

import (
	"regexp"
	"strconv"
	"strings"
)

type Orchestrator struct {
	games []game
}

func (o *Orchestrator) Load(lines []string) error {
	for _, line := range lines {
		if err := o.load(line); err != nil {
			return err
		}
	}
	return nil
}

func (o *Orchestrator) load(line string) error {
	idRe := regexp.MustCompile("Game (?P<id>[0-9]+): ")
	matches := idRe.FindStringSubmatch(line)
	idMatch := matches[idRe.SubexpIndex("id")]
	id, err := strconv.Atoi(idMatch)
	if err != nil {
		return err
	}

	redRe := regexp.MustCompile("(?P<red>[0-9]+) red")
	greenRe := regexp.MustCompile("(?P<green>[0-9]+) green")
	blueRe := regexp.MustCompile("(?P<blue>[0-9]+) blue")
	draws := make([]*draw, 0)
	for _, draw := range strings.Split(line, ";") {
		red, err := getColour(redRe, draw)
		if err != nil {
			return err
		}
		green, err := getColour(greenRe, draw)
		if err != nil {
			return err
		}
		blue, err := getColour(blueRe, draw)
		if err != nil {
			return err
		}
		draws = append(draws, NewDraw(red, green, blue))
	}
	o.games = append(o.games, *NewGame(id, draws))
	return nil
}

func getColour(re *regexp.Regexp, draw string) (int, error) {
	matches := re.FindStringSubmatch(draw)
	if matches == nil {
		return 0, nil
	}
	count, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (o *Orchestrator) Answer() (string, error) {
	result := 0
	for _, game := range o.games {
		result += game.Power()
	}
	return strconv.Itoa(result), nil
}
