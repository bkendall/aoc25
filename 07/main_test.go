package main

import (
	"strings"
	"testing"
)

var sampleInput = strings.TrimSpace(`
.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............
`)

func TestPart1(t *testing.T) {
	w := parseInput(sampleInput)
	startX := findStart(w)
	if startX == -1 {
		t.Fatal("Could not find start")
	}
	expected := 21
	if got := solvePart1(w, startX); got != expected {
		t.Errorf("Part1 = %d, want %d", got, expected)
	}
}

func TestPart2(t *testing.T) {
	w := parseInput(sampleInput)
	startX := findStart(w)
	if startX == -1 {
		t.Fatal("Could not find start")
	}
	expected := 40
	if got := solvePart2(w, startX); got != expected {
		t.Errorf("Part2 = %d, want %d", got, expected)
	}
}
