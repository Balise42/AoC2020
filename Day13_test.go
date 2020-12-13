package main

import "testing"

func TestComputeDay13a(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Compute Day 13 - part 1 - example", args {"939\n7,13,x,x,59,x,31,19"} , 295},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay13a(tt.args.input); got != tt.want {
				t.Errorf("ComputeDay13a() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputeDay13b(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Compute Day 13 - part 2 - example 1", args {"939\n7,13,x,x,59,x,31,19"} , 1068781},
		{"Compute Day 13 - part 2 - example 2", args {"939\n17,x,13,19"} , 3417},
		{"Compute Day 13 - part 2 - example 3", args {"939\n67,7,59,61"} , 754018},
		{"Compute Day 13 - part 2 - example 4", args {"939\n67,x,7,59,61"} , 779210},
		{"Compute Day 13 - part 2 - example 5", args {"939\n67,7,x,59,61"} , 1261476},
		{"Compute Day 13 - part 2 - example 6", args {"939\n1789,37,47,1889"} , 1202161486},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay13b(tt.args.input); got != tt.want {
				t.Errorf("ComputeDay13b() = %v, want %v", got, tt.want)
			}
		})
	}
}