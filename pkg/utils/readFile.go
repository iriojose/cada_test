package utils

import (
	"io"
	"os"
)

// reads and returns bytes of the read file
func ReadFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	byteValue, _ := io.ReadAll(file)

	return byteValue, nil
}
