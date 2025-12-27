package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseLists(input string) ([]int, []int, error) {
	var left, right []int

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) != 2 {
			return nil, nil, fmt.Errorf("invalid line: %s", line)
		}

		leftVal, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid left value: %s", fields[0])
		}

		rightVal, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid right value: %s", fields[1])
		}

		left = append(left, leftVal)
		right = append(right, rightVal)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return left, right, nil
}

func calculateTotalDistance(left, right []int) int {
	leftSorted := make([]int, len(left))
	rightSorted := make([]int, len(right))

	copy(leftSorted, left)
	copy(rightSorted, right)

	sort.Ints(leftSorted)
	sort.Ints(rightSorted)

	totalDistance := 0
	for i := 0; i < len(leftSorted); i++ {
		distance := leftSorted[i] - rightSorted[i]
		if distance < 0 {
			distance = -distance
		}
		totalDistance += distance
	}

	return totalDistance
}

func calculateSimilarityScore(left, right []int) int {
	rightCounts := make(map[int]int)
	for _, num := range right {
		rightCounts[num]++
	}

	similarityScore := 0
	for _, num := range left {
		count := rightCounts[num]
		similarityScore += num * count
	}

	return similarityScore
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input.txt: %v\n", err)
		os.Exit(1)
	}

	left, right, err := parseLists(string(data))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing lists: %v\n", err)
		os.Exit(1)
	}

	part1 := calculateTotalDistance(left, right)
	fmt.Printf("Part 1 - Total distance: %d\n", part1)

	part2 := calculateSimilarityScore(left, right)
	fmt.Printf("Part 2 - Similarity score: %d\n", part2)
}
