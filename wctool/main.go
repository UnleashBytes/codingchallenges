package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	byteCountFlag := fs.Bool("c", false, "Count bytes")
	lineCountFlag := fs.Bool("l", false, "Count lines")
	wordCountFlag := fs.Bool("w", false, "Count words")
	charCountFlag := fs.Bool("m", false, "Count characters")

	fs.Parse(os.Args[1:])

	if fs.NArg() > 1 {
		fmt.Println("Usage: gowc -c <file> or gowc -l <file>")
		os.Exit(1)
	}
	var filename string
	var data []byte
	var err error

	if fs.NArg() == 1 {
		filename = fs.Arg(0)
		data, err = os.ReadFile(filename)
		if err != nil {
			fmt.Println("Error reading file:", err)
			os.Exit(1)
		}
	} else {
		data, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("Error reading from stdin:", err)
			os.Exit(1)
		}
	}

	var byteCount, lineCount, wordCount, charCount int
	var output strings.Builder
	if err != nil {
		fmt.Println("Error reading from stdin:", err)
		os.Exit(1)
	}

	if *byteCountFlag {
		byteCount = countBytes(data)
		output.WriteString(fmt.Sprintf("%d ", byteCount))
	} else if *lineCountFlag {
		lineCount = countLines(data)
		output.WriteString(fmt.Sprintf("%d ", lineCount))
	} else if *wordCountFlag {
		wordCount = countWords(data)
		output.WriteString(fmt.Sprintf("%d ", wordCount))
	} else if *charCountFlag {
		charCount = countCharacters(data)
		output.WriteString(fmt.Sprintf("%d ", charCount))
	} else {
		byteCount = countBytes(data)
		lineCount = countLines(data)
		wordCount = countWords(data)
		output.WriteString(fmt.Sprintf("%d %d %d ", lineCount, wordCount, byteCount))
	}

	fmt.Printf("%s%s\n", output.String(), filename)
}

func countBytes(data []byte) int {
	return len(data)
}

func countLines(data []byte) int {
	scanner := bufio.NewScanner(bytes.NewReader(data))

	lines := 0
	for scanner.Scan() {
		lines++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading data:", err)
		os.Exit(1)
	}

	return lines
}

func countWords(data []byte) int {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	scanner.Split(bufio.ScanWords)

	words := 0
	for scanner.Scan() {
		words++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading data:", err)
		os.Exit(1)
	}

	return words
}

func countCharacters(data []byte) int {
	text := string(data)
	return len([]rune(text))
}
