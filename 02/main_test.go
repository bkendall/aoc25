package main

import "testing"

func TestHasRepeatingHalves(t *testing.T) {
	tests := []struct {
		input int
		want  bool
	}{
		{123123, true},
		{1212, true},
		{1234, false},
		{121, false},
		{11, true},
		{1, false},
		{1010, true},
		{99, true},
		{123456, false},
	}

	for _, tt := range tests {
		if got := hasRepeatingHalves(tt.input); got != tt.want {
			t.Errorf("hasRepeatingHalves(%d) = %v; want %v", tt.input, got, tt.want)
		}
	}
}

func TestHasRepeatingPattern(t *testing.T) {
	tests := []struct {
		input int
		want  bool
	}{
		{123123, true},
		{1212, true},
		{121212, true},
		{1234, false},
		{121, false},
		{111, true},
		{999999, true},
		{123123123, true},
		{123124, false},
		{1, false},  // pattern length must be at least 1, but maxLen is len/2 = 0, loop doesn't run -> false
		{12, false}, // maxLen 1. pattern "1", check "2" != "1" -> false
		{101010, true},
	}

	for _, tt := range tests {
		if got := hasRepeatingPattern(tt.input); got != tt.want {
			t.Errorf("hasRepeatingPattern(%d) = %v; want %v", tt.input, got, tt.want)
		}
	}
}
