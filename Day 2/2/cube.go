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

	totalPowers := 0.0
	for i := range lines {
		line := lines[i]
		totalPowers += getPower(line)

	}
	fmt.Printf("Output: %v", totalPowers)

}

func getPower(line string) (power float64) {
	// ignore empty lines
	if line == "" {
		return 0
	}

	// cut off "Game x:"
	semiColonIndex := strings.Index(line, ":")
	line = line[semiColonIndex+1:]

	usedColorsMin := map[string]int{"red": 0, "green": 0, "blue": 0}

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
			// add to tracker
			if usedColorsMin[pair[1]] < pairAmount {
				usedColorsMin[pair[1]] = pairAmount
			}
		}
	}

	power = float64(usedColorsMin["red"]) * float64(usedColorsMin["green"]) * float64(usedColorsMin["blue"])
	return power
}
