package main

import (
	"bytes"
	"fmt"
	"testing"
)

type File struct {
	buffer *bytes.Buffer
	size   int
}

func TestCountBytes(t *testing.T) {
	t.Run("reads the size in bytes of a buffered file", func(t *testing.T) {
		mockFile := createMockFile(
			File{
				new(bytes.Buffer),
				1024,
			})
		got := CountBytes(mockFile)
		want := 1024
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("reads a file of size zero bytes", func(t *testing.T) {
		mockFile := createMockFile(
			File{
				new(bytes.Buffer),
				0,
			})
		got := CountBytes(mockFile)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestCountLines(t *testing.T) {
	t.Run("reads the number of lines in a given file", func(t *testing.T) {
		mockFile := createMockFile(File{
			new(bytes.Buffer),
			1024,
		})
		addMockData(mockFile, 20)
		fmt.Printf("%q\n", mockFile)
	})
}

func addMockData(file []byte, lineCount int) {
	for i := 0; i < lineCount; i++ {
		file[i] = '\n'
	}
}

func createMockFile(file File) []byte {
	mockData := make([]byte, file.size)
	file.buffer.Write(mockData)
	return file.buffer.Bytes()
}
