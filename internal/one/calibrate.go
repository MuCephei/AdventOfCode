package one

import (
	"strconv"
)

func isNumber(char rune) bool {
	return char >= '0' && char <= '9'
}

func getFirstDigit(line string) string {
	for _, char := range line {
		if isNumber(char) {
			return string(char)
		}
	}
	return ""
}

func getLastDigit(line string) string {
	for i := len(line) - 1; i >= 0; i-- {
		char := rune(line[i])
		if isNumber(char) {
			return string(char)
		}
	}
	return ""
}

func getCalibrationNumber(line string) (int, error) {
	first := getFirstDigit(line)
	last := getLastDigit(line)
	return strconv.Atoi(first + last)
}
