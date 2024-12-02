package five

import (
	"testing"
)

func TestConvert(t *testing.T) {
	converter := &converter{}
	converter.AddSubConverter(50, 98, 2)
	converter.AddSubConverter(52, 50, 48)

	conversions := []struct {
		in  int64
		out int64
	}{
		{in: 0, out: 0},
		{in: 49, out: 49},
		{in: 50, out: 52},
		{in: 51, out: 53},
		{in: 97, out: 99},
		{in: 98, out: 50},
		{in: 99, out: 51},
		{in: 100, out: 100},
	}
	for _, c := range conversions {
		if out := converter.Convert(c.in); out != c.out {
			t.Fatalf(`Incorrect conversion for %d: expected %d got %d`, c.in, c.out, out)
		}
		if in := converter.Invert(c.out); in != c.in {
			t.Fatalf(`Incorrect inversion for %d: expected %d got %d`, c.out, c.in, in)
		}
	}
}
