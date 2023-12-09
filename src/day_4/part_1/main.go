package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		points, err := getPoints(scanner.Text())
		if err != nil {
			log.Fatalf("failed to get calibration value: %s", err)
		}

		sum += points
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("failed reading file: %s", err)
	}

	log.Printf("They are worth %d points", sum)
}

// getPoints returns the number of points for a line of text
func getPoints(text string) (int, error) {
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
	matchCount := getMatchCount(numbersToMatch, winningNumbers)

	res := 0
	if matchCount != 0 {
		res = int(math.Pow(2, float64(matchCount-1)))
	}

	return res, nil
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
