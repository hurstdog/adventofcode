package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const dialSize = 100
const startPosition = 50

// Rotation represents a single rotation instruction
type Rotation struct {
	Direction string // "L" or "R"
	Distance  int
}

// ParseRotation parses a rotation string like "L68" or "R48"
func ParseRotation(s string) (Rotation, error) {
	s = strings.TrimSpace(s)
	if len(s) < 2 {
		return Rotation{}, fmt.Errorf("invalid rotation: %s", s)
	}

	direction := s[:1]
	if direction != "L" && direction != "R" {
		return Rotation{}, fmt.Errorf("invalid direction: %s", direction)
	}

	distance, err := strconv.Atoi(s[1:])
	if err != nil {
		return Rotation{}, fmt.Errorf("invalid distance in %s: %w", s, err)
	}

	return Rotation{Direction: direction, Distance: distance}, nil
}

// ApplyRotation applies a rotation to the current position and returns the new position
func ApplyRotation(currentPos int, rotation Rotation) int {
	if rotation.Direction == "R" {
		return (currentPos + rotation.Distance) % dialSize
	}
	// Left rotation
	return (currentPos - rotation.Distance%dialSize + dialSize) % dialSize
}

// CountZeroCrossings counts how many times the dial passes through position 0
// during a rotation (including if it lands on 0)
func CountZeroCrossings(currentPos int, rotation Rotation) int {
	if rotation.Direction == "R" {
		// Right rotation: count how many times we cross 0
		// Going from currentPos to currentPos+distance
		return (currentPos + rotation.Distance) / dialSize - currentPos / dialSize
	}
	// Left rotation: count how many times we pass through 0 going backwards
	// When going left by D from position X, we visit X-1, X-2, ..., X-D (mod 100)
	// We hit 0 when k ≡ X (mod 100) for k in [1, D]
	// If X = 0: we hit 0 at k = 100, 200, ... (count = D / 100)
	// If X > 0: we hit 0 at k = X, X+100, X+200, ... (count = (D + 100 - X) / 100)
	if currentPos == 0 {
		return rotation.Distance / dialSize
	}
	return (rotation.Distance + dialSize - currentPos) / dialSize
}

// SolvePart1 counts how many times the dial lands on 0 after any rotation
func SolvePart1(rotations []Rotation) int {
	position := startPosition
	count := 0

	for _, rotation := range rotations {
		position = ApplyRotation(position, rotation)
		if position == 0 {
			count++
		}
	}

	return count
}

// SolvePart2 counts how many times the dial points at 0 during any click
func SolvePart2(rotations []Rotation) int {
	position := startPosition
	count := 0

	for _, rotation := range rotations {
		count += CountZeroCrossings(position, rotation)
		position = ApplyRotation(position, rotation)
	}

	return count
}

// ReadRotations reads rotations from a file
func ReadRotations(filename string) ([]Rotation, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var rotations []Rotation
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		rotation, err := ParseRotation(line)
		if err != nil {
			return nil, err
		}
		rotations = append(rotations, rotation)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return rotations, nil
}

func main() {
	rotations, err := ReadRotations("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	part1 := SolvePart1(rotations)
	part2 := SolvePart2(rotations)

	// Part 1 Answer: 1102
	fmt.Printf("Part 1: %d\n", part1)

	// Part 2 Answer: 6175
	fmt.Printf("Part 2: %d\n", part2)
}
