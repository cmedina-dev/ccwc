package main

import (
	"bytes"
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
		assertCount(t, got, want)
	})
}

func TestCountLines(t *testing.T) {
	t.Run("counts the number of \\n characters", func(t *testing.T) {
		mockFile := createMockFile(File{
			new(bytes.Buffer),
			8,
		})
		addChar(&mockFile, 4, '\n')
		got := CountLines(mockFile)
		want := 4
		assertCount(t, got, want)
	})
	t.Run("counts the number of \\r characters", func(t *testing.T) {
		mockFile := createMockFile(File{
			new(bytes.Buffer),
			8,
		})
		addChar(&mockFile, 2, '\r')
		got := CountLines(mockFile)
		want := 2
		assertCount(t, got, want)
	})
	t.Run("counts the number of \\r\\n characters", func(t *testing.T) {
		mockFile := createMockFile(File{
			new(bytes.Buffer),
			8,
		})
		addChar(&mockFile, 1, '\r')
		addChar(&mockFile, 1, '\n')
		addChar(&mockFile, 1, '\r')
		got := CountLines(mockFile)
		want := 2
		assertCount(t, got, want)
	})
}

func TestCountWords(t *testing.T) {
	t.Run("counts the number of words in a file", func(t *testing.T) {
		mockFile := createMockFile(File{
			new(bytes.Buffer),
			16,
		})
		addWord(&mockFile, "Hello ")
		addWord(&mockFile, "World!")
		got := CountWords(mockFile)
		want := 2
		assertCount(t, got, want)
	})
}

func TestCountCharacters(t *testing.T) {
	t.Run("counts the number of characters in a file", func(t *testing.T) {
		mockFile := createMockFile(File{
			new(bytes.Buffer),
			16,
		})
		addWord(&mockFile, "Hello ")
		addWord(&mockFile, "World!")
		got := CountCharacters(mockFile)
		want := 28
		assertCount(t, got, want)
	})
}

func addChar(file *[]byte, lineCount int, lineBreak byte) {
	for i := 0; i < lineCount; i++ {
		*file = append(*file, lineBreak)
	}
}

func addWord(file *[]byte, word string) {
	for i := 0; i < len(word); i++ {
		*file = append(*file, word[i])
	}
}

func createMockFile(file File) []byte {
	mockData := make([]byte, file.size)
	file.buffer.Write(mockData)
	return file.buffer.Bytes()
}

func assertCount(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
