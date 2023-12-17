package one

import (
	"strconv"
)

type Orchestrator struct {
	input []string
}

func (o *Orchestrator) Load(lines []string) error {
	o.input = lines
	return nil
}

func (o *Orchestrator) Answer() (string, error){
	result := 0
	for _, line := range o.input {
		num, err := getCalibrationNumber(line)
		if err != nil {
			return "", err
		}
		result += num
	}
	return strconv.Itoa(result), nil
}