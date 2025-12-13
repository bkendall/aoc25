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
				TargetState: []bool{false, true, true, false},
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
