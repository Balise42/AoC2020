package main

import "testing"

func TestComputeDay10a(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Day 10 part 1 example 1", args{"16\n10\n15\n5\n1\n11\n7\n19\n6\n12\n4"}, 35},
		{"Day 10 part 1 example 2", args{"28\n33\n18\n42\n31\n14\n46\n20\n48\n47\n24\n23\n49\n45\n19\n38\n39\n11\n1\n32\n25\n35\n8\n17\n7\n9\n4\n2\n34\n10\n3"}, 220},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay10a(tt.args.input); got != tt.want {
				t.Errorf("ComputeDay10a() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputeDay10b(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{"Day 10 part 2 example 1", args{"16\n10\n15\n5\n1\n11\n7\n19\n6\n12\n4"}, 8},
		{"Day 10 part 2 example 2", args{"28\n33\n18\n42\n31\n14\n46\n20\n48\n47\n24\n23\n49\n45\n19\n38\n39\n11\n1\n32\n25\n35\n8\n17\n7\n9\n4\n2\n34\n10\n3"}, 19208},
		{ "Day 10 part 2 easy example", args {"1\n2\n3\n4"}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay10b(tt.args.input); got != tt.want {
				t.Errorf("ComputeDay10b() = %v, want %v", got, tt.want)
			}
		})
	}
}