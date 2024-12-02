package adventofcode

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Comparer struct {
	lists  [][]int
	counts map[int]int
}

func (c *Comparer) Load(lines []string) error {
	c.lists = make([][]int, 2)
	for i := range c.lists {
		c.lists[i] = make([]int, 0)
	}
	for _, line := range lines {
		entrys := strings.Fields(line)
		if len(entrys) != 2 {
			_ = fmt.Sprint(line)
			return errors.New("unexpected line")
		}
		for i, e := range entrys {
			value, err := strconv.Atoi(e)
			if err != nil {
				return err
			}
			c.lists[i] = append(c.lists[i], value)
		}
	}
	for i := range c.lists {
		sort.Ints(c.lists[i])
	}

	c.counts = make(map[int]int)
	for _, value := range c.lists[1] {
		c.counts[value] += 1
	}
	return nil
}

func (c *Comparer) AnswerA() (string, error) {
	result := 0
	a := c.lists[0]
	b := c.lists[1]
	for i := range a {
		result += abs(a[i], b[i])
	}
	return strconv.Itoa(result), nil
}

func (c *Comparer) AnswerB() (string, error) {
	result := 0
	for _, a := range c.lists[0] {
		result += a * c.counts[a]
	}
	return strconv.Itoa(result), nil
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
