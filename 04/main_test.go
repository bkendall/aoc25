package main

import "testing"

func TestCountNeighbors(t *testing.T) {
	// Setup a simple world
	// @ . .
	// . @ .
	// @ @ @
	//
	// Positions:
	// (0,0), (1,1), (0,2), (1,2), (2,2)

	world := World{
		0: {0: true, 2: true},
		1: {1: true, 2: true},
		2: {2: true},
	}

	tests := []struct {
		x, y     int
		expected int
	}{
		{1, 1, 4}, // Center
		{0, 0, 1}, // Top-left corner
		{1, 2, 3}, // Bottom-middle
		{0, 1, 4}, // Middle-left (empty spot checking neighbors)
	}

	for _, tt := range tests {
		got := countNeighbors(world, tt.x, tt.y)
		if got != tt.expected {
			t.Errorf("countNeighbors(world, %d, %d) = %d; want %d", tt.x, tt.y, got, tt.expected)
		}
	}
}
