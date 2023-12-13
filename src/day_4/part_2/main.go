package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	content string
	copies  int
}

func main() {
	scratchCardBytes, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatalf("failed reading file: %s", err)
	}

	// Create a slice of cards with one copy of each card
	scratchCards := make([]Card, 0)
	for _, scratchCard := range strings.Split(string(scratchCardBytes), "\n") {
		scratchCards = append(scratchCards, Card{content: scratchCard, copies: 1})
	}

	count := 0
	for len(scratchCards) > 0 {
		matches, err := getMatches(scratchCards[0].content)
		if err != nil {
			log.Fatalf("failed getting points: %s", err)
		}

		// For each match, add a copy of the scratch card
		for i := 0; i < matches; i++ {
			scratchCards[i+1].copies += scratchCards[0].copies
		}

		// Increase the total number of cards processed
		count += scratchCards[0].copies

		// Remove the processed card
		scratchCards = scratchCards[1:]
	}

	log.Printf("The total number of cards is %d", count)
}

// getMatchCount returns the number of matching numbers between the winning and the played numbers
func getMatches(text string) (int, error) {
	numbers := strings.Split(strings.Split(text, ":")[1], "|")

	// Get the winning numbers
	winningNumbersStrings := strings.Split(strings.TrimSpace(numbers[0]), " ")
	winningNumbers, err := getNumbers(winningNumbersStrings)
	if err != nil {
		return 0, fmt.Errorf("failed to get winning numbers: %s", err)
	}

	// Get matching numbers count
	numbersToMatchString := strings.Split(strings.TrimSpace(numbers[1]), " ")
	numbersToMatch, err := getNumbers(numbersToMatchString)
	if err != nil {
		return 0, fmt.Errorf("failed to get numbers to match: %s", err)
	}

	// Get the number of points
	return getMatchCount(numbersToMatch, winningNumbers), nil
}

// getMatchCount returns the number of matching numbers between the two given slices
func getMatchCount(numbers []int, winningNumbers []int) int {
	count := 0
	for _, number := range numbers {
		for _, winningNumber := range winningNumbers {
			if number == winningNumber {
				count++
			}
		}
	}

	return count
}

// getNumbers returns a slice of integers from a string containing numbers space separated
func getNumbers(numbersString []string) ([]int, error) {
	numbers := make([]int, 0)
	for _, numberString := range numbersString {
		if numberString == "" {
			continue
		}

		number, err := strconv.Atoi(numberString)
		if err != nil {
			return nil, fmt.Errorf("failed parsing int: %s", err)
		}

		numbers = append(numbers, number)
	}

	return numbers, nil
}
