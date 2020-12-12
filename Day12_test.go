package main

import "testing"

func TestComputeDay12a(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{ "Day 12 - Part 1 - example", args {"F10\nN3\nF7\nR90\nF11"}, 25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay12a(tt.args.input); got != tt.want {
				t.Errorf("ComputeDay12a() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputeDay12b(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{ "Day 12 - Part 2 - example", args {"F10\nN3\nF7\nR90\nF11"}, 286},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay12b(tt.args.input); got != tt.want {
				t.Errorf("ComputeDay12b() = %v, want %v", got, tt.want)
			}
		})
	}
}