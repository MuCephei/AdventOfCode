package general

import (
	"bufio"
	"os"
)

type DataStore interface {
	Load(lines []string)
}

func Load(ds DataStore, filename string) (error) {
	lines, err := loadFile(filename)
	if err != nil {
		return err
	}
	ds.Load(lines)
	return nil
}

// load loads a file line by line into a string slice.
func loadFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
