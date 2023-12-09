package main

import (
	"bufio"
	"fmt"
	"log"
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

	// Read file line by line and sum the the IDs for each possible game
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		gameID, err := getGameID(line)
		if err != nil {
			log.Fatalf("failed to get game ID: %s", err)
		}

		if isPossible(line) {
			sum += gameID
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("failed reading file: %s", err)
	}

	log.Printf("The sum of the IDs for the possible games is %d", sum)
}

// getGameID parses the game ID from the text
func getGameID(text string) (int, error) {
	var gameID int
	_, err := fmt.Sscanf(text, "Game %d:", &gameID)
	if err != nil {
		return 0, fmt.Errorf("failed to parse game ID: %s", err)
	}

	return gameID, nil
}

// isPossible checks if the game is possible
func isPossible(text string) bool {
	// Remove the "Game X:" part and trim the spaces
	sets := strings.TrimSpace(strings.SplitAfter(text, ":")[1])

	cubeMaximums := map[string]int{
		"red":   0,
		"blue":  0,
		"green": 0,
	}

	num := 0
	words := strings.Split(sets, " ")
	for _, word := range words {
		// Remove the trailing and leading commas and semicolons
		word = strings.Trim(word, ",;")

		// Check if the word is a number or a color
		if numOfCubes, err := strconv.Atoi(word); err == nil {
			num = numOfCubes
		} else if max, ok := cubeMaximums[word]; ok && num > max {
			cubeMaximums[word] = num
		}
	}

	return cubeMaximums["red"] <= 12 && cubeMaximums["green"] <= 13 && cubeMaximums["blue"] <= 14
}
