package day3

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

type NumPair struct {
	first  int
	second int
}

func Day3() {
	start := time.Now()
	var result int

	file, err := os.Open("./day3/input.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	input, err := os.ReadFile("./day3/input.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	pairs := FindCommands(string(input))
	result = ReduceCommands(pairs)

	elapsed := time.Since(start)
	fmt.Printf("Day 3 result: %d\r\nCompleted in %s", result, elapsed)
}

func FindCommands(input string) []NumPair {
	output := []NumPair{}

	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	digits := regexp.MustCompile(`\d+`)

	matches := re.FindAllString(input, -1)

	for _, match := range matches {
		rawDigits := digits.FindAllString(match, -1)

		first, err := strconv.Atoi(rawDigits[0])
		if err != nil {
			log.Fatalf("Could not properly convert string to int: %v", err)
		}
		second, err := strconv.Atoi(rawDigits[1])
		if err != nil {
			log.Fatalf("Could not properly convert string to int: %v", err)
		}

		output = append(output, NumPair{first: first, second: second})
	}

	return output
}

func ReduceCommands(pairs []NumPair) int {
	total := 0

	for _, pair := range pairs {
		mult := pair.first * pair.second
		total = total + mult
	}

	return total
}
