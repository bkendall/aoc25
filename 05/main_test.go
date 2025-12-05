package main

import (
	"testing"
)

func TestSolvePart2(t *testing.T) {
	tests := []struct {
		name     string
		ranges   []Range
		expected int
	}{
		{
			name:     "No overlap",
			ranges:   []Range{{1, 5}, {10, 15}},
			expected: 5 + 6, // 11
		},
		{
			name:     "Overlap",
			ranges:   []Range{{1, 5}, {4, 8}},
			expected: 8, // 1..8
		},
		{
			name:     "Contained",
			ranges:   []Range{{1, 10}, {2, 5}},
			expected: 10,
		},
		{
			name:     "Touching",
			ranges:   []Range{{1, 5}, {6, 10}},
			expected: 10,
		},
		{
			name:     "Unsorted",
			ranges:   []Range{{10, 15}, {1, 5}},
			expected: 11,
		},
		{
			name:     "Multiple overlaps",
			ranges:   []Range{{1, 5}, {2, 6}, {8, 10}, {9, 12}},
			// 1-5, 2-6 -> 1-6 (size 6)
			// 8-10, 9-12 -> 8-12 (size 5)
			// Total 11
			expected: 11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solvePart2(tt.ranges); got != tt.expected {
				t.Errorf("solvePart2() = %v, want %v", got, tt.expected)
			}
		})
	}
}