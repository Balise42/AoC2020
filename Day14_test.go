package main

import "testing"

func TestComputeDay14a(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{ "Day 14 Part 1 example", args {"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X\nmem[8] = 11\nmem[7] = 101\nmem[8] = 0"}, 165},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay14a(tt.args.input); got != tt.want {
				t.Errorf("ComputeDay14a() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputeDay14b(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{ "Day 14 Part 2 example", args {"mask = 000000000000000000000000000000X1001X\nmem[42] = 100\nmask = 00000000000000000000000000000000X0XX\nmem[26] = 1"}, 208},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay14b(tt.args.input); got != tt.want {
				t.Errorf("ComputeDay14b() = %v, want %v", got, tt.want)
			}
		})
	}
}