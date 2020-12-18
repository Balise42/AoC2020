package main

import (
	"fmt"
	"strconv"
	"strings"
)

func (day *Day) Day18a() {
	fmt.Printf("Part 1: %d\n", ComputeDay18a(day.input))
}

func (day *Day) Day18b() {
	fmt.Printf("Part 2: %d\n", ComputeDay18b(day.input))
}

func ComputeDay18a(input string) int {
	lines := strings.Split(input, "\n")
	res := 0
	for _, line := range lines {
		res = res + ComputeExpressionDay18a(line, false)
	}
	return res
}

func ComputeDay18b(input string) int {
	lines := strings.Split(input, "\n")
	res := 0
	for _, line := range lines {
		res = res + ComputeExpressionDay18a(line, true)
	}
	return res
}

func ComputeExpressionDay18a(line string, part2 bool) int {
	expr := line

	leftPar := strings.Index(expr, "(")

	for leftPar >= 0 {
		rightPar := findFirstMatchingRightPar(expr)
		enclosedResult := ComputeExpressionDay18a(expr[leftPar + 1:rightPar], part2)

		suffix := ""
		if rightPar < len(expr) - 1 {
			suffix = expr[rightPar + 1:]
		}
		expr = expr[:leftPar] + strconv.Itoa(enclosedResult) + suffix
		leftPar = strings.Index(expr, "(")
	}

	if !part2 {
		return ComputeParentheseFreeExpressionPart1(expr)
	} else {
		return ComputeParentheseFreeExpressionPart2(expr)
	}
}

func findFirstMatchingRightPar(expr string) int {
	nesting := 0
	for i, c := range expr {
		if c == '(' {
			nesting++
		} else if c == ')' {
			nesting--
			if nesting == 0 {
				return i
			}
		}
	}
	panic("Could not find matching parenthesis " + expr)
}

func ComputeParentheseFreeExpressionPart1(expr string) int {
	lhs := -1
	rhs := -1
	var op int32
	var err error

	buf := ""
	for _, c := range expr {
		if c == '*' || c == '+' {
			if lhs == - 1 {
				lhs, err = strconv.Atoi(buf)
				if err != nil {
					panic("Could not parse first number "+ buf)
				}
				buf = ""
			} else {
				rhs, err = strconv.Atoi(buf)
				if err != nil {
					panic("Could not parse rhs " + buf)
				}
				buf = ""
				if op == '+' {
					lhs = lhs + rhs
				} else if op == '*' {
					lhs = lhs * rhs
				} else {
					panic("Unknown operator " + string(op))
				}
			}
			op = c
		} else if c == ' ' {
			continue
		} else {
			buf = buf + string(c)
		}
	}
	rhs, err = strconv.Atoi(buf)
	if err != nil {
		panic("Could not parse rhs " + buf)
	}
	if op == '+' {
		lhs = lhs + rhs
	} else if op == '*' {
		lhs = lhs * rhs
	} else {
		return rhs
	}
	return lhs
}

func ComputeParentheseFreeExpressionPart2(line string) int {
	expr := ComputeAdditions(line)
	return ComputeParentheseFreeExpressionPart1(expr)
}

func ComputeAdditions(line string) string {
	expr := line
	multPos := strings.Index(line, "+")
	if multPos < 0 {
		return expr
	}

	startPos := 0
	for i := multPos - 1; i >= 0; i-- {
		if expr[i] == '*' {
			startPos = i + 1
			break
		}
	}

	endPos := multPos + 1
	for ; endPos < len(expr); endPos++ {
		if expr[endPos] == '+' || expr[endPos] == '*' {
			break
		}
	}

	return ComputeAdditions(expr[:startPos] + strconv.Itoa(ComputeParentheseFreeExpressionPart1(expr[startPos:endPos]))+ expr[endPos:])
}