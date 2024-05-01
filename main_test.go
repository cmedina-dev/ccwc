package main

import (
	"os"
	"testing"
)

const (
	path = "./"
	file = "test.txt"
)

func TestRun(t *testing.T) {
	t.Run("returns error when Stdin is invalid", func(t *testing.T) {
		writer, _, _ := os.Pipe()
		err := writer.Close()
		assertNoError(t, err)
		oldStdin := os.Stdin
		defer func() { os.Stdin = oldStdin }()
		os.Stdin = writer
		args := []string{path, "-m", file}
		err = run(args)
		assertError(t, err)
	})
	t.Run("accepts piped non-interactive input", func(t *testing.T) {
		reader, writer, err := os.Pipe()
		assertNoError(t, err)
		_, err = writer.Write([]byte("Hello World!\n\n\n"))
		assertNoError(t, err)
		err = writer.Close()
		assertNoError(t, err)
		oldStdin := os.Stdin
		defer func() {
			os.Stdin = oldStdin
			err = reader.Close()
			assertNoError(t, err)
		}()
		os.Stdin = reader
		args := []string{path, "-c"}
		err = run(args)
		assertNoError(t, err)
	})
	t.Run("accepts input with a flag", func(t *testing.T) {
		args := []string{path, "-c", file}
		err := run(args)
		assertNoError(t, err)
	})
	t.Run("accepts input without a flag", func(t *testing.T) {
		args := []string{path, file}
		err := run(args)
		assertNoError(t, err)
	})
	t.Run("rejects input with an unknown flag", func(t *testing.T) {
		args := []string{path, "-q", file}
		err := run(args)
		assertError(t, err)
	})
	t.Run("rejects input with an invalid number of arguments", func(t *testing.T) {
		args := []string{path, "-q", "-c", "-l", file}
		err := run(args)
		assertError(t, err)
	})
}

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
		_, writer, _ := os.Pipe()
		err := writer.Close()
		assertNoError(t, err)
		originalStdout := os.Stdout
		defer func() {
			os.Stdout = originalStdout
		}()
		os.Stdout = writer
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
	t.Helper()
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Error(err)
	}
}
