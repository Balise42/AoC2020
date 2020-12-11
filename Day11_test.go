package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestComputeDay11a(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{ "Day 11 Part 1 - Example", args {"L.LL.LL.LL\nLLLLLLL.LL\nL.L.L..L..\nLLLL.LL.LL\nL.LL.LL.LL\nL.LLLLL.LL\n..L.L.....\nLLLLLLLLLL\nL.LLLLLL.L\nL.LLLLL.LL"}, 37 },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay11a(tt.args.input); got != tt.want {
				t.Errorf("ComputeDay11a() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputeDay11b(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{ "Day 11 Part 2 - Example", args {"L.LL.LL.LL\nLLLLLLL.LL\nL.L.L..L..\nLLLL.LL.LL\nL.LL.LL.LL\nL.LLLLL.LL\n..L.L.....\nLLLLLLLLLL\nL.LLLLLL.L\nL.LLLLL.LL"}, 26 },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay11b(tt.args.input); got != tt.want {
				t.Errorf("ComputeDay11b() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputeNumberOccupiedVisibleSeats(t *testing.T) {
	type args struct {
		grid []string
		i    int
		j    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{  "Test visible occupied seats 1", args {strings.Split(".......#.\n...#.....\n.#.......\n.........\n..#L....#\n....#....\n.........\n#........\n...#.....", "\n"), 4, 3 }, 8},
		{  "Test visible occupied seats 2", args {strings.Split(".##.##.\n#.#.#.#\n##...##\n...L...\n##...##\n#.#.#.#\n.##.##.", "\n"), 3, 3 }, 0},
		{  "Test visible occupied seats 3", args {strings.Split(".............\n.L.L.#.#.#.#.\n.............", "\n"), 1, 1 }, 0},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberOccupiedVisibleSeats(tt.args.grid, tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("NumberOccupiedVisibleSeats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputeNextSeating(t *testing.T) {
	type args struct {
		grid  []string
		part2 bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		/*{"Test next seating 2", args {strings.Split("#.L#.L#.L#\n#LLLLLL.LL\nL.L.L..#..\n##L#.#L.L#\nL.L#.LL.L#\n#.LLLL#.LL\n..#.L.....\nLLL###LLL#\n#.LLLLL#.L\n#.L#LL#.L#","\n"), true},
			strings.Split("#.L#.L#.L#\n#LLLLLL.LL\nL.L.L..#..\n##L#.#L.L#\nL.L#.LL.L#\n#.LLLL#.LL\n..#.L.....\nLLL###LLL#\n#.LLLLL#.L\n#.L#LL#.L#", "\n")},*/
		{"Test next seating 2", args {strings.Split("#.L#.##.L#\n#L#####.LL\nL.#.#..#..\n##L#.##.##\n#.##.#L.##\n#.#####.#L\n..#.#.....\nLLL####LL#\n#.L#####.L\n#.L####.L#","\n"), true},
			strings.Split("#.L#.L#.L#\n#LLLLLL.LL\nL.L.L..#..\n##LL.LL.L#\nL.LL.LL.L#\n#.LLLLL.LL\n..L.L.....\nLLLLLLLLL#\n#.LLLLL#.L\n#.L#LL#.L#", "\n")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeNextSeating(tt.args.grid, tt.args.part2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("computeNextSeating() = %v, want %v", got, tt.want)
			}
		})
	}
}