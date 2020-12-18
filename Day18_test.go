package main

import "testing"

/*func TestComputeExpressionDay18a(t *testing.T) {
	type args struct {
		expr string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Day 18 Part 1 example no parentheses", args{"1 + 2 * 3 + 4 * 5 + 6"}, 71},
		{"Day 18 Part 1 example 0", args{"1 + (2 * 3) + (4 * (5 + 6))"}, 51},
		{"Day 18 Part 1 example 1", args{"2 * 3 + (4 * 5)"}, 26},
		{"Day 18 Part 1 example 2", args{"5 + (8 * 3 + 9 + 3 * 4 * 3)"}, 437},
		{"Day 18 Part 1 example 3", args{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"}, 12240 },
		{"Day 18 Part 1 example 4", args{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"}, 13632 },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay18a(tt.args.expr); got != tt.want {
				t.Errorf("ComputeParentheseFreeExpression() = %v, want %v", got, tt.want)
			}
		})
	}
}*/

func TestComputeDay18b(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Day 18 Part 2 example no parentheses", args{"1 + 2 * 3 + 4 * 5 + 6"}, 231},
		{"Day 18 Part 2 example 0", args{"1 + (2 * 3) + (4 * (5 + 6))"}, 51},
		{"Day 18 Part 2 example 1", args{"2 * 3 + (4 * 5)"}, 46},
		{"Day 18 Part 2 example 2", args{"5 + (8 * 3 + 9 + 3 * 4 * 3)"}, 1445},
		{"Day 18 Part 2 example 3", args{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"}, 669060 },
		{"Day 18 Part 2 example 4", args{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"}, 23340 },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeDay18b(tt.args.input); got != tt.want {
				t.Errorf("ComputeDay18b() = %v, want %v", got, tt.want)
			}
		})
	}
}