package main

import (
	"aoc25/lib"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Bank []int

func main() {
	sourcePtr := flag.String("source", "03/sample.txt", "input file")
	flag.Parse()

	buff, err := os.ReadFile(*sourcePtr)
	if err != nil {
		lib.Fatalf("Error: %v", err)
	}
	input := string(buff)
	input = strings.TrimSpace(input)

	banks := []Bank{}
	for _, l := range strings.Split(input, "\n") {
		// Parse each line into a Bank (slice of ints)
		bank := Bank{}
		for _, s := range strings.Split(l, "") {
			bank = append(bank, lib.ToInt(s))
		}
		banks = append(banks, bank)
	}

	// Part 1.
	// For each bank, find the largest 2-digit number that
	// can be formed.
	sum := 0
	for _, bank := range banks {
		maxNum := -1
		for i := 0; i < len(bank)-1; i++ {
			for j := i + 1; j < len(bank); j++ {
				num := joinDigits(bank[i], bank[j])
				if num > maxNum {
					maxNum = num
				}
			}
		}
		// fmt.Printf("Bank: %v, Max 2-digit number: %d\n", bank, maxNum)
		sum += maxNum
	}
	fmt.Printf("Part 1: Sum of largest 2-digit numbers: %d\n", sum)

	// Part 2.
	// Now we try the same thing, but we find the largest 12-digit number.
	sum = 0
	for _, bank := range banks {
		// Greedy approach to find the largest 12-digit subsequence.
		maxNum := -1
		if len(bank) >= 12 {
			maxNum = 0
			currentStart := 0
			for i := 0; i < 12; i++ {
				// We need to pick the (i+1)-th digit.
				// We must leave at least (11-i) digits after it.
				// So the index j can go up to len(bank) - (12-i).
				limit := len(bank) - (12 - i)
				bestDigit := -1
				bestIdx := -1

				for j := currentStart; j <= limit; j++ {
					if bank[j] > bestDigit {
						bestDigit = bank[j]
						bestIdx = j
						if bestDigit == 9 {
							break
						}
					}
				}
				maxNum = joinDigits(maxNum, bestDigit)
				currentStart = bestIdx + 1
			}
		}
		// fmt.Printf("Bank: %v, Max 12-digit number: %d\n", bank, maxNum)
		sum += maxNum
	}
	fmt.Printf("Part 2: Sum of largest 12-digit numbers: %d\n", sum)
}

// joinDigits takes to integers and returns the integer represented by joining their digits.
func joinDigits(a, b int) int {
	return a*10 + b
}
