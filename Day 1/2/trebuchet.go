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

	textByteArray, err := os.ReadFile("input.txt")
	if err != nil || textByteArray == nil {
		log.Fatal("Reading file went wrong. Does input.txt exist?")
	}
	lines := strings.Split(string(textByteArray), "\n")

	totalcoords := 0

	for index := range lines {
		// fmt.Println(findCoordsAdvanced(lines[index]))
		totalcoords += findCoordsAdvanced(lines[index])
	}
	fmt.Printf("Total: %v\n", totalcoords)

}

func findCoordsAdvanced(line string) (out int) {
	numbers := [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	first := "-1"
	last := "-1"

	// iterate over line
	for i := range line {
		// check if character is a valid number first
		num, err := strconv.Atoi(string(line[i]))
		if err == nil {
			fmt.Printf("found %v in '%s'\n", num, line)
			if first == "-1" {
				first = strconv.Itoa(num)
			}
			last = strconv.Itoa(num)
			fmt.Printf("first: %v, last: %v\n", first, last)

		}

		// otherwise, check if it's the first letter in a string number by iterating over numbers
		for j := range numbers {

			// if the rest of the line is less long than the length of the number, pass and check next number
			if i+len(numbers[j]) > len(line) {
				continue
			}

			// get a slice that starts at the index and ends at the index + the length of the possible number ("zero" is 4 chars for ex.)
			possibleNumber := line[i:(i + len(numbers[j]))]

			// check if slice is equal to number being checked
			if possibleNumber == numbers[j] {
				fmt.Printf("found %v at %v in %v\n", numbers[j], j, possibleNumber)
				if first == "-1" {
					first = strconv.Itoa(j)
				}
				last = strconv.Itoa(j)
				fmt.Printf("first: %v, last: %v\n", first, last)
			} else {
			}

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
	// fmt.Printf("first: %s, last: %s\n	in '%s'", first, last, line)
	return out
}

/*
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
*/
