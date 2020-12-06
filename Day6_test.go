package main

import "testing"

func TestComputeDay6a(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Day 6 part 1 - example", args{"abc\n\na\nb\nc\n\nab\nac\n\na\na\na\na\n\nb"}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay6a(tt.args.input); got != tt.want {
				t.Errorf("ComputeDay6a() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputeDay6b(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Day 6 part 2 - example", args{"abc\n\na\nb\nc\n\nab\nac\n\na\na\na\na\n\nb"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay6b(tt.args.input); got != tt.want {
				t.Errorf("ComputeDay6b() = %v, want %v", got, tt.want)
			}
		})
	}
}