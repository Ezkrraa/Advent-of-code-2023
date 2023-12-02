package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")
	if lines == nil {
		log.Fatal("File 'input.txt' was empty.")
	}

	validGames := 0
	for i := range lines {
		line := lines[i]
		if validateLine(line) {
			validGames += i + 1
			// fmt.Printf("Game %v was valid\n", i+1)
		} else {
			// fmt.Printf("Game %v was invalid\n", i+1)
		}

	}
	fmt.Printf("Output: %v", validGames)

}

func validateLine(line string) (validity bool) {
	// ignore empty lines
	if line == "" {
		return false
	}

	// cut off "Game x:"
	semiColonIndex := strings.Index(line, ":")
	line = line[semiColonIndex+1:]

	validity = true

	// use a map to store per-color maximums
	colorMax := map[string]int{"red": 12, "green": 13, "blue": 14}

	//iterate over each set in a line
	sets := strings.Split(line, ";")
	for i := range sets {
		set := sets[i]

		// iterate over each color in a set
		colors := strings.Split(set, ", ")
		for j := range colors {
			colors[j] = strings.TrimSpace(colors[j])

			// split pair and check map to see if below max
			pair := strings.Split(colors[j], " ")
			pairAmount, err := strconv.Atoi(pair[0])
			if err != nil {
				fmt.Printf("tried to convert %v to string in line %v\n", pair[0], line)
				log.Fatal(err)
			}
			if pairAmount > colorMax[pair[1]] {
				validity = false
			}

		}
	}

	return validity
}
