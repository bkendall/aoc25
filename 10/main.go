package main

import (
	"aoc25/lib"
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
)

type Machine struct {
	NumLights      int
	LightStates    []bool
	Buttons        int
	ButtonTriggers map[int][]int
	TargetState    []bool
	TargetJoltage  []int
}

func (m Machine) String() string {
	s := "["
	for _, b := range m.LightStates {
		switch b {
		case true:
			s += "#"
		case false:
			s += "."
		}
	}
	s += "]"
	for i := 0; i < m.Buttons; i++ {
		s += " "
		s += fmt.Sprintf("%v", m.ButtonTriggers[i])
	}
	return s
}

func (m *Machine) pressButton(button int) {
	for _, l := range m.ButtonTriggers[button] {
		m.LightStates[l] = !m.LightStates[l]
	}
}

func (m *Machine) isTargetState() bool {
	for i := 0; i < m.NumLights; i++ {
		if m.LightStates[i] != m.TargetState[i] {
			return false
		}
	}
	return true
}

func main() {
	sourcePtr := flag.String("source", "10/sample.txt", "input file")
	flag.Parse()

	buff, err := os.ReadFile(*sourcePtr)
	if err != nil {
		lib.Fatalf("Error: %v", err)
	}
	input := string(buff)
	input = strings.TrimSpace(input)

	machines := parseInput(input)

	sum := solvePartOne(machines)
	fmt.Println("Part 1:", sum)

	sum = solvePartTwo(machines)
	fmt.Println("Part 2:", sum)
}

func parseInput(input string) []*Machine {
	var ret []*Machine
	// Sample: `[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}`
	for _, line := range strings.Split(input, "\n") {
		m := &Machine{}
		arr := strings.Split(line, " ")
		var t string
		t, arr = arr[0], arr[1:]
		t = strings.Trim(t, "[]")
		m.NumLights = len(t)
		// All the lights start off.
		m.LightStates = make([]bool, len(t))
		m.TargetState = make([]bool, len(t))
		for i, c := range t {
			if c == '#' {
				m.TargetState[i] = true
			}
		}

		last, arr := arr[len(arr)-1], arr[:len(arr)-1]
		last = strings.Trim(last, "{}")
		lastArr := strings.Split(last, ",")
		m.TargetJoltage = make([]int, len(lastArr))
		for i, j := range lastArr {
			m.TargetJoltage[i] = lib.ToInt(j)
		}

		// Each arr is a button that triggers a set of lights.
		m.Buttons = len(arr)
		m.ButtonTriggers = make(map[int][]int, m.Buttons)
		for i, a := range arr {
			a = strings.Trim(a, "()")
			var lights []int
			for _, l := range strings.Split(a, ",") {
				lights = append(lights, lib.ToInt(l))
			}
			m.ButtonTriggers[i] = lights
		}

		ret = append(ret, m)
	}
	return ret
}

// For each machine, find the minimum number of button presses required to reach the target state.
// This will take some finagling to get right, and likely some dynamic programming.
func solvePartOne(machines []*Machine) int {
	sum := 0
	for _, m := range machines {
		presses := findMinPresses(m)
		if presses == -1 {
			lib.Fatalf("Impossible configuration found")
		}
		sum += presses
	}
	return sum
}

func findMinPresses(m *Machine) int {
	rows := m.NumLights
	cols := m.Buttons

	// Augmented matrix: last column is target state
	// mat[i][j] = 1 if button j affects light i
	// mat[i][cols] = target state of light i
	mat := make([][]int, rows)
	for i := range mat {
		mat[i] = make([]int, cols+1)
		if m.TargetState[i] {
			mat[i][cols] = 1
		}
	}

	for b := 0; b < cols; b++ {
		for _, l := range m.ButtonTriggers[b] {
			if l < rows {
				mat[l][b] = 1
			}
		}
	}

	// Gaussian elimination to RREF
	pivotRow := 0
	pivots := make([]int, rows) // stores the column index of the pivot for each row, or -1
	for i := range pivots {
		pivots[i] = -1
	}

	colToPivotRow := make(map[int]int)

	for j := 0; j < cols && pivotRow < rows; j++ {
		// Find row with 1 in col j at or below pivotRow
		sel := -1
		for i := pivotRow; i < rows; i++ {
			if mat[i][j] == 1 {
				sel = i
				break
			}
		}

		if sel == -1 {
			continue
		}

		// Swap rows
		mat[pivotRow], mat[sel] = mat[sel], mat[pivotRow]

		// Eliminate other rows
		for i := 0; i < rows; i++ {
			if i != pivotRow && mat[i][j] == 1 {
				for k := j; k <= cols; k++ {
					mat[i][k] ^= mat[pivotRow][k]
				}
			}
		}

		pivots[pivotRow] = j
		colToPivotRow[j] = pivotRow
		pivotRow++
	}

	// Check for inconsistency
	for i := pivotRow; i < rows; i++ {
		if mat[i][cols] == 1 {
			return -1 // Impossible
		}
	}

	// Identify free variables
	var freeVars []int
	for j := 0; j < cols; j++ {
		if _, ok := colToPivotRow[j]; !ok {
			freeVars = append(freeVars, j)
		}
	}

	minPresses := -1

	// Iterate 2^len(freeVars)
	count := 1 << len(freeVars)
	for i := 0; i < count; i++ {
		presses := 0
		assignments := make([]int, cols)

		tempVal := i
		for _, fv := range freeVars {
			assignments[fv] = tempVal & 1
			if assignments[fv] == 1 {
				presses++
			}
			tempVal >>= 1
		}

		// Calculate dependent variables
		for r := 0; r < pivotRow; r++ {
			pVal := pivots[r]
			val := mat[r][cols]
			for c := pVal + 1; c < cols; c++ {
				if mat[r][c] == 1 {
					val ^= assignments[c]
				}
			}
			assignments[pVal] = val
			if val == 1 {
				presses++
			}
		}

		if minPresses == -1 || presses < minPresses {
			minPresses = presses
		}
	}

	return minPresses
}

// In this part, we're going to need to find the minimum number of button presses
// to reach the "TargetJoltage". The index in the TargetJoltage array cooresponds
// each light that the button affects. Pressing a button increments the "joltage"
// by one for each light it affects.
func solvePartTwo(machines []*Machine) int {
	sum := 0
	for _, m := range machines {
		presses := findMinPressesPartTwo(m)
		if presses == -1 {
			lib.Fatalf("Impossible configuration found in Part 2")
		}
		sum += presses
	}
	return sum
}

func findMinPressesPartTwo(m *Machine) int {
	rows := m.NumLights
	cols := m.Buttons

	// Build Augmented Matrix
	mat := make([][]float64, rows)
	for i := range mat {
		mat[i] = make([]float64, cols+1)
		mat[i][cols] = float64(m.TargetJoltage[i])
	}

	// Upper bounds for variables
	maxVal := make([]int, cols)
	for j := 0; j < cols; j++ {
		minB := -1
		hasConstraint := false
		for _, l := range m.ButtonTriggers[j] {
			if l < rows {
				mat[l][j] = 1.0
				if minB == -1 || m.TargetJoltage[l] < minB {
					minB = m.TargetJoltage[l]
				}
				hasConstraint = true
			}
		}
		if !hasConstraint {
			maxVal[j] = 0 // Button affects nothing useful
		} else {
			maxVal[j] = minB
		}
	}

	// Gaussian Elimination
	pivotRow := 0
	pivots := make([]int, rows) // pivots[row] = col
	for i := range pivots {
		pivots[i] = -1
	}

	colToPivotRow := make(map[int]int)

	for j := 0; j < cols && pivotRow < rows; j++ {
		// Find pivot
		sel := -1
		for i := pivotRow; i < rows; i++ {
			if math.Abs(mat[i][j]) > 1e-9 {
				sel = i
				break
			}
		}
		if sel == -1 {
			continue
		}

		// Swap
		mat[pivotRow], mat[sel] = mat[sel], mat[pivotRow]

		// Normalize pivot row
		pivotVal := mat[pivotRow][j]
		for k := j; k <= cols; k++ {
			mat[pivotRow][k] /= pivotVal
		}

		// Eliminate
		for i := 0; i < rows; i++ {
			if i != pivotRow {
				factor := mat[i][j]
				if math.Abs(factor) > 1e-9 {
					for k := j; k <= cols; k++ {
						mat[i][k] -= factor * mat[pivotRow][k]
					}
				}
			}
		}
		pivots[pivotRow] = j
		colToPivotRow[j] = pivotRow
		pivotRow++
	}

	// Check consistency of remaining rows
	for i := pivotRow; i < rows; i++ {
		if math.Abs(mat[i][cols]) > 1e-9 {
			return -1 // Impossible
		}
	}

	// Identify free vars
	var freeVars []int
	for j := 0; j < cols; j++ {
		if _, ok := colToPivotRow[j]; !ok {
			freeVars = append(freeVars, j)
		}
	}

	minPresses := -1

	// Recursive search
	var solve func(idx int, currentSum int, assignment []int)
	solve = func(idx int, currentSum int, assignment []int) {
		if minPresses != -1 && currentSum >= minPresses {
			return
		}

		if idx == len(freeVars) {
			// Calculate pivot vars
			thisSum := currentSum
			valid := true

			for r := 0; r < pivotRow; r++ {
				val := mat[r][cols]
				for _, fv := range freeVars {
					val -= mat[r][fv] * float64(assignment[fv])
				}

				// Check integer and non-negative
				if val < -1e-9 {
					valid = false
					break
				}

				intVal := int(math.Round(val))
				if math.Abs(val-float64(intVal)) > 1e-9 {
					valid = false
					break
				}
				thisSum += intVal
			}

			if valid {
				if minPresses == -1 || thisSum < minPresses {
					minPresses = thisSum
				}
			}
			return
		}

		fv := freeVars[idx]
		limit := maxVal[fv]
		for v := 0; v <= limit; v++ {
			assignment[fv] = v
			solve(idx+1, currentSum+v, assignment)
		}
	}

	assignment := make([]int, cols)
	solve(0, 0, assignment)

	return minPresses
}
