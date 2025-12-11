package lib

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func Fatalf(format string, args ...any) {
	fmt.Printf(format, args...)
	os.Exit(1)
}

func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		Fatalf("%v", err)
	}
	return i
}

func Abs(i int) int {
	return int(math.Abs(float64(i)))
}

func Sqrt(i int) int {
	return int(math.Sqrt(float64(i)))
}

func Square(i int) int {
	return i * i
}

func Distance3D(x1, y1, z1, x2, y2, z2 int) int {
	return Sqrt(Square(x2-x1) + Square(y2-y1) + Square(z2-z1))
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
