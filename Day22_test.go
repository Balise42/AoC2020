package main

import "testing"

func TestComputeDay22a(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Day 22 a example", args{"Player 1:\n9\n2\n6\n3\n1\n\nPlayer 2:\n5\n8\n4\n7\n10"}, 306},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay22a(tt.args.input); got != tt.want {
				t.Errorf("ComputeDay22a() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputeDay22b(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Day 22 b example", args{"Player 1:\n9\n2\n6\n3\n1\n\nPlayer 2:\n5\n8\n4\n7\n10"}, 291},
		{ "Day 22 b infinite loop detection", args {"Player 1:\n43\n19\n\nPlayer 2:\n2\n29\n14"}, 105},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay22b(tt.args.input); got != tt.want {
				t.Errorf("ComputeDay22b() = %v, want %v", got, tt.want)
			}
		})
	}
}