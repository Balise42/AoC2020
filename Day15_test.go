package main

import "testing"

func TestComputeDay15a(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{  "Day 15 Part 1 example 0", args{"0,3,6"}, 436},
		{  "Day 15 Part 1 example 1", args{"1,3,2"}, 1},
		{  "Day 15 Part 1 example 2", args{"2,1,3"}, 10},
		{  "Day 15 Part 1 example 3", args{"1,2,3"}, 27},
		{  "Day 15 Part 1 example 4", args{"2,3,1"}, 78},
		{  "Day 15 Part 1 example 5", args{"3,2,1"}, 438},
		{  "Day 15 Part 1 example 6", args{"3,1,2"}, 1836},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay15a(tt.args.input); got != tt.want {
				t.Errorf("ComputeDay15a() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputeDay15b(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{  "Day 15 Part 2 example 0", args{"0,3,6"}, 175594},
		/*{  "Day 15 Part 2 example 1", args{"1,3,2"}, 2578},
		{  "Day 15 Part 2 example 2", args{"2,1,3"}, 3544142},
		{  "Day 15 Part 2 example 3", args{"1,2,3"}, 261214},
		{  "Day 15 Part 2 example 4", args{"2,3,1"}, 6895259},
		{  "Day 15 Part 2 example 5", args{"3,2,1"}, 18},
		{  "Day 15 Part 2 example 6", args{"3,1,2"}, 362},*/
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay15b(tt.args.input); got != tt.want {
				t.Errorf("ComputeDay15b() = %v, want %v", got, tt.want)
			}
		})
	}
}