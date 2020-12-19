package main

import "testing"

func TestComputeDay19a(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Day 19 part 1 - example 1", args{"0: 4 1 5\n1: 2 3 | 3 2\n2: 4 4 | 5 5\n3: 4 5 | 5 4\n4: \"a\"\n5: \"b\"\n\nababbb\nbababa\nabbbab\naaabbb\naaaabbb"}, 2},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay19a(tt.args.input); got != tt.want {
				t.Errorf("ComputeDay19a() = %v, want %v", got, tt.want)
			}
		})
	}
}
