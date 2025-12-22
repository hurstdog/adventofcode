package main

import (
	"testing"
)

func TestParseRotation(t *testing.T) {
	tests := []struct {
		input    string
		expected Rotation
		hasError bool
	}{
		{"L68", Rotation{"L", 68}, false},
		{"R48", Rotation{"R", 48}, false},
		{"L1", Rotation{"L", 1}, false},
		{"R1000", Rotation{"R", 1000}, false},
		{"L0", Rotation{"L", 0}, false},
		{"  R48  ", Rotation{"R", 48}, false}, // with whitespace
		{"X48", Rotation{}, true},              // invalid direction
		{"L", Rotation{}, true},                // no distance
		{"", Rotation{}, true},                 // empty string
		{"Labc", Rotation{}, true},             // invalid distance
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseRotation(tt.input)
			if tt.hasError {
				if err == nil {
					t.Errorf("ParseRotation(%q) expected error, got nil", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("ParseRotation(%q) unexpected error: %v", tt.input, err)
				}
				if result != tt.expected {
					t.Errorf("ParseRotation(%q) = %v, want %v", tt.input, result, tt.expected)
				}
			}
		})
	}
}

func TestApplyRotation(t *testing.T) {
	tests := []struct {
		position int
		rotation Rotation
		expected int
	}{
		// Examples from problem
		{50, Rotation{"L", 68}, 82},
		{82, Rotation{"L", 30}, 52},
		{52, Rotation{"R", 48}, 0},
		{0, Rotation{"L", 5}, 95},
		{95, Rotation{"R", 60}, 55},
		{55, Rotation{"L", 55}, 0},
		{0, Rotation{"L", 1}, 99},
		{99, Rotation{"L", 99}, 0},
		{0, Rotation{"R", 14}, 14},
		{14, Rotation{"L", 82}, 32},

		// Edge cases
		{0, Rotation{"L", 1}, 99},   // wrap from 0 to 99
		{99, Rotation{"R", 1}, 0},   // wrap from 99 to 0
		{50, Rotation{"R", 0}, 50},  // no movement
		{50, Rotation{"L", 0}, 50},  // no movement
		{50, Rotation{"R", 100}, 50}, // full circle
		{50, Rotation{"L", 100}, 50}, // full circle
		{50, Rotation{"R", 250}, 0},  // multiple circles
		{11, Rotation{"R", 8}, 19},   // example from problem
		{19, Rotation{"L", 19}, 0},   // example from problem
		{5, Rotation{"L", 10}, 95},   // example from problem
		{95, Rotation{"R", 5}, 0},    // example from problem (could also be 0)
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := ApplyRotation(tt.position, tt.rotation)
			if result != tt.expected {
				t.Errorf("ApplyRotation(%d, %v) = %d, want %d",
					tt.position, tt.rotation, result, tt.expected)
			}
		})
	}
}

func TestCountZeroCrossings(t *testing.T) {
	tests := []struct {
		position int
		rotation Rotation
		expected int
		desc     string
	}{
		// From example (Part 2 specific details)
		{50, Rotation{"L", 68}, 1, "L68 from 50 crosses 0 once"},
		{82, Rotation{"L", 30}, 0, "L30 from 82 doesn't cross 0"},
		{52, Rotation{"R", 48}, 1, "R48 from 52 lands on 0"},
		{0, Rotation{"L", 5}, 0, "L5 from 0 doesn't cross 0"},
		{95, Rotation{"R", 60}, 1, "R60 from 95 crosses 0 once"},
		{55, Rotation{"L", 55}, 1, "L55 from 55 lands on 0"},
		{0, Rotation{"L", 1}, 0, "L1 from 0 doesn't cross 0"},
		{99, Rotation{"L", 99}, 1, "L99 from 99 lands on 0"},
		{0, Rotation{"R", 14}, 0, "R14 from 0 doesn't cross 0"},
		{14, Rotation{"L", 82}, 1, "L82 from 14 crosses 0 once"},

		// Edge cases
		{50, Rotation{"R", 1000}, 10, "R1000 from 50 crosses 0 ten times"},
		{0, Rotation{"R", 100}, 1, "R100 from 0 crosses 0 once"},
		{0, Rotation{"R", 200}, 2, "R200 from 0 crosses 0 twice"},
		{99, Rotation{"R", 1}, 1, "R1 from 99 lands on 0"},
		{0, Rotation{"L", 100}, 1, "L100 from 0 crosses 0 once (backwards)"},
		{1, Rotation{"L", 1}, 1, "L1 from 1 lands on 0"},
		{50, Rotation{"R", 0}, 0, "no rotation"},
		{50, Rotation{"L", 0}, 0, "no rotation"},
		{0, Rotation{"R", 0}, 0, "no rotation from 0"},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			result := CountZeroCrossings(tt.position, tt.rotation)
			if result != tt.expected {
				t.Errorf("CountZeroCrossings(%d, %v) = %d, want %d (%s)",
					tt.position, tt.rotation, result, tt.expected, tt.desc)
			}
		})
	}
}

func TestSolvePart1(t *testing.T) {
	// Example from problem
	exampleRotations := []Rotation{
		{"L", 68},
		{"L", 30},
		{"R", 48},
		{"L", 5},
		{"R", 60},
		{"L", 55},
		{"L", 1},
		{"L", 99},
		{"R", 14},
		{"L", 82},
	}

	result := SolvePart1(exampleRotations)
	expected := 3
	if result != expected {
		t.Errorf("SolvePart1(example) = %d, want %d", result, expected)
	}
}

func TestSolvePart2(t *testing.T) {
	// Example from problem
	exampleRotations := []Rotation{
		{"L", 68},
		{"L", 30},
		{"R", 48},
		{"L", 5},
		{"R", 60},
		{"L", 55},
		{"L", 1},
		{"L", 99},
		{"R", 14},
		{"L", 82},
	}

	result := SolvePart2(exampleRotations)
	expected := 6
	if result != expected {
		t.Errorf("SolvePart2(example) = %d, want %d", result, expected)
	}
}

// Test the step-by-step process for Part 1 to verify the example
func TestPart1StepByStep(t *testing.T) {
	rotations := []Rotation{
		{"L", 68}, // 50 -> 82
		{"L", 30}, // 82 -> 52
		{"R", 48}, // 52 -> 0  (count: 1)
		{"L", 5},  // 0 -> 95
		{"R", 60}, // 95 -> 55
		{"L", 55}, // 55 -> 0  (count: 2)
		{"L", 1},  // 0 -> 99
		{"L", 99}, // 99 -> 0  (count: 3)
		{"R", 14}, // 0 -> 14
		{"L", 82}, // 14 -> 32
	}

	expectedPositions := []int{82, 52, 0, 95, 55, 0, 99, 0, 14, 32}
	expectedZeros := 3

	position := startPosition
	zerosCount := 0

	for i, rotation := range rotations {
		position = ApplyRotation(position, rotation)
		if position == 0 {
			zerosCount++
		}

		if position != expectedPositions[i] {
			t.Errorf("After rotation %d (%v): position = %d, want %d",
				i, rotation, position, expectedPositions[i])
		}
	}

	if zerosCount != expectedZeros {
		t.Errorf("Total zeros count = %d, want %d", zerosCount, expectedZeros)
	}
}

// Test the step-by-step process for Part 2 to verify the example
func TestPart2StepByStep(t *testing.T) {
	rotations := []Rotation{
		{"L", 68}, // 50 -> 82, crosses 0 once
		{"L", 30}, // 82 -> 52, no crossing
		{"R", 48}, // 52 -> 0, lands on 0 (count 1)
		{"L", 5},  // 0 -> 95, no crossing
		{"R", 60}, // 95 -> 55, crosses 0 once
		{"L", 55}, // 55 -> 0, lands on 0 (count 1)
		{"L", 1},  // 0 -> 99, no crossing
		{"L", 99}, // 99 -> 0, lands on 0 (count 1)
		{"R", 14}, // 0 -> 14, no crossing
		{"L", 82}, // 14 -> 32, crosses 0 once
	}

	expectedCrossings := []int{1, 0, 1, 0, 1, 1, 0, 1, 0, 1}
	expectedTotal := 6

	position := startPosition
	totalCrossings := 0

	for i, rotation := range rotations {
		crossings := CountZeroCrossings(position, rotation)
		totalCrossings += crossings

		if crossings != expectedCrossings[i] {
			t.Errorf("Rotation %d (%v from %d): crossings = %d, want %d",
				i, rotation, position, crossings, expectedCrossings[i])
		}

		position = ApplyRotation(position, rotation)
	}

	if totalCrossings != expectedTotal {
		t.Errorf("Total crossings = %d, want %d", totalCrossings, expectedTotal)
	}
}
