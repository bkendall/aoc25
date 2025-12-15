package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	R, C int
}

type Variation struct {
	Cells  []Point // Anchored at (0,0), which is the first cell in row-major order
	Width  int     // Effective width (max C + 1)
	Height int     // Effective height (max R + 1)
}

type Shape struct {
	ID         int
	Original   []string
	Variations []Variation
}

type Region struct {
	Width, Height int
	Counts        []int // Count for each Shape ID
}

var shapes []Shape
var regions []Region
var debug bool

func main() {
	sourcePtr := flag.String("source", "12/sample.txt", "input file")
	debugPtr := flag.Bool("debug", false, "debug output")
	flag.Parse()
	debug = *debugPtr

	buff, err := os.ReadFile(*sourcePtr)
	if err != nil {
		panic(err)
	}
	input := string(buff)
	input = strings.TrimSpace(input)

	parseInput(input)

	solve()
}

func parseInput(input string) {
	// Split into sections
	// The first section is shapes. The second is regions.
	// They are separated by... looking at sample, there isn't a clear delimiter except format.
	// We can iterate line by line.

	lines := strings.Split(input, "\n")
	var shapeLines []string
	var regionLines []string

	inShape := true
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}
		// Shapes start with "N:"
		// Regions start with "WxH:"
		if strings.Contains(line, ":") {
			parts := strings.Split(line, ":")
			firstPart := parts[0]
			if strings.Contains(firstPart, "x") {
				inShape = false
			}
		}

		if inShape {
			shapeLines = append(shapeLines, line)
		} else {
			regionLines = append(regionLines, line)
		}
	}

	parseShapes(shapeLines)
	parseRegions(regionLines)
}

func parseShapes(lines []string) {
	// Group lines by shape
	// Format: "ID:" then shape lines
	var currentID int = -1
	var currentBuffer []string

	processBuffer := func() {
		if currentID != -1 && len(currentBuffer) > 0 {
			shape := Shape{
				ID:       currentID,
				Original: make([]string, len(currentBuffer)),
			}
			copy(shape.Original, currentBuffer)
			shape.Variations = generateVariations(currentBuffer)
			shapes = append(shapes, shape)
		}
	}

	for _, line := range lines {
		if strings.HasSuffix(line, ":") {
			processBuffer()
			idStr := strings.TrimSuffix(line, ":")
			var err error
			currentID, err = strconvAtoi(idStr)
			if err != nil {
				panic(fmt.Sprintf("Invalid shape ID: %s", idStr))
			}
			currentBuffer = []string{}
		} else {
			currentBuffer = append(currentBuffer, line)
		}
	}
	processBuffer()

	// Ensure shapes are sorted by ID just in case
	sort.Slice(shapes, func(i, j int) bool {
		return shapes[i].ID < shapes[j].ID
	})
}

func generateVariations(lines []string) []Variation {
	// Convert lines to set of points
	var points []Point
	for r, line := range lines {
		for c, char := range line {
			if char == '#' {
				points = append(points, Point{r, c})
			}
		}
	}

	uniqueVars := make(map[string]Variation)

	// 8 symmetries
	// 0: Original
	// 1: Rot 90
	// 2: Rot 180
	// 3: Rot 270
	// 4: Flip H
	// 5: Flip H + Rot 90
	// 6: Flip H + Rot 180
	// 7: Flip H + Rot 270

	// Helper to rotate point 90 deg clockwise
	rot90 := func(p Point) Point {
		return Point{p.C, -p.R}
	}

	// Helper to flip horizontal
	flipH := func(p Point) Point {
		return Point{p.R, -p.C}
	}

	var current []Point

	// Base
	current = make([]Point, len(points))
	copy(current, points)

	for i := 0; i < 8; i++ {
		// Normalize and add
		v := normalize(current)
		key := fmt.Sprintf("%v", v.Cells)
		uniqueVars[key] = v

		if i == 3 {
			// Reset to original and flip
			current = make([]Point, len(points))
			copy(current, points)
			for j := range current {
				current[j] = flipH(current[j])
			}
		} else {
			// Rotate
			for j := range current {
				current[j] = rot90(current[j])
			}
		}
	}

	res := make([]Variation, 0, len(uniqueVars))
	for _, v := range uniqueVars {
		res = append(res, v)
	}
	return res
}

func normalize(points []Point) Variation {
	if len(points) == 0 {
		return Variation{}
	}

	// Sort by R then C to find the "first" cell
	// We want the reading-order first cell to be at (0,0)
	sorted := make([]Point, len(points))
	copy(sorted, points)
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].R != sorted[j].R {
			return sorted[i].R < sorted[j].R
		}
		return sorted[i].C < sorted[j].C
	})

	anchor := sorted[0]

	// Shift all points
	shifted := make([]Point, len(points))
	maxR, maxC := 0, 0
	for i, p := range sorted {
		nr, nc := p.R-anchor.R, p.C-anchor.C
		shifted[i] = Point{nr, nc}
		if nr > maxR {
			maxR = nr
		}
		if nc > maxC {
			maxC = nc
		} // Note: nc can be negative? No, because we sorted.
		// Wait. If anchor is (0,0) in sorted list.
		// A point could be (1, -1) relative to anchor?
		// Example: .#
		//          ##
		// Cells: (0,1), (1,0), (1,1).
		// Sorted: (0,1), (1,0), (1,1). Anchor (0,1).
		// Shifted: (0,0), (1,-1), (1,0).
		// (1,-1) means valid col index is -1.
		// BUT we established we only place anchor at current cursor (r,c).
		// So (r+1, c-1).
		// Is this valid? Yes.
		// But in our logic "previous cells are filled".
		// (r+1, c-1) is in the next row, so it is NOT processed yet.
		// However, does maxC calculate width correctly?
		// We might need minC too.
	}

	minC := 0
	for _, p := range shifted {
		if p.C < minC {
			minC = p.C
		}
	}

	// We do NOT shift to make minC 0. We keep anchor at (0,0).
	// But for bounds checking we need to know the extent.

	return Variation{
		Cells:  shifted,
		Width:  0, // Not strictly "width", but handled dynamically
		Height: 0,
	}
}

func parseRegions(lines []string) {
	for _, line := range lines {
		// "4x4: 0 0 0 0 2 0"
		parts := strings.Split(line, ":")
		dims := strings.Split(parts[0], "x")
		w, _ := strconvAtoi(dims[0])
		h, _ := strconvAtoi(dims[1])

		countsStr := strings.Fields(parts[1])
		counts := make([]int, len(countsStr))
		for i, s := range countsStr {
			c, _ := strconvAtoi(s)
			counts[i] = c
		}

		regions = append(regions, Region{w, h, counts})
	}
}

func strconvAtoi(s string) (int, error) {
	return strconv.Atoi(strings.TrimSpace(s))
}

func solve() {
	count := 0
	for i, region := range regions {
		if solveRegion(region) {
			if debug {
				fmt.Printf("Region %d: SUCCESS\n", i)
			}
			count++
		} else {
			if debug {
				fmt.Printf("Region %d: FAIL\n", i)
			}
		}
	}
	fmt.Printf("%d\n", count)
}

func solveRegion(region Region) bool {
	// Calculate areas
	presentArea := 0
	for id, count := range region.Counts {
		if id < len(shapes) {
			presentArea += count * len(shapes[id].Variations[0].Cells)
		}
	}

	totalArea := region.Width * region.Height
	if presentArea > totalArea {
		if debug {
			fmt.Printf("Area fail: %d > %d\n", presentArea, totalArea)
		}
		return false
	}

	slack := totalArea - presentArea
	if debug {
		fmt.Printf("Solving %dx%d, Slack: %d\n", region.Width, region.Height, slack)
	}

	grid := make([]uint64, region.Height)
	counts := make([]int, len(region.Counts))
	copy(counts, region.Counts)

	// Create map of ID -> variations to avoid lookup in loop
	// Actually shapes global is fine

	return backtrack(0, 0, grid, counts, slack, region.Width, region.Height)
}

func backtrack(r, c int, grid []uint64, counts []int, slack int, W, H int) bool {
	// Find next empty
	for r < H {
		if c >= W {
			c = 0
			r++
			continue
		}

		// Check bit
		if (grid[r] & (1 << c)) != 0 {
			c++
			continue
		}

		// Found empty
		break
	}

	if r == H {
		// All filled (or skipped). Check if we used all presents.
		for _, count := range counts {
			if count > 0 {
				return false
			}
		}
		return true
	}

	// Option 1: Place a present
	// Optimization: Sort available presents?
	// For now, iterate
	for id, count := range counts {
		if count > 0 {
			// Try this shape
			counts[id]--

			for _, v := range shapes[id].Variations {
				if canPlace(grid, v, r, c, W, H) {
					place(grid, v, r, c)
					if backtrack(r, c+1, grid, counts, slack, W, H) { // Move to next cell? Or just recurse (loop will find next)
						// Optimization: passing c+1 helps
						return true
					}
					unplace(grid, v, r, c)
				}
			}

			counts[id]++
		}
	}

	// Option 2: Skip (use slack)
	if slack > 0 {
		grid[r] |= (1 << c)
		if backtrack(r, c+1, grid, counts, slack-1, W, H) {
			return true
		}
		grid[r] &= ^(1 << c)
	}

	return false
}

func canPlace(grid []uint64, v Variation, r, c, W, H int) bool {
	for _, p := range v.Cells {
		nr, nc := r+p.R, c+p.C
		if nr < 0 || nr >= H || nc < 0 || nc >= W {
			return false
		}
		if (grid[nr] & (1 << nc)) != 0 {
			return false
		}
	}
	return true
}

func place(grid []uint64, v Variation, r, c int) {
	for _, p := range v.Cells {
		nr, nc := r+p.R, c+p.C
		grid[nr] |= (1 << nc)
	}
}

func unplace(grid []uint64, v Variation, r, c int) {
	for _, p := range v.Cells {
		nr, nc := r+p.R, c+p.C
		grid[nr] &= ^(1 << nc)
	}
}
