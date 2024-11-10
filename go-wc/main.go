package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	printByteCount := flag.Bool("c", false, "Prints the byte count of a file")
	printLineCount := flag.Bool("l", false, "Prints the line count of a file")
	printWordCount := flag.Bool("w", false, "Prints the word count of a file")
	printCharCount := flag.Bool("m", false, "Prints the character count of a file")
	flag.Parse()

	textData, filePath := getInputTextData()

	if *printByteCount {
		fmt.Printf("%v %v\n", len(textData), filePath)
	}

	if *printLineCount {
		fmt.Printf("%v %v", getLineCount(textData), filePath)
	}

	if *printWordCount {
		fmt.Printf("%v %v", getWordCount(textData), filePath)
	}

	if *printCharCount {
		fmt.Printf("%v %v", getCharacterCount(textData), filePath)
	}

	// If no flag is passed then print lines, words, and byte counts as default
	if !*printByteCount &&
		!*printLineCount &&
		!*printWordCount &&
		!*printCharCount {
		fmt.Printf("%v %v %v %v", getLineCount(textData), getWordCount(textData), len(textData), filePath)
	}
}

func getInputTextData() ([]byte, string) {
	filePath, err := getFilePathArg()
	var textData []byte

	if err == nil {
		textData, err = os.ReadFile("./" + filePath)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		textData, err = io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
	}

	return textData, filePath
}

func getFilePathArg() (string, error) {
	args := os.Args[1:]

	if len(args) == 0 {
		return "", fmt.Errorf("No arguments were passed")
	}

	filePath := args[len(args)-1]

	if strings.HasPrefix(filePath, "-") {
		return "", fmt.Errorf("Last argument must be a file path")
	}

	return filePath, nil
}

func getLineCount(textData []byte) int {
	text := string(textData)
	lines := 0

	for _, char := range text {
		if char == '\n' {
			lines += 1
		}
	}

	return lines
}

func getWordCount(textData []byte) int {
	words := 0
	text := string(textData)
	whiteSpaceChars := map[rune]bool{
		' ':  true,
		'\t': true,
		'\r': true,
		'\n': true,
		'\v': true,
		'\f': true,
	}

	prevChar := ' '

	for _, char := range text {
		_, prevCharIsWhitespace := whiteSpaceChars[prevChar]
		_, currCharIsWhitespace := whiteSpaceChars[char]
		if !prevCharIsWhitespace && currCharIsWhitespace {
			words += 1
		}
		prevChar = char
	}

	return words
}

func getCharacterCount(textData []byte) int {
	characters := 0
	text := string(textData)

	// range iterates over each unicode code point as opposed to
	// len which would give count multi-byte code points as multiple chars
	for range text {
		characters += 1
	}

	return characters
}
