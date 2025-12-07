package main

import (
	"os"
	"strings"
	"testing"
)

func readSampleInput(t *testing.T) string {
	t.Helper()
	buff, err := os.ReadFile("sample.txt")
	if err != nil {
		t.Fatalf("Error reading sample.txt: %v", err)
	}
	return strings.TrimSpace(string(buff))
}

func TestPart1(t *testing.T) {
	input := readSampleInput(t)
	w := parseInput(input)
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
	input := readSampleInput(t)
	w := parseInput(input)
	startX := findStart(w)
	if startX == -1 {
		t.Fatal("Could not find start")
	}
	expected := 40
	if got := solvePart2(w, startX); got != expected {
		t.Errorf("Part2 = %d, want %d", got, expected)
	}
}
