package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	// Read file line by line and sum the calibration values from each line
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := getCalibrationValue(scanner.Text())
		if err != nil {
			log.Fatalf("failed to get calibration value: %s", err)
		}

		sum += value
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("failed reading file: %s", err)
	}

	log.Printf("The sum of all calibration values is: %d", sum)
}

// getCalibrationValue returns the calibration value from a line of text
func getCalibrationValue(text string) (int, error) {
	firstDigit := -1
	lastDigit := -1
	for _, char := range text {
		// If the character is a digit, set the first and last digit
		if unicode.IsDigit(char) {
			var err error
			if firstDigit == -1 {
				firstDigit, err = strconv.Atoi(string(char))
				if err != nil {
					return 0, fmt.Errorf("failed parsing int: %s", err)
				}
			}

			lastDigit, err = strconv.Atoi(string(char))
			if err != nil {
				return 0, fmt.Errorf("failed parsing int: %s", err)
			}
		}
	}

	return firstDigit*10 + lastDigit, nil
}
