package main

import "testing"

func TestComputeDay23a(t *testing.T) {
	type args struct {
		s      string
		rounds int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{ "Day 23 part 1 example 1", args{"389125467", 10}, "92658374"},
		{ "Day 23 part 1 example 2", args{"389125467", 100}, "67384529"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay23a(tt.args.s, tt.args.rounds); got != tt.want {
				t.Errorf("ComputeDay23a() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputeDay23b(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{ "Day 23 Part 2 example 1", args{"389125467"}, 149245887792},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay23b(tt.args.s); got != tt.want {
				t.Errorf("ComputeDay23b() = %v, want %v", got, tt.want)
			}
		})
	}
}