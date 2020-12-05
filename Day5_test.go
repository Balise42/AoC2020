package main

import "testing"

func TestComputeSeatId(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Day 5 ID example 1", args{"BFFFBBFRRR"}, 567 },
		{"Day 5 ID example 2", args{"FFFBBBFRRR"}, 119 },
		{"Day 5 ID example 2", args{"BBFFBBFRLL"}, 820 },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeSeatId(tt.args.line); got != tt.want {
				t.Errorf("ComputeSeatId() = %v, want %v", got, tt.want)
			}
		})
	}
}