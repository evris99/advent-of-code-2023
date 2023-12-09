package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	schematicsBytes, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatalf("failed reading file: %s", err)
	}

	// Split the schematics into lines
	schematics := strings.Split(string(schematicsBytes), "\n")

	sum := 0
	gearRegexp := regexp.MustCompile(`\*`)
	for i, line := range schematics {
		// Find all the gear symbols in the line
		gears := gearRegexp.FindAllStringIndex(line, -1)
		if gears == nil {
			continue
		}

		for _, gear := range gears {
			gearRatio, err := getGearRatio(schematics, i, gear[0])
			if err != nil {
				log.Fatalf("failed to get gear ratio: %s", err)
			}
			sum += gearRatio
		}
	}

	log.Printf("The of the gear ratios is %d", sum)
}

// getGearRatio returns the gear ratio for the given coordinates
// This functions does not check if the coordinates are a valid gear symbol (*)
// It returns 0 if there is no gear ratio
func getGearRatio(data []string, i, j int) (int, error) {
	partNumbers := make([]int, 0)
	for k := i - 1; k <= i+1; k++ {
		for l := j - 1; l <= j+1; l++ {

			if isOutOfBounds(data, k, l) {
				continue
			}

			if unicode.IsDigit(rune(data[k][l])) {
				startIndex, endIndex := getPartNumberIndexes(data[k], l)

				partNumber, err := strconv.Atoi(data[k][startIndex : endIndex+1])
				if err != nil {
					return 0, fmt.Errorf("failed to convert part number to int: %s", err)
				}

				partNumbers = append(partNumbers, partNumber)

				// Replace the part number with dots, so we don't count it again
				line := []rune(data[k])
				for m := startIndex; m <= endIndex; m++ {
					line[m] = '.'
				}

				data[k] = string(line)
			}
		}
	}

	res := 0
	if len(partNumbers) == 2 {
		res = partNumbers[0] * partNumbers[1]
	}

	return res, nil
}

// getPartNumber returns the part number that is at the given coordinates
func getPartNumberIndexes(data string, x int) (int, int) {
	// Find the start index of the part number
	startIndex := x
	for i := x; i >= 0 && unicode.IsDigit(rune(data[i])); i-- {
		startIndex = i
	}

	// Find the end index of the part number
	endIndex := startIndex
	for i := startIndex; i < len(data) && unicode.IsDigit(rune(data[i])); i++ {
		endIndex = i
	}

	return startIndex, endIndex
}

// isOutOfBounds checks if the given coordinates are out of bounds
func isOutOfBounds(data []string, x, y int) bool {
	return x < 0 || x >= len(data) || y < 0 || y >= len(data[x])
}
