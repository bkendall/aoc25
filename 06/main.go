package main

import (
	"aoc25/lib"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Worksheet struct {
	Numbers   [][]int
	Operation []string
}

func main() {
	sourcePtr := flag.String("source", "06/sample.txt", "input file")
	flag.Parse()

	buff, err := os.ReadFile(*sourcePtr)
	if err != nil {
		lib.Fatalf("Error: %v", err)
	}
	input := string(buff)
	input = strings.TrimSpace(input)

	ws := parseInput(input)

	// Part 1
	_, sum := doOperations(ws)
	fmt.Printf("Sum of results: %d\n", sum)

	world := map[int]map[int]string{}
	// Part 2. Little harder. Re-parse.
	maxX := 0
	maxY := len(strings.Split(input, "\n"))
	for y, line := range strings.Split(input, "\n") {
		// line = strings.TrimSpace(line)
		for x := range line {
			if _, ok := world[x]; !ok {
				world[x] = map[int]string{}
			}
			world[x][y] = string(line[x])
			if x > maxX {
				maxX = x
			}
		}
	}

	var nums []int
	var results2 []int
	for x := maxX; x >= 0; x-- {
		n := ""
		operation := ""
		for y := 0; y < maxY; y++ {
			c := world[x][y]
			if c == "*" || c == "+" {
				operation = c
				break
			}
			n += c
		}
		n = strings.TrimSpace(n)
		if n != "" {
			nums = append(nums, lib.ToInt(n))
		}
		if operation != "" {
			switch operation {
			case "+":
				sum := 0
				for _, v := range nums {
					sum += v
				}
				results2 = append(results2, sum)
			case "*":
				prod := 1
				for _, v := range nums {
					prod *= v
				}
				results2 = append(results2, prod)
			default:
				lib.Fatalf("Unknown operation: %s", operation)
			}
			nums = []int{}
		}
	}
	finalSum := 0
	for _, v := range results2 {
		finalSum += v
	}
	fmt.Printf("Sum of results (part 2): %d\n", finalSum)
}

func parseInput(input string) Worksheet {
	var ws Worksheet
	// This is odd because the worksheet is organized vertically.
	// So, every number in a row belongs do a different set of numbers.
	// The last line is not numbers, but operations.
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		line = strings.TrimSpace(line)
		parts := strings.Fields(line)
		if i == len(lines)-1 {
			// Last line: operations
			ws.Operation = parts
			continue
		}
		// Numbers
		if i == 0 {
			ws.Numbers = make([][]int, len(parts))
		}
		for j, part := range parts {
			num := lib.ToInt(part)
			ws.Numbers[j] = append(ws.Numbers[j], num)
		}
	}
	return ws
}

func doOperations(ws Worksheet) ([]int, int) {
	results := make([]int, len(ws.Numbers))
	for i, nums := range ws.Numbers {
		result := 0
		switch ws.Operation[i] {
		case "+":
			for _, n := range nums {
				result += n
			}
		case "*":
			result = 1
			for _, n := range nums {
				result *= n
			}
		default:
			lib.Fatalf("Unknown operation: %s", ws.Operation)
		}
		results[i] = result
	}

	sum := 0
	for _, res := range results {
		sum += res
	}

	return results, sum
}

func rToLNumbers([]int) []int {
	return nil
}
