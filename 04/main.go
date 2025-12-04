package main

import (
	"aoc25/lib"
	"flag"
	"fmt"
	"os"
	"strings"
)

type World map[int]map[int]bool

func countNeighbors(world World, x, y int) int {
	directions := [][2]int{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}
	neighbors := 0
	for _, d := range directions {
		dx, dy := d[0], d[1]
		if world[x+dx][y+dy] {
			neighbors++
		}
	}
	return neighbors
}

func main() {
	sourcePtr := flag.String("source", "04/sample.txt", "input file")
	flag.Parse()

	buff, err := os.ReadFile(*sourcePtr)
	if err != nil {
		lib.Fatalf("Error: %v", err)
	}
	input := string(buff)
	input = strings.TrimSpace(input)

	world := World{}
	for y, l := range strings.Split(input, "\n") {
		// For each line, record if there is a `@` in each spot.
		// The top right corner is 0, 0.
		for x, r := range l {
			if r == '@' {
				if _, ok := world[x]; !ok {
					world[x] = map[int]bool{}
				}
				world[x][y] = true
			}
		}
	}

	// Part 1.
	// For every position in the world where there is a `@`, count how many
	// are around it in all 8 directions.
	// If there are less than 4, add it to the count.
	count := 0
	for x, col := range world {
		for y := range col {
			if countNeighbors(world, x, y) < 4 {
				count++
			}
		}
	}
	fmt.Printf("Part 1: %d\n", count)

	// Part 2.
	// Like before, we're going to see what neighbors each have less than 4 neighbors.
	// But this time, we will remove them from the world, and repeat until no more can be removed.
	removed := 0
	for {
		toRemove := [][2]int{}
		for x, col := range world {
			for y := range col {
				if countNeighbors(world, x, y) < 4 {
					toRemove = append(toRemove, [2]int{x, y})
				}
			}
		}
		if len(toRemove) == 0 {
			break
		}
		for _, pos := range toRemove {
			x, y := pos[0], pos[1]
			delete(world[x], y)
			if len(world[x]) == 0 {
				delete(world, x)
			}
		}
		removed += len(toRemove)
	}
	fmt.Printf("Part 2: %d\n", removed)
}
