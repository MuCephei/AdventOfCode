package two

type game struct {
	id   int
	draws []*draw
}

func NewGame(id int, draws []*draw) *game {
	return &game{
		id:   id,
		draws: draws,
	}
}

func (g *game) Possible(reds, greens, blues int) bool {
	for _, d := range g.draws {
		if !d.Possible(reds, greens, blues) {
			return false
		}
	}
	return true
}

func (g *game) Power() int {
	red, green, blue := 0, 0, 0
	for _, draw := range g.draws {
		if draw.reds > red {
			red = draw.reds
		}
		if draw.greens > green {
			green = draw.greens
		}
		if draw.blues > blue {
			blue = draw.blues
		}
	}
	return red * green * blue
}

type draw struct {
	reds   int
	greens int
	blues  int
}

func NewDraw(reds, greens, blues int) *draw {
	return &draw{
		reds:   reds,
		greens: greens,
		blues:  blues,
	}
}

func (d *draw) Possible(reds, greens, blues int) bool {
	return d.reds <= reds && d.greens <= greens && d.blues <= blues
}
