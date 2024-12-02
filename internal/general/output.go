package adventofcode

import (
	"os"
	"path"
)

// createDirectories makes any directories that don't exist under a filepath.
func createDirectories(filepath string) error {
	return os.MkdirAll(path.Dir(filepath), os.ModePerm)
}

// Save creates directories and then saves the specified contents to a filepath.
func Save(filepath string, contents string) error {
	err := createDirectories(filepath)
	if err != nil {
		return err
	}

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte(contents))
	if err != nil {
		return err
	}
	return nil
}
