package main

import (
	"aoc25/lib"
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
)

type JunctionBox struct {
	ID        int
	X, Y, Z   int
	Distances map[int]int // distances to other junction boxes
}

func main() {
	sourcePtr := flag.String("source", "07/sample.txt", "input file")
	connections := flag.Int("connections", 10, "number of connections to form")
	flag.Parse()

	buff, err := os.ReadFile(*sourcePtr)
	if err != nil {
		lib.Fatalf("Error: %v", err)
	}
	input := string(buff)
	input = strings.TrimSpace(input)

	world := parseInput(input)
	distances := calculate(world)
	sortedDistances := sortedByDistance(distances)

	res := solvePartOne(sortedDistances, *connections)
	fmt.Printf("Part One: %d\n", res)

	dist := solvePartTwo(world, sortedDistances)
	fmt.Printf("Part Two: %d\n", dist)
}

func parseInput(input string) map[int]*JunctionBox {
	m := make(map[int]*JunctionBox)
	for id, line := range strings.Split(input, "\n") {
		jb := &JunctionBox{ID: id, Distances: make(map[int]int)}
		arr := strings.Split(line, ",")
		jb.X, jb.Y, jb.Z = lib.ToInt(arr[0]), lib.ToInt(arr[1]), lib.ToInt(arr[2])
		m[id] = jb
	}
	return m
}

func calculate(world map[int]*JunctionBox) map[int]map[int]int {
	distances := make(map[int]map[int]int)
	for id1, jb1 := range world {
		distances[id1] = make(map[int]int)
		for id2, jb2 := range world {
			if id1 == id2 {
				continue
			}
			dist := lib.Distance3D(jb1.X, jb1.Y, jb1.Z, jb2.X, jb2.Y, jb2.Z)
			distances[id1][id2] = dist
			jb1.Distances[id2] = dist
		}
	}
	return distances
}

type Pair struct {
	jb1ID int
	jb2ID int
	dist  int
}

func sortedByDistance(distances map[int]map[int]int) []Pair {
	allDistances := []Pair{}
	for id1, dists := range distances {
		for id2, dist := range dists {
			if id1 < id2 {
				allDistances = append(allDistances, Pair{
					jb1ID: id1,
					jb2ID: id2,
					dist:  dist,
				})
			}
		}
	}
	// And sort them...
	slices.SortFunc(allDistances, func(a, b Pair) int {
		return a.dist - b.dist
	})
	return allDistances
}

// We want to start making circuits by connecting the closesst junction
// boxes together. As they are joined, they may or may not join existing
// circuts. Once we make the required number of connections, we stop and
// return the product of the sizes of the circuits.
func solvePartOne(allDistances []Pair, connections int) int {
	// We can assume any junction box that isn't in a circuit is its own circuit.
	// Let's start connecting them up and tracking the circuits as we go.
	circuitMap := make(map[int]int) // junction box ID -> circuit ID
	circuitSizes := make(map[int]int)
	nextCircuitID := 0

	// Now, we're going to make the required number of connections starting
	// from the shortest distance.
	circuitNodes := make(map[int][]int)
	for i := 0; i < connections && i < len(allDistances); i++ {
		pair := allDistances[i]
		id1, id2 := pair.jb1ID, pair.jb2ID
		c1, ok1 := circuitMap[id1]
		c2, ok2 := circuitMap[id2]

		if !ok1 && !ok2 {
			newID := nextCircuitID
			nextCircuitID++
			circuitMap[id1] = newID
			circuitMap[id2] = newID
			circuitSizes[newID] = 2
			circuitNodes[newID] = []int{id1, id2}
		} else if ok1 && !ok2 {
			circuitMap[id2] = c1
			circuitSizes[c1]++
			circuitNodes[c1] = append(circuitNodes[c1], id2)
		} else if !ok1 && ok2 {
			circuitMap[id1] = c2
			circuitSizes[c2]++
			circuitNodes[c2] = append(circuitNodes[c2], id1)
		} else {
			if c1 == c2 {
				continue
			}
			// Merge c2 into c1
			for _, node := range circuitNodes[c2] {
				circuitMap[node] = c1
			}
			circuitNodes[c1] = append(circuitNodes[c1], circuitNodes[c2]...)
			circuitSizes[c1] += circuitSizes[c2]
			delete(circuitSizes, c2)
			delete(circuitNodes, c2)
		}
	}

	// Calculate the product of the sizes of the three largest circuits.
	sizes := []int{}
	for _, size := range circuitSizes {
		sizes = append(sizes, size)
	}
	slices.Sort(sizes)
	product := 1
	for i := 0; i < 3 && i < len(sizes); i++ {
		product *= sizes[len(sizes)-1-i]
	}
	return product
}

func solvePartTwo(world map[int]*JunctionBox, allDistances []Pair) int {
	// Connect junction boxes until all are connected. Save the coordinates of
	// the last two connected boxes. Return the product of their X coordinates.
	var lastJb1, lastJb2 *JunctionBox
	connected := make(map[int]bool)
	for i := 0; len(connected) < len(world) && i < len(allDistances); i++ {
		pair := allDistances[i]
		id1, id2 := pair.jb1ID, pair.jb2ID
		_, ok1 := connected[id1]
		_, ok2 := connected[id2]

		if !ok1 || !ok2 {
			connected[id1] = true
			connected[id2] = true
			lastJb1 = world[id1]
			lastJb2 = world[id2]
		}
	}

	if lastJb1 == nil || lastJb2 == nil {
		return 0
	}
	return lastJb1.X * lastJb2.X
}
