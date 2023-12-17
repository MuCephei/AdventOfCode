package one

import (
	"strconv"
)

type orchestrator struct {
	input []string
}

func NewOrchestrator(lines []string) *orchestrator {
	o := orchestrator{input: lines}
	return &o
}

func (o *orchestrator) Answer() (string, error){
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