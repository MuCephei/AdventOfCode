package adventofcode

import (
	"fmt"
	"strconv"
)

type Orchestrator struct {
	input []string
}

func (o *Orchestrator) Load(lines []string) error {
	o.input = lines
	return nil
}

func (o *Orchestrator) Answer() (string, error) {
	result := 0
	for _, line := range o.input {
		num := 3
		fmt.Println(line)
		// if err != nil {
		// 	return "", err
		// }
		result += num
	}
	return strconv.Itoa(result), nil
}
