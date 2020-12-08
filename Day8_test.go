package main

import "testing"

func TestComputeDay8a(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Day8 - Sample - Part 1", args{"nop +0\nacc +1\njmp +4\nacc +3\njmp -3\nacc -99\nacc +1\njmp -4\nacc +6"}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay8a(tt.args.input); got != tt.want {
				t.Errorf("ComputeDay8a() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputeDay8b(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Day8 - Sample - Part 2", args{"nop +0\nacc +1\njmp +4\nacc +3\njmp -3\nacc -99\nacc +1\njmp -4\nacc +6"}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay8b(tt.args.input); got != tt.want {
				t.Errorf("ComputeDay8a() = %v, want %v", got, tt.want)
			}
		})
	}
}
