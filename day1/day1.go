package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// get puzzle input here:
// https://adventofcode.com/2024/day/1/input

func Day1() {
	file, err := os.Open("./day1/input.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}

	defer file.Close()

	input, err := os.ReadFile("./day1/input.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// fmt.Println(string(input))
	result := GetTotal(string(input))

	fmt.Println(result)
}

func GetTotal(input string) int {
	locations1, locations2, err := ParseInput(input)

	if err != nil {
		log.Fatalf("could not parse input: %v", err)
	}

	sort.Ints(locations1)
	sort.Ints(locations2)

	total := 0

	for i := 0; i < len(locations1); i++ {
		if locations1[i] > locations2[i] {
			diff := locations1[i] - locations2[i]
			total = total + diff
		} else if locations2[i] > locations1[i] {
			diff := locations2[i] - locations1[i]
			total = total + diff
		}
	}

	return total

}

func ParseInput(input string) ([]int, []int, error) {
	locations1 := []int{}
	locations2 := []int{}

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		fields := strings.Fields(line)

		firstNum, err := strconv.Atoi(string(fields[0]))
		if err != nil {
			return nil, nil, err
		}

		secondNum, err := strconv.Atoi(string(fields[1]))
		if err != nil {
			return nil, nil, err
		}

		locations1 = append(locations1, firstNum)
		locations2 = append(locations2, secondNum)

	}
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return locations1, locations2, nil
}
