package two

type game struct {
	id   int
	draw []*draw
}

func NewGame(id int, draws []*draw) *game {
	return &game{
		id:   id,
		draw: draws,
	}
}

func (g *game) Possible(reds, greens, blues int) bool {
	for _, d := range g.draw {
		if !d.Possible(reds, greens, blues) {
			return false
		}
	}
	return true
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
