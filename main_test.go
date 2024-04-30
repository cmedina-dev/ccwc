package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestHandleFlagInput(t *testing.T) {
	t.Run("accepts a file as input with a flag", func(t *testing.T) {

	})
}
