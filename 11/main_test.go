package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		input string
		want  *Node
	}{
		{
			input: "you: b c",
			want: &Node{
				ID: "you",
				Connections: map[string]*Node{
					"b": {ID: "b"},
					"c": {ID: "c"},
				},
			},
		},
	}

	for _, tt := range tests {
		got, _ := parseInput(tt.input)
		if diff := cmp.Diff(got, tt.want); diff != "" {
			t.Errorf("parseInput(%q) mismatch (-want +got):\n%s", tt.input, diff)
		}
	}
}
