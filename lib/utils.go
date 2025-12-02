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
