package main

import (
	"bufio"
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
		sum += getPower(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("failed reading file: %s", err)
	}

	log.Printf("The sum of power is %d", sum)
}

// isPossible checks if the game is possible
func getPower(text string) int {
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

	return cubeMaximums["red"] * cubeMaximums["green"] * cubeMaximums["blue"]
}
