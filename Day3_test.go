package main

import (
	"strings"
	"testing"
)

func TestComputeDay3a(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{ "Day 3 Part A example", args{  "..##.......\n#...#...#..\n.#....#..#.\n..#.#...#.#\n.#...##..#.\n..#.##.....\n.#.#.#....#\n.#........#\n#.##...#...\n#...##....#\n.#..#...#.#"}, 7 },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay3a(tt.args.input); got != tt.want {
				t.Errorf("ComputeDay3a() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputeNumTrees(t *testing.T) {
	type args struct {
		lines []string
		xinc  int
		yinc  int
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{ "Day 3 Part B example (1, 1)", args{  strings.Split("..##.......\n#...#...#..\n.#....#..#.\n..#.#...#.#\n.#...##..#.\n..#.##.....\n.#.#.#....#\n.#........#\n#.##...#...\n#...##....#\n.#..#...#.#", "\n"), 1, 1}, 2 },
		{ "Day 3 Part B example (3, 1)", args{  strings.Split("..##.......\n#...#...#..\n.#....#..#.\n..#.#...#.#\n.#...##..#.\n..#.##.....\n.#.#.#....#\n.#........#\n#.##...#...\n#...##....#\n.#..#...#.#", "\n"), 3, 1}, 7 },
		{ "Day 3 Part B example (5, 1)", args{  strings.Split("..##.......\n#...#...#..\n.#....#..#.\n..#.#...#.#\n.#...##..#.\n..#.##.....\n.#.#.#....#\n.#........#\n#.##...#...\n#...##....#\n.#..#...#.#", "\n"), 5, 1}, 3 },
		{ "Day 3 Part B example (7, 1)", args{  strings.Split("..##.......\n#...#...#..\n.#....#..#.\n..#.#...#.#\n.#...##..#.\n..#.##.....\n.#.#.#....#\n.#........#\n#.##...#...\n#...##....#\n.#..#...#.#", "\n"), 7, 1}, 4 },
		{ "Day 3 Part B example (1, 2)", args{  strings.Split("..##.......\n#...#...#..\n.#....#..#.\n..#.#...#.#\n.#...##..#.\n..#.##.....\n.#.#.#....#\n.#........#\n#.##...#...\n#...##....#\n.#..#...#.#", "\n"), 1, 2}, 2 },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumTrees(tt.args.lines, tt.args.xinc, tt.args.yinc); got != tt.want {
				t.Errorf("NumTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputeDay3b(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{ "Day 3 Part B example", args{  "..##.......\n#...#...#..\n.#....#..#.\n..#.#...#.#\n.#...##..#.\n..#.##.....\n.#.#.#....#\n.#........#\n#.##...#...\n#...##....#\n.#..#...#.#"}, 336 },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay3b(tt.args.input); got != tt.want {
				t.Errorf("ComputeDay3b() = %v, want %v", got, tt.want)
			}
		})
	}
}