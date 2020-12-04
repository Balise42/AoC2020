package main

import (
	"fmt"
	"strconv"
	"strings"
)

func (day *Day) Day4a() {
	fmt.Printf("Part1: %d\n", ComputeDay4a(day.input))
}

func (day *Day) Day4b() {
	fmt.Printf("Part2: %d\n", ComputeDay4b(day.input))
}

func ComputeDay4a(input string) int {
	passports := strings.Split(input, "\n\n")
	validPassports := 0
	for _, passport := range passports {
		if isValidPassportPart1(passport) {
			validPassports++
		}
	}
	return validPassports
}

func ComputeDay4b(input string) int {
	passports := strings.Split(input, "\n\n")
	validPassports := 0
	for _, passport := range passports {
		if isValidPassportPart1(passport) && isValidPassportPart2(passport) {
			validPassports++
		}
	}
	return validPassports
}

func isValidPassportPart1(passport string) bool {
	return strings.Contains(passport, "byr:") && strings.Contains(passport, "iyr:") &&
		strings.Contains(passport, "eyr:") && strings.Contains(passport, "hgt:") &&
		strings.Contains(passport, "hcl:") && strings.Contains(passport, "ecl:") &&
		strings.Contains(passport, "pid:")
}

type Day4 struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
}

func isValidPassportPart2(passport string) bool {
	fields := strings.Fields(passport)
	var passFields Day4
	for _, field := range fields {
		kv := strings.Split(field, ":")
		switch kv[0] {
		case "byr":
			passFields.byr = kv[1]
		case "iyr":
			passFields.iyr = kv[1]
		case "eyr":
			passFields.eyr = kv[1]
		case "hgt":
			passFields.hgt = kv[1]
		case "hcl":
			passFields.hcl = kv[1]
		case "ecl":
			passFields.ecl = kv[1]
		case "pid":
			passFields.pid = kv[1]
		}
	}

	return validatePassport(passFields)
}

func validatePassport(passport Day4) bool {
	// birth year validation
	if len(passport.byr) != 4 {
		return false
	}
	byr, err := strconv.Atoi(passport.byr)
	if err != nil {
		return false
	}
	if byr < 1920 || byr > 2002 {
		return false
	}

	// issue year validation
	if len(passport.iyr) != 4 {
		return false
	}
	iyr, err := strconv.Atoi(passport.iyr)
	if err != nil {
		return false
	}
	if iyr < 2010 || iyr > 2020 {
		return false
	}

	// expiration year validation
	if len(passport.eyr) != 4 {
		return false
	}
	eyr, err := strconv.Atoi(passport.eyr)
	if err != nil {
		return false
	}
	if eyr < 2020 || eyr > 2030 {
		return false
	}

	// height validation
	if len(passport.hgt) != 5 && len(passport.hgt) != 4 {
		return false
	}
	if len(passport.hgt) == 4 {
		if passport.hgt[2] != 'i' || passport.hgt[3] != 'n' {
			return false
		}
		hgt, err := strconv.Atoi(passport.hgt[:2])
		if err != nil {
			return false
		}
		if hgt < 59 || hgt > 76 {
			return false
		}
	}

	if len(passport.hgt) == 5 {
		if passport.hgt[3] != 'c' || passport.hgt[4] != 'm' {
			return false
		}
		hgt, err := strconv.Atoi(passport.hgt[:3])
		if err != nil {
			return false
		}
		if hgt < 150 || hgt > 193 {
			return false
		}
	}

	// hair color validation
	if len(passport.hcl) != 7 {
		return false
	}
	if passport.hcl[0] != '#' {
		return false
	}
	// if we have a hex number then we're all good
	_, err = strconv.ParseInt(passport.hcl[1:], 16, 64)
	if err != nil {
		return false
	}

	// eye color validation
	ecl := passport.ecl
	if ecl != "amb" && ecl != "blu" && ecl != "brn" && ecl != "gry" && ecl != "grn" && ecl != "hzl" && ecl != "oth" {
		return false
	}

	// passport id validation
	if len(passport.pid) != 9 {
		return false
	}
	_, err = strconv.Atoi(passport.pid)
	if err != nil {
		return false
	}

	return true
}
