package main

import "testing"

func TestComputeDay2a(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Basic test case part 1", args{"1-3 a: abcde\n1-3 b: cdefg\n2-9 c: ccccccccc"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay2a(tt.args.input); got != tt.want {
				t.Errorf("ComputeDay2a() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputeDay2b(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Basic test case part 1", args{"1-3 a: abcde\n1-3 b: cdefg\n2-9 c: ccccccccc"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay2b(tt.args.input); got != tt.want {
				t.Errorf("ComputeDay2b() = %v, want %v", got, tt.want)
			}
		})
	}
}