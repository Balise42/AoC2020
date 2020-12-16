package main

import "testing"

func TestComputeDay16a(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Compute day 16 part 1 example", args{"class: 1-3 or 5-7\nrow: 6-11 or 33-44\nseat: 13-40 or 45-50\n\nyour ticket:\n7,1,14\n\nnearby tickets:\n7,3,47\n40,4,50\n55,2,20\n38,6,12"}, 71},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay16a(tt.args.input); got != tt.want {
				t.Errorf("ComputeDay16a() = %v, want %v", got, tt.want)
			}
		})
	}
}