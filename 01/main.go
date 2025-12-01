package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func fatalf(format string, args ...any) {
	fmt.Printf(format, args...)
	os.Exit(1)
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fatalf("%v", err)
	}
	return i
}

func abs(i int) int {
	return int(math.Abs(float64(i)))
}

type Turn struct {
	direction string
	steps     int
}

func main() {
	sourcePtr := flag.String("source", "./sample.txt", "input file")
	flag.Parse()

	buff, err := os.ReadFile(*sourcePtr)
	if err != nil {
		fatalf("Error: %v", err)
	}
	input := string(buff)
	input = strings.TrimSpace(input)

	turns := []Turn{}
	for _, l := range strings.Split(input, "\n") {
		t := string(l[0])
		s := toInt(l[1:])
		turns = append(turns, Turn{direction: t, steps: s})
	}

	// Part 1
	pointer := 50
	count := 0
	for _, t := range turns {
		switch t.direction {
		case "R":
			pointer += t.steps
		case "L":
			pointer -= t.steps
		}
		pointer = pointer % 100

		if pointer == 0 {
			count++
		}
	}

	fmt.Printf("Count: %d\n", count)

	// Part 2
	pointer = 50
	count = 0
	for _, t := range turns {
		// This is bad, but it works.
		// I got stuck always counting when we went negative as "we passed 0".
		steps := t.steps
		for range steps {
			switch t.direction {
			case "R":
				pointer++
			case "L":
				pointer--
			}
			if pointer == 100 {
				pointer = 0
			}
			if pointer == -1 {
				pointer = 99
			}
			if pointer == 0 {
				count++
			}
		}
	}

	fmt.Printf("Count 2: %d\n", count)
}
