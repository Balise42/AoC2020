package main

import "testing"

func TestComputeDay9a(t *testing.T) {
	type args struct {
		input  string
		window int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Day 9 example - part 1", args {"35\n20\n15\n25\n47\n40\n62\n55\n65\n95\n102\n117\n150\n182\n127\n219\n299\n277\n309\n576", 5}, 127 },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay9a(tt.args.input, tt.args.window); got != tt.want {
				t.Errorf("ComputeDay9a() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputeDayb(t *testing.T) {
	type args struct {
		input  string
		window int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Day 9 example - part 2", args {"35\n20\n15\n25\n47\n40\n62\n55\n65\n95\n102\n117\n150\n182\n127\n219\n299\n277\n309\n576", 5}, 62 },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay9b(tt.args.input, tt.args.window); got != tt.want {
				t.Errorf("ComputeDay9a() = %v, want %v", got, tt.want)
			}
		})
	}
}