package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ccwc [OPTIONS] FILENAME")
		_, err := fmt.Fprintf(os.Stdout, "Got %d args\n", len(os.Args))
		assertNoError(err)
		fmt.Println(os.Args)
		os.Exit(1)
	}

	flag := os.Args[1]
	fileName := os.Args[2]
	dat, err := os.ReadFile(fileName)
	assertNoError(err)

	switch flag {
	case "-c":
		fileSize := CountBytes(dat)
		fmt.Printf("%d %s\n", fileSize, fileName)
	case "-l":
		lineCount := CountLines(dat)
		fmt.Printf("%d %s\n", lineCount, fileName)
	case "-w":
		wordCount := CountWords(dat)
		fmt.Printf("%d %s\n", wordCount, fileName)
	default:
		fmt.Println("Usage: ccwc [OPTIONS] FILENAME")
	}
}

func assertNoError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
