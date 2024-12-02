package one

import (
	"strconv"
)

var digits []string = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

func isSpelledDigit(line string, index int) int {
	for n, word := range digits {
		if len(line) >= index + len(word) && line[index:index+len(word)] == word {
			return n
		}
	}
	return -1
}

func getFirstDigit(line string) string {
	for i, char := range line {
		if isDigit(char) {
			return string(char)
		}
		digit := isSpelledDigit(line, i)
		if digit != -1 {
			return strconv.Itoa(digit)
		}
	}
	return ""
}

func getLastDigit(line string) string {
	for i := len(line) - 1; i >= 0; i-- {
		char := rune(line[i])
		if isDigit(char) {
			return string(char)
		}
		digit := isSpelledDigit(line, i)
		if digit != -1 {
			return strconv.Itoa(digit)
		}
	}
	return ""
}

func getCalibrationNumber(line string) (int, error) {
	first := getFirstDigit(line)
	last := getLastDigit(line)
	return strconv.Atoi(first + last)
}
