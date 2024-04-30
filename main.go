package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	noFlag = 2
	flag   = 3
)

func main() {
	content, err := io.ReadAll(os.Stdin)
	if err == nil {
		handleStdInput(os.Args, content)
	} else {
		switch len(os.Args) {
		case flag:
			handleFlagInput(os.Args)
		case noFlag:
			handleNoFlagInput(os.Args)
		default:
			fmt.Println("Usage: ccwc [OPTIONS] FILENAME")
			_, err := fmt.Fprintf(os.Stdout, "Got %d args\n", len(os.Args))
			assertNoError(err)
			fmt.Println(os.Args)
			os.Exit(1)
		}
	}
}

func handleStdInput(args []string, dat []byte) {
	flagType := args[1]
	switch flagType {
	case "-c":
		fileSize := CountBytes(dat)
		fmt.Printf("%d\n", fileSize)
	case "-l":
		lineCount := CountLines(dat)
		fmt.Printf("%d\n", lineCount)
	case "-w":
		wordCount := CountWords(dat)
		fmt.Printf("%d\n", wordCount)
	case "-m":
		characterCount := CountCharacters(dat)
		fmt.Printf("%d\n", characterCount)
	default:
		fmt.Println("Usage: ccwc [OPTIONS] FILENAME")
	}
}

func handleNoFlagInput(args []string) {
	fileName := args[1]
	dat, err := os.ReadFile(fileName)
	assertNoError(err)
	fileSize := CountBytes(dat)
	lineCount := CountLines(dat)
	wordCount := CountWords(dat)
	_, err = fmt.Fprintf(os.Stdout, "%d %d %d %s\n", lineCount, wordCount, fileSize, fileName)
	assertNoError(err)
}

func handleFlagInput(args []string) {
	flagType := args[1]
	fileName := args[2]
	dat, err := os.ReadFile(fileName)
	assertNoError(err)

	switch flagType {
	case "-c":
		fileSize := CountBytes(dat)
		fmt.Printf("%d %s\n", fileSize, fileName)
	case "-l":
		lineCount := CountLines(dat)
		fmt.Printf("%d %s\n", lineCount, fileName)
	case "-w":
		wordCount := CountWords(dat)
		fmt.Printf("%d %s\n", wordCount, fileName)
	case "-m":
		characterCount := CountCharacters(dat)
		fmt.Printf("%d %s\n", characterCount, fileName)
	default:
		fmt.Println("Usage: ccwc [OPTIONS] FILENAME")
	}
}

func assertNoError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
