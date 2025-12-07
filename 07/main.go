package main

import (
	"aoc25/lib"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Entity string

const (
	Start    Entity = "S"
	Empty    Entity = "."
	Splitter Entity = "^"
	Beam     Entity = "|"
)

type World map[int]map[int]Entity

func main() {
	sourcePtr := flag.String("source", "07/sample.txt", "input file")
	flag.Parse()

	buff, err := os.ReadFile(*sourcePtr)
	if err != nil {
		lib.Fatalf("Error: %v", err)
	}
	input := string(buff)
	input = strings.TrimSpace(input)

	w := parseInput(input)
	startX := findStart(w)
	fmt.Printf("Starting at %d,0\n", startX)

	p1 := solvePart1(w, startX)
	fmt.Printf("Part 1: hit %d splitters\n", p1)

	p2 := solvePart2(w, startX)
	fmt.Printf("Part 2: total paths to bottom: %d\n", p2)
}

func parseInput(input string) World {
	w := make(World)
	for y, line := range strings.Split(input, "\n") {
		w[y] = make(map[int]Entity)
		for x, char := range line {
			w[y][x] = Entity(char)
		}
	}
	return w
}

func findStart(w World) int {
	for x := 0; x < len(w[0]); x++ {
		if w[0][x] == Start {
			return x
		}
	}
	return -1
}

func solvePart1(w World, startX int) int {
	beams := map[int]bool{startX: true}
	splitterCount := 0
	for y := 1; y < len(w); y++ {
		newBeams := map[int]bool{}
		for x := range beams {
			entity := w[y][x]
			switch entity {
			case Splitter:
				splitterCount++
				newBeams[x-1] = true
				newBeams[x+1] = true
			case Empty, Beam:
				newBeams[x] = true
			}
		}
		beams = newBeams
	}
	return splitterCount
}

func solvePart2(w World, startX int) int {
	paths := map[int]int{startX: 1}
	for y := 1; y < len(w); y++ {
		newPaths := map[int]int{}
		for x, count := range paths {
			entity := w[y][x]
			switch entity {
			case Splitter:
				newPaths[x-1] += count
				newPaths[x+1] += count
			case Empty, Beam:
				newPaths[x] += count
			}
		}
		paths = newPaths
	}
	totalPaths := 0
	for _, count := range paths {
		totalPaths += count
	}
	return totalPaths
}
