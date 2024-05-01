package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

const (
	path = "./"
	file = "test.txt"
)

func TestHandleStdInput(t *testing.T) {
	testCases := []struct {
		name string
		path string
		flag string
		file string
	}{
		{
			"accepts the flag to count piped bytes",
			path,
			"-c",
			file,
		},
		{
			"accepts the flag to count piped words",
			path,
			"-w",
			file,
		},
		{
			"accepts the flag to count piped lines",
			path,
			"-l",
			file,
		},
		{
			"accepts the flag to count piped characters",
			path,
			"-m",
			file,
		},
	}
	var buffer []byte
	buffer = append(buffer, "Hello "...)
	buffer = append(buffer, "World!"...)

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			args := []string{path, testCase.flag}
			err := handleStdInput(args, buffer)
			assertNoError(t, err)
		})
	}
	t.Run("returns error when invalid flag is supplied", func(t *testing.T) {
		args := []string{path, "-q"}
		err := handleStdInput(args, buffer)
		assertError(t, err)
	})
}

func TestHandleNoFlagInput(t *testing.T) {
	t.Run("accepts input without flag", func(t *testing.T) {
		args := []string{path, file}
		err := handleNoFlagInput(args)
		assertNoError(t, err)
	})
	t.Run("returns error when file does not exist", func(t *testing.T) {
		args := []string{path, "nonexistent.txt"}
		err := handleNoFlagInput(args)
		assertError(t, err)
	})
	t.Run("returns error when writing to stdout fails", func(t *testing.T) {
		_, w, _ := os.Pipe()
		err := w.Close()
		if err != nil {
			t.Fatalf("Failed to close writer: %v", err)
		}

		originalStdout := os.Stdout
		defer func() {
			os.Stdout = originalStdout
		}()
		os.Stdout = w

		args := []string{path, file}
		err = handleNoFlagInput(args)
		assertError(t, err)
	})
}

func TestHandleFlagInput(t *testing.T) {
	testCases := []struct {
		name string
		path string
		flag string
		file string
	}{
		{
			"accepts the flag to count bytes",
			path,
			"-c",
			file,
		},
		{
			"accepts the flag to count words",
			path,
			"-w",
			file,
		},
		{
			"accepts the flag to count lines",
			path,
			"-l",
			file,
		},
		{
			"accepts the flag to count characters",
			path,
			"-m",
			file,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			args := []string{testCase.path, testCase.flag, testCase.file}
			err := handleFlagInput(args)
			assertNoError(t, err)
		})
	}
	t.Run("returns error when file does not exist", func(t *testing.T) {
		args := []string{path, "-l", "nonexistent.txt"}
		err := handleFlagInput(args)
		assertError(t, err)
	})
	t.Run("returns error when invalid flag supplied", func(t *testing.T) {
		args := []string{path, "-q", "test.txt"}
		err := handleFlagInput(args)
		assertError(t, err)
	})
}

func assertError(t *testing.T, err error) {
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
