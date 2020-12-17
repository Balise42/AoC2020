package main

import "testing"

func TestComputeDay17a(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Day 17 Part 1 example", args{".#.\n..#\n###"}, 112},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay17a(tt.args.input); got != tt.want {
				t.Errorf("ComputeDay17a() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputeDay17b(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Day 17 Part 2 example", args{".#.\n..#\n###"}, 848},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay17b(tt.args.input); got != tt.want {
				t.Errorf("ComputeDay17b() = %v, want %v", got, tt.want)
			}
		})
	}
}