package main

import (
	"aoc25/lib"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Range struct {
	Min, Max int
}

func main() {
	sourcePtr := flag.String("source", "05/sample.txt", "input file")
	flag.Parse()

	buff, err := os.ReadFile(*sourcePtr)
	if err != nil {
		lib.Fatalf("Error: %v", err)
	}
	input := string(buff)
	input = strings.TrimSpace(input)

	ranges, ingredients := parseInput(input)

	// Part 1: Check if an ingrediant ID is in one of the ranges.
	freshCount := 0
	for _, id := range ingredients {
		for _, r := range ranges {
			if id >= r.Min && id <= r.Max {
				freshCount++
				break
			}
		}
	}
	fmt.Printf("Part 1: Fresh ingredients: %d\n", freshCount)

	// Part 2: Count the number of IDs that are included in all the ranges.
	totalFresh := solvePart2(ranges)
	fmt.Printf("Part 2: Total fresh ingredient IDs: %d\n", totalFresh)
}

func parseInput(input string) ([]Range, []int) {
	ranges := []Range{}
	ingredients := []int{}
	partOne := true
	for _, l := range strings.Split(input, "\n") {
		l = strings.TrimSpace(l)
		if l == "" {
			partOne = false
			continue
		}
		if partOne {
			var r Range
			arr := strings.Split(l, "-")
			r.Min, r.Max = lib.ToInt(arr[0]), lib.ToInt(arr[1])
			ranges = append(ranges, r)
		} else {
			ingredients = append(ingredients, lib.ToInt(l))
		}
	}
	return ranges, ingredients
}

func solvePart2(ranges []Range) int {
	if len(ranges) == 0 {
		return 0
	}

	// Make a copy to avoid modifying the original slice if it's used elsewhere
	sortedRanges := make([]Range, len(ranges))
	copy(sortedRanges, ranges)

	sort.Slice(sortedRanges, func(i, j int) bool {
		return sortedRanges[i].Min < sortedRanges[j].Min
	})

	total := 0
	current := sortedRanges[0]

	for _, r := range sortedRanges[1:] {
		if r.Min <= current.Max {
			// Overlap, extend current max if needed
			if r.Max > current.Max {
				current.Max = r.Max
			}
		} else {
			// No overlap, add current range length
			total += current.Max - current.Min + 1
			current = r
		}
	}
	// Add the last range
	total += current.Max - current.Min + 1

	return total
}