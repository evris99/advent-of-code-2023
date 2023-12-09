package main

import (
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

	schematics := strings.Split(string(schematicsBytes), "\n")

	sum := 0
	partNumberRegex := regexp.MustCompile("[0-9]+")
	for i, line := range schematics {
		// Find all the numbers in the line
		partNumberIndexes := partNumberRegex.FindAllStringIndex(line, -1)
		if partNumberIndexes == nil {
			continue
		}

		for _, partNumberIndex := range partNumberIndexes {

			// For each digit found, check if it is adjacent to a symbol
			isSymbolAdjacent := false
			for j := partNumberIndex[0]; j < partNumberIndex[1]; j++ {
				if hasAdjacentSymbol(schematics, i, j) {
					isSymbolAdjacent = true
					break
				}
			}

			// If any digit is adjacent to a symbol, add it to the sum
			if isSymbolAdjacent {
				// Convert the part number string to an integer
				partNumber, err := strconv.Atoi(line[partNumberIndex[0]:partNumberIndex[1]])
				if err != nil {
					log.Fatalf("failed to convert part number to integer: %s", err)
				}

				sum += partNumber
			}
		}
	}

	log.Printf("The sum of the part numbers is %d", sum)
}

// hasAdjacentSymbol checks if the symbol is adjacent to the given coordinates
func hasAdjacentSymbol(data []string, x, y int) bool {
	for k := x - 1; k <= x+1; k++ {
		for l := y - 1; l <= y+1; l++ {
			if k < 0 || k >= len(data) || l < 0 || l >= len(data[k]) {
				continue
			}

			if IsSymbol(data[k][l]) {
				return true
			}
		}
	}

	return false
}

// IsSymbol checks if the given character is a symbol
func IsSymbol(char byte) bool {
	return !unicode.IsNumber(rune(char)) && char != '.'
}
