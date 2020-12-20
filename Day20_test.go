package main

import (
	"io/ioutil"
	"testing"
)

func TestComputeDay20a(t *testing.T) {
	inputpath := "D:/Home/projects/adventcalendar/2020/inputs/day20-test.dat"
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
		want int64
	}{
		{"Day 20 example part 1", args{string(bytes)}, 20899048083289},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay20a(tt.args.input); got != tt.want {
				t.Errorf("ComputeDay20b() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputeDay20b(t *testing.T) {
	inputpath := "D:/Home/projects/adventcalendar/2020/inputs/day20-test.dat"
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
		{"Day 20 example part 2", args{string(bytes)}, 273},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay20b(tt.args.input); got != tt.want {
				t.Errorf("ComputeDay20b() = %v, want %v", got, tt.want)
			}
		})
	}
}

