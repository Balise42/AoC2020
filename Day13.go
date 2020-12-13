package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

func (day *Day) Day13a() {
	fmt.Printf("Part 1: %d\n", ComputeDay13a(day.input))
}

func (day *Day) Day13b() {
	fmt.Printf("Part 2: %d\n", ComputeDay13b(day.input))
}

func ComputeDay13a(input string) int {
	lines := strings.Split(input, "\n")
	earliest, err := strconv.Atoi(lines[0])
	if err != nil {
		panic("Cannot parse earliest time in inpus")
	}
	busstr := strings.Split(lines[1], ",")
	buses := make([]int, 0)
	for _, str := range busstr {
		if str != "x" {
			newBus, err := strconv.Atoi(str)
			if err != nil {
				panic("Cannot parse bus")
			}
			buses = append(buses, newBus)
		}
	}

	waittime := math.MaxInt64
	busnr := -1
	for _, bus := range buses {
		newWaittime := (earliest/bus+1)*bus - earliest
		if newWaittime < waittime {
			waittime = newWaittime
			busnr = bus
		}
	}
	return busnr * waittime
}

func ComputeDay13b(input string) int64 {
	lines := strings.Split(input, "\n")
	busstr := strings.Split(lines[1], ",")
	modulos := make([]int64, 0)
	remainders := make([]int64, 0)
	for i, str := range busstr {
		if str != "x" {
			newBus, err := strconv.Atoi(str)
			if err != nil {
				panic("Cannot parse bus")
			}
			modulos = append(modulos, int64(newBus))
			rem := -i
			for rem < 0 {
				rem = rem + newBus
			}
			remainders = append(remainders, int64(rem))
		}
	}

	res, mod := solveTwo(modulos[0], modulos[1], remainders[0], remainders[1])
	for i := 2; i < len(modulos); i++ {
		res, mod = solveTwo(mod, modulos[i], res, remainders[i])
	}
	return res
}

func solveTwo(mod1 int64, mod2 int64, a1 int64, a2 int64) (int64, int64) {
	m1, m2 := computeBezoutCoefficients(mod1, mod2)

	bm1 := new(big.Int).SetInt64(m1)
	bm2 := new(big.Int).SetInt64(m2)
	ba1 := new(big.Int).SetInt64(a1)
	ba2 := new(big.Int).SetInt64(a2)
	bmod1 := new(big.Int).SetInt64(mod1)
	bmod2 := new(big.Int).SetInt64(mod2)

	x := new(big.Int).SetInt64(0)
	x = x.Mul(ba1, bm2)
	x = x.Mul(x, bmod2)

	y := new(big.Int).SetInt64(0)
	y = y.Mul(ba2, bm1)
	y = y.Mul(y, bmod1)

	x = x.Add(x, y)

	prod := new(big.Int).SetInt64(0)
	prod = prod.Mul(bmod1, bmod2)

	x = x.Mod(x, prod)

	return x.Int64(), mod1 * mod2
}

func computeBezoutCoefficients(a int64, b int64) (int64, int64) {
	oldR, r := a, b
	oldS, s := int64(1), int64(0)
	oldT, t := int64(0), int64(1)

	for r != 0 {
		quotient := oldR / r
		oldR, r = r, oldR-quotient*r
		oldS, s = s, oldS-quotient*s
		oldT, t = t, oldT-quotient*t
	}

	return oldS, oldT
}