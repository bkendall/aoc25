package main

import (
	"aoc25/lib"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Coordinate struct {
	X, Y int
}

func main() {
	sourcePtr := flag.String("source", "09/sample.txt", "input file")
	flag.Parse()

	buff, err := os.ReadFile(*sourcePtr)
	if err != nil {
		lib.Fatalf("Error: %v", err)
	}
	input := string(buff)
	input = strings.TrimSpace(input)

	coordinates := parseInput(input)

	size := findLargestArea(coordinates)
	fmt.Printf("Largest area (part 1): %d\n", size)

	size = findLargestAreaInBoundary(coordinates)
	fmt.Printf("Largest area (part 2): %d\n", size)
}

func parseInput(input string) []*Coordinate {
	ret := []*Coordinate{}
	for _, line := range strings.Split(input, "\n") {
		arr := strings.Split(line, ",")
		x, y := lib.ToInt(arr[0]), lib.ToInt(arr[1])
		c := &Coordinate{X: x, Y: y}
		ret = append(ret, c)
	}
	return ret
}

func findLargestArea(coordinates []*Coordinate) int {
	// Find the largest rectangle that can be made with any pair of coordinates.
	// The minimum width or hight is 1 (because we're counting tiles).
	// We can only create rectangles using a pair of coordinates from the given list.
	maxArea := 0
	for i := 0; i < len(coordinates); i++ {
		for j := i + 1; j < len(coordinates); j++ {
			c1 := coordinates[i]
			c2 := coordinates[j]
			width := lib.Abs(c1.X-c2.X) + 1
			height := lib.Abs(c1.Y-c2.Y) + 1
			area := width * height
			// fmt.Printf("Coordinates: %+v and %+v => width: %d, height: %d, area: %d\n", c1, c2, width, height, area)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

// Note: Okay, this went really hard on the LLM solution. Interesting to see...

func findLargestAreaInBoundary(coordinates []*Coordinate) int {
	// Collect unique X and Y coordinates to compress the grid.
	uniqueX := make(map[int]bool)
	uniqueY := make(map[int]bool)
	for _, c := range coordinates {
		uniqueX[c.X] = true
		uniqueY[c.Y] = true
	}
	var sortedX, sortedY []int
	for x := range uniqueX {
		sortedX = append(sortedX, x)
	}
	for y := range uniqueY {
		sortedY = append(sortedY, y)
	}
	sort.Ints(sortedX)
	sort.Ints(sortedY)

	xMap := make(map[int]int)
	for i, x := range sortedX {
		xMap[x] = i
	}
	yMap := make(map[int]int)
	for i, y := range sortedY {
		yMap[y] = i
	}

	// Create a compressed grid.
	// We use 2*N+1 size to represent intervals and points.
	// Index 2*i+1 corresponds to sortedX[i].
	// Index 2*i corresponds to the interval (sortedX[i-1], sortedX[i]).
	W := 2*len(sortedX) + 1
	H := 2*len(sortedY) + 1
	grid := make([][]bool, W)
	for i := range grid {
		grid[i] = make([]bool, H)
	}

	// Draw boundary on the compressed grid.
	for i := 0; i < len(coordinates); i++ {
		c1 := coordinates[i]
		c2 := coordinates[0]
		if i < len(coordinates)-1 {
			c2 = coordinates[i+1]
		}

		x1, y1 := 2*xMap[c1.X]+1, 2*yMap[c1.Y]+1
		x2, y2 := 2*xMap[c2.X]+1, 2*yMap[c2.Y]+1

		if x1 == x2 { // Vertical segment
			minY, maxY := y1, y2
			if minY > maxY {
				minY, maxY = maxY, minY
			}
			for y := minY; y <= maxY; y++ {
				grid[x1][y] = true
			}
		} else { // Horizontal segment
			minX, maxX := x1, x2
			if minX > maxX {
				minX, maxX = maxX, minX
			}
			for x := minX; x <= maxX; x++ {
				grid[x][y1] = true
			}
		}
	}

	// Flood fill from outside (0,0) to identify outside regions.
	// 0,0 is guaranteed to be outside because indices start at 1 for coordinates.
	queue := []struct{ x, y int }{{0, 0}}
	visited := make([][]bool, W)
	for i := range visited {
		visited[i] = make([]bool, H)
	}
	visited[0][0] = true

	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]

		dirs := []struct{ dx, dy int }{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for _, d := range dirs {
			nx, ny := c.x+d.dx, c.y+d.dy
			if nx >= 0 && nx < W && ny >= 0 && ny < H {
				if !visited[nx][ny] && !grid[nx][ny] {
					visited[nx][ny] = true
					queue = append(queue, struct{ x, y int }{nx, ny})
				}
			}
		}
	}

	// Determine valid cells (Inside or Boundary).
	isValid := make([][]int, W)
	for i := range isValid {
		isValid[i] = make([]int, H)
		for j := range isValid[i] {
			if !visited[i][j] { // Not visited from outside -> Inside or Boundary
				isValid[i][j] = 1
			}
		}
	}

	// Build 2D Prefix Sum array.
	prefixSum := make([][]int, W)
	for i := 0; i < W; i++ {
		prefixSum[i] = make([]int, H)
		for j := 0; j < H; j++ {
			val := isValid[i][j]
			left := 0
			if i > 0 {
				left = prefixSum[i-1][j]
			}
			top := 0
			if j > 0 {
				top = prefixSum[i][j-1]
			}
			diag := 0
			if i > 0 && j > 0 {
				diag = prefixSum[i-1][j-1]
			}
			prefixSum[i][j] = val + left + top - diag
		}
	}

	getSum := func(x1, y1, x2, y2 int) int {
		total := prefixSum[x2][y2]
		left := 0
		if x1 > 0 {
			left = prefixSum[x1-1][y2]
		}
		top := 0
		if y1 > 0 {
			top = prefixSum[x2][y1-1]
		}
		diag := 0
		if x1 > 0 && y1 > 0 {
			diag = prefixSum[x1-1][y1-1]
		}
		return total - left - top + diag
	}

	maxArea := 0
	for i := 0; i < len(coordinates); i++ {
		for j := i + 1; j < len(coordinates); j++ {
			c1 := coordinates[i]
			c2 := coordinates[j]

			// Map to compressed grid coordinates
			u1, v1 := 2*xMap[c1.X]+1, 2*yMap[c1.Y]+1
			u2, v2 := 2*xMap[c2.X]+1, 2*yMap[c2.Y]+1

			minU, maxU := u1, u2
			if minU > maxU {
				minU, maxU = maxU, minU
			}
			minV, maxV := v1, v2
			if minV > maxV {
				minV, maxV = maxV, minV
			}

			// Calculate expected count of valid cells in the rectangle
			widthCells := maxU - minU + 1
			heightCells := maxV - minV + 1
			expectedSum := widthCells * heightCells

			actualSum := getSum(minU, minV, maxU, maxV)

			if actualSum == expectedSum {
				realWidth := lib.Abs(c1.X-c2.X) + 1
				realHeight := lib.Abs(c1.Y-c2.Y) + 1
				area := realWidth * realHeight
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}
