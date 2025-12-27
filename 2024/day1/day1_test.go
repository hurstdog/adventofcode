package main

import (
	"testing"
)

func TestParseLists(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`

	left, right, err := parseLists(input)
	if err != nil {
		t.Fatalf("parseLists failed: %v", err)
	}

	expectedLeft := []int{3, 4, 2, 1, 3, 3}
	expectedRight := []int{4, 3, 5, 3, 9, 3}

	if len(left) != len(expectedLeft) {
		t.Errorf("left list length = %d, want %d", len(left), len(expectedLeft))
	}

	for i := range left {
		if left[i] != expectedLeft[i] {
			t.Errorf("left[%d] = %d, want %d", i, left[i], expectedLeft[i])
		}
	}

	if len(right) != len(expectedRight) {
		t.Errorf("right list length = %d, want %d", len(right), len(expectedRight))
	}

	for i := range right {
		if right[i] != expectedRight[i] {
			t.Errorf("right[%d] = %d, want %d", i, right[i], expectedRight[i])
		}
	}
}

func TestCalculateTotalDistance(t *testing.T) {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}

	result := calculateTotalDistance(left, right)
	expected := 11

	if result != expected {
		t.Errorf("calculateTotalDistance = %d, want %d", result, expected)
	}
}

func TestCalculateSimilarityScore(t *testing.T) {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}

	result := calculateSimilarityScore(left, right)
	expected := 31

	if result != expected {
		t.Errorf("calculateSimilarityScore = %d, want %d", result, expected)
	}
}
