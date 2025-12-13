package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		input string
		want  *Machine
	}{
		{
			input: "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}",
			want: &Machine{
				NumLights:   4,
				LightStates: make([]bool, 4),
				Buttons:     6,
				ButtonTriggers: map[int][]int{
					0: {3},
					1: {1, 3},
					2: {2},
					3: {2, 3},
					4: {0, 2},
					5: {0, 1},
				},
				TargetState:   []bool{false, true, true, false},
				TargetJoltage: []int{3, 5, 4, 7},
			},
		},
	}

	for _, tt := range tests {
		g := parseInput(tt.input)
		want := []*Machine{tt.want}
		if diff := cmp.Diff(want, g); diff != "" {
			t.Errorf("ParseInput(%q) mismatch (-want +got):\n%s", tt.input, diff)
		}
	}
}

func TestSolvePartOne(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}",
			want:  2,
		},
	}

	for _, tt := range tests {
		got := solvePartOne(parseInput(tt.input))
		if got != tt.want {
			t.Errorf("SolvePartOne(%q) = %d, want %d", tt.input, got, tt.want)
		}
	}
}

func TestSolvePartTwo(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}",
			want:  10,
		},
	}

	for _, tt := range tests {
		got := solvePartTwo(parseInput(tt.input))
		if got != tt.want {
			t.Errorf("SolvePartTwo(%q) = %d, want %d", tt.input, got, tt.want)
		}
	}
}
