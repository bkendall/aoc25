package main

import (
	"aoc25/lib"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Range struct {
	start int
	end   int
}

func main() {
	sourcePtr := flag.String("source", "02/sample.txt", "input file")
	flag.Parse()

	buff, err := os.ReadFile(*sourcePtr)
	if err != nil {
		lib.Fatalf("Error: %v", err)
	}
	input := string(buff)
	input = strings.TrimSpace(input)

	ranges := []Range{}
	for _, l := range strings.Split(input, "\n") {
		for _, r := range strings.Split(l, ",") {
			parts := strings.Split(r, "-")
			start, end := lib.ToInt(parts[0]), lib.ToInt(parts[1])
			ranges = append(ranges, Range{start: start, end: end})
		}
	}

	// Part 1
	// See if the string representing each number in the range
	// has repeated character patterns.
	count := 0
	sum := 0
	for _, r := range ranges {
		start, end := r.start, r.end
		for i := start; i <= end; i++ {
			s := fmt.Sprintf("%d", i)
			if len(s)%2 == 1 {
				// Odd length cannot have repeated pattern.
				continue
			}
			mid := len(s) / 2
			first, second := s[0:mid], s[mid:]
			if first == second {
				count++
				sum += i
			}
		}
	}
	fmt.Printf("Count: %d\n", count)
	fmt.Printf("Sum: %d\n", sum)

	// Part 2
	// Now we're looking for any repeating pattern in the string.
	count = 0
	sum = 0
	for _, r := range ranges {
		start, end := r.start, r.end
		for i := start; i <= end; i++ {
			s := fmt.Sprintf("%d", i)
			maxLen := len(s) / 2
			for l := 1; l <= maxLen; l++ {
				if len(s)%l != 0 {
					// Pattern length must divide string length.
					continue
				}
				pattern := s[0:l]
				// Check if the pattern repeats in the string.
				repeated := true
				for j := l; j < len(s); j += l {
					if s[j:j+l] != pattern {
						repeated = false
						break
					}
				}
				if repeated {
					count++
					sum += i
					break
				}
			}
		}
	}
	fmt.Printf("Count: %d\n", count)
	fmt.Printf("Sum: %d\n", sum)
}
