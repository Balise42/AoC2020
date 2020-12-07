package main

import (
	"io/ioutil"
	"testing"
)

func TestComputeDay7a(t *testing.T) {
	inputpath := "D:/Home/projects/adventcalendar/2020/inputs/day7-test.dat"
	bytes, err := ioutil.ReadFile(inputpath)
	if err != nil {
		panic("No input file found")
	}

	type args struct {
		input string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"Day 8 example part 1", args{string(bytes)}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay7a(tt.args.input); got != tt.want {
				t.Errorf("ComputeDay7a() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputeDay7b(t *testing.T) {
	inputpath := "D:/Home/projects/adventcalendar/2020/inputs/day7-test.dat"
	bytes, err := ioutil.ReadFile(inputpath)
	if err != nil {
		panic("No input file found")
	}

	type args struct {
		input string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"Day 8 example part 2", args{string(bytes)}, 32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay7b(tt.args.input); got != tt.want {
				t.Errorf("ComputeDay7b() = %v, want %v", got, tt.want)
			}
		})
	}
}
