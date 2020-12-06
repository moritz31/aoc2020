package main

import "testing"

func TestCalculation(t *testing.T) {
	sampleInput := []int{1721, 979, 366, 299, 675, 1456}

	resultTuple, resultTriple := findMatching(sampleInput)

	if resultTuple != 514579 {
		t.Errorf("Sum was incorrect, got %d, want: %d", resultTuple, 514579)
	}

	if resultTriple != 241861950 {
		t.Errorf("Sum2 was incorrect, got %d, want: %d", resultTriple, 241861950)
	}
}
