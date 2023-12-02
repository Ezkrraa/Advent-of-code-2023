package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// will read "input.txt" from this directory
// then get first and last number of each row
// concat into string, then convert into int.
// add all ints together and return that.

func main() {
	//
	textByteArray, err := os.ReadFile("input.txt")
	if err != nil || textByteArray == nil {
		log.Fatal("Reading file went wrong. Does input.txt exist?")
	}

	lines := strings.Split(string(textByteArray), "\n")

	totalcoords := 0

	for index := range lines {
		totalcoords += findCoords(lines[index])
	}
	fmt.Println(totalcoords)
}

func findCoords(line string) (out int) {
	first := "-1"
	last := "-1"

	for index := range line {
		char := string(line[index])

		num, err := strconv.Atoi(char)
		if err == nil {
			if first == "-1" {
				first = strconv.Itoa(num)
			}
			last = strconv.Itoa(num)
		}
	}
	if first == "-1" || last == "-1" {
		log.Fatalf("No 2 ints found in %s.", line)
	}

	out, err := strconv.Atoi(first + last)
	if err != nil {
		fmt.Printf("first: %s, last: %s\n", first, last)
		log.Fatal(err)
	}
	return out
}
