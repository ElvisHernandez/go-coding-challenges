package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	filePath := getFilePathArg()
	data, err := os.ReadFile("./" + filePath)
	if err != nil {
		log.Fatal(err)
	}

	printByteCount := flag.Bool("c", false, "Prints the byte count of a file")
	printLineCount := flag.Bool("l", false, "Prints the line count of a file")
	flag.Parse()

	if *printByteCount {
		fmt.Printf("%v %v\n", len(data), filePath)
	}

	if *printLineCount {
		fmt.Printf("%v %v", getLineCount(data), filePath)
	}
}

func getFilePathArg() string {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatal("No arguments were passed")
	}

	filePath := args[len(args)-1]

	if strings.HasPrefix(filePath, "-") {
		log.Fatal("Last argument must be a file path")
	}

	return filePath
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
