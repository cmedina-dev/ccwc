package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	countTypes := flag.NewFlagSet("countTypes", flag.ExitOnError)
	countTypes.Bool("c", false, "prints the file size in bytes")
	countTypes.Bool("l", false, "prints the number of lines in the file")

	if len(os.Args) != 3 {
		fmt.Println("Usage: ccwc [OPTIONS] FILENAME")
		fmt.Fprintf(os.Stdout, "Got %d args\n", len(os.Args))
		fmt.Println(os.Args)
		os.Exit(1)
	}

	err := countTypes.Parse(os.Args[2:])
	assertNoError(err)

	fileName := countTypes.Arg(0)
	dat, err := os.ReadFile(fileName)
	assertNoError(err)

	switch os.Args[1] {
	case "-c":
		fileSize := CountBytes(dat)
		fmt.Printf("%d %s\n", fileSize, fileName)
	case "-l":
		lineCount := CountLines(dat)
		fmt.Printf("%d %s\n", lineCount, fileName)
	}
}

func assertNoError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
