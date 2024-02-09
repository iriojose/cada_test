package utils

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	testContent := []byte("Hello, world!")
	tempFile, err := ioutil.TempFile("", "testfile")

	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.Write(testContent); err != nil {
		tempFile.Close()
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	tempFile.Close()

	content, err := ReadFile(tempFile.Name())

	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	if string(content) != string(testContent) {
		t.Errorf("Incorrect file content. expected %s but got %s", string(testContent), string(content))
	}
}

func TestReadFileError(t *testing.T) {
	_, err := ReadFile("fake.txt")

	if err == nil {
		t.Error("No errors returned when reading a non-existent file")
	}
	if _, ok := err.(*os.PathError); !ok {
		t.Errorf("Type error expected os.PathError, but it was received: %v", err)
	}
}
