package main

import "testing"

func TestComputeDay21a(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Day 21 part 1 example", args {"mxmxvkd kfcds sqjhc nhms (contains dairy, fish)\ntrh fvjkl sbzzf mxmxvkd (contains dairy)\nsqjhc fvjkl (contains soy)\nsqjhc mxmxvkd sbzzf (contains fish)"}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := ComputeDay21a(tt.args.input); got != tt.want {
				t.Errorf("ComputeDay21a() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputeDay21b(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Day 21 part 1 example", args {"mxmxvkd kfcds sqjhc nhms (contains dairy, fish)\ntrh fvjkl sbzzf mxmxvkd (contains dairy)\nsqjhc fvjkl (contains soy)\nsqjhc mxmxvkd sbzzf (contains fish)"}, "mxmxvkd,sqjhc,fvjkl"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay21b(tt.args.input); got != tt.want {
				t.Errorf("ComputeDay21b() = %v, want %v", got, tt.want)
			}
		})
	}
}