package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"runtime"
)

const (
	noFlag = 2
	flag   = 3
)

func main() {
	err := run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

func run(args []string) error {
	fileInfo, err := os.Stdin.Stat()
	if err != nil {
		return err
	}
	if (fileInfo.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		var buffer bytes.Buffer
		var lineBreak []byte
		setLinebreak(&lineBreak)
		for scanner.Scan() {
			buffer.Write(scanner.Bytes())
			buffer.Write([]byte("\r\n"))
		}
		data := buffer.Bytes()
		err := handleStdInput(args, data)
		if err != nil {
			return err
		}
	} else {
		switch len(args) {
		case flag:
			err := handleFlagInput(args)
			if err != nil {
				return err
			}
		case noFlag:
			err := handleNoFlagInput(args)
			if err != nil {
				return err
			}
		default:
			fmt.Println("Usage: ccwc [OPTIONS] FILENAME")
			_, err := fmt.Fprintf(os.Stdout, "Got %d args\n", len(args))
			if err != nil {
				return err
			}
			return errors.New("flag: invalid option supplied")
		}
	}
	return nil
}

func handleStdInput(args []string, dat []byte) error {
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
		return errors.New("flag: invalid option supplied")
	}
	return nil
}

func handleNoFlagInput(args []string) error {
	fileName := args[1]
	dat, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	fileSize := CountBytes(dat)
	lineCount := CountLines(dat)
	wordCount := CountWords(dat)
	_, err = fmt.Fprintf(os.Stdout, "%d %d %d %s\n", lineCount, wordCount, fileSize, fileName)
	if err != nil {
		return err
	}
	return nil
}

func handleFlagInput(args []string) error {
	flagType := args[1]
	fileName := args[2]
	dat, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

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
		return errors.New("flag: invalid option supplied")
	}
	return nil
}

func setLinebreak(lineBreak *[]byte) {
	var lineEnd string
	switch runtime.GOOS {
	case "windows":
		lineEnd = "\r\n"
	default:
		lineEnd = "\n"
	}
	*lineBreak = []byte(lineEnd)
}
