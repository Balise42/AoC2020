package main

import (
	"fmt"
	"strings"
)

func (day *Day) Day19a() {
	fmt.Printf("Part 1: %d\n", ComputeDay19a(day.input))
}

func (day *Day) Day19b() {
	fmt.Printf("Part 2: %d\n", ComputeDay19b(day.input))
}

type Day19Rule interface {
	matches(message string) (bool, []string)
}

type Day19Sequence struct {
	rules []Day19Rule
}

type Day19Options struct {
	rules []Day19Rule
}

type Day19Char struct {
	val uint8
}

func (seq Day19Sequence) matches(message string) (bool, []string) {
	prefixes := make([]string, 1)
	prefixes[0] = ""

	for _, rule := range seq.rules {
		newPrefixes := make([]string, 0)
		for _, p := range prefixes {
			str := message
			if p != "" {
				if !strings.HasPrefix(str, p) {
					panic("Unexpected: string [" + str + "] is supposed to start with [" + p + "]")
				}
				str = strings.TrimPrefix(str, p)
			}
			matchRule, rulePrefixes := rule.matches(str)
			if !matchRule {
				return false, []string{}
			} else {
				for _, rulePrefix := range rulePrefixes {
					newPrefixes = append(newPrefixes, p+rulePrefix)
				}
			}
		}
		prefixes = newPrefixes
	}

	prefixes = cleanupPrefixes(prefixes)

	return true, prefixes
}

func (options Day19Options) matches(message string) (bool, []string) {
	match := false
	prefixes := make([]string, 0)

	for _, rule := range options.rules {
		matchRule, rulePrefixes := rule.matches(message)
		match = match || matchRule
		prefixes = append(prefixes, rulePrefixes...)
	}

	prefixes = cleanupPrefixes(prefixes)
	return match, prefixes
}

func cleanupPrefixes(prefixes []string) []string {
	prefixSet := make(map[string]bool, len(prefixes))
	for _, p := range prefixes {
		prefixSet[p] = true
	}
	cleanList := make([]string, 0, len(prefixSet))
	for k := range prefixSet {
		cleanList = append(cleanList, k)
	}
	return cleanList
}

func (val Day19Char) matches(message string) (bool, []string) {
	if len(message) == 0 {
		return false, []string{}
	}
	if message[0] == val.val {
		return true, []string{message[:1]}
	}
	return false, []string{}
}

func ComputeDay19a(input string) int {
	parts := strings.Split(input, "\n\n")
	rulesStr := strings.Split(parts[0], "\n")
	messages := strings.Split(parts[1], "\n")

	rules := parseDay19Rules(rulesStr, false)

	res := 0
	for _, message := range messages {
		ok, prefixes := rules["0"].matches(message)
		if !ok {
			continue
		}
		for _, prefix := range prefixes {
			if message == prefix {
				res++
				break
			}
		}
	}
	return res
}

func parseDay19Rules(rulesStr []string, part2 bool) map[string]Day19Rule {
	rulesStrMap := make(map[string]string, len(rulesStr))
	for _, str := range rulesStr {
		toks := strings.Split(str, ": ")
		rulesStrMap[toks[0]] = toks[1]
	}

	rulesMap := make(map[string]Day19Rule, len(rulesStr))

	for key := range rulesStrMap {
		if !part2 {
			parseDay19Rule(key, &rulesStrMap, &rulesMap)
		} else if key != "0" {
			parseDay19Rule(key, &rulesStrMap, &rulesMap)
		}
	}

	return rulesMap
}

func parseDay19Rule(key string, strMap *map[string]string, ruleMap *map[string]Day19Rule) Day19Rule {
	if rule, ok := (*ruleMap)[key]; ok {
		return rule
	}

	ruleStr, ok := (*strMap)[key]
	if !ok {
		panic("Could not find rule with key [" + key + "]")
	}
	if ruleStr[0] == '"' {
		(*ruleMap)[key] = Day19Char{ruleStr[1]}
	} else if strings.Index(ruleStr, " | ") >= 0 {
		var newRule Day19Options
		newRule.rules = make([]Day19Rule, 0)
		for _, subRule := range strings.Split(ruleStr, " | ") {
			newRule.rules = append(newRule.rules, parseSequence(subRule, strMap, ruleMap))
		}
		(*ruleMap)[key] = newRule
	} else {
		(*ruleMap)[key] = parseSequence(ruleStr, strMap, ruleMap)
	}
	return (*ruleMap)[key]
}

func parseSequence(str string, strMap *map[string]string, ruleMap *map[string]Day19Rule) Day19Sequence {
	toks := strings.Split(str, " ")
	var newRule Day19Sequence
	newRule.rules = make([]Day19Rule, 0)
	for _, subRule := range toks {
		newRule.rules = append(newRule.rules, parseDay19Rule(subRule, strMap, ruleMap))
	}
	return newRule
}

func ComputeDay19b(input string) int {
	parts := strings.Split(input, "\n\n")
	rulesStr := strings.Split(parts[0], "\n")
	messages := strings.Split(parts[1], "\n")

	rules := parseDay19Rules(rulesStr, true)

	res := 0
	for _, message := range messages {
		messageOk := false
		if rule8prefixes, ok := matchesDay19Rule8(message, rules); ok {
			for _, prefix := range rule8prefixes {
				if !messageOk {
					toCheck := strings.TrimPrefix(message, prefix)
					rule11prefixes, ok := matchesDay19Rule11(toCheck, rules)
					if !ok {
						continue
					}
					for _, rule11prefix := range rule11prefixes {
						if toCheck == rule11prefix {
							messageOk = true
							break
						}
					}
				}
			}
		}
		if messageOk {
			res++
		}
	}
	return res
}

func matchesDay19Rule8(message string, rules map[string]Day19Rule) ([]string, bool) {
	match, prefixes := rules["42"].matches(message)
	if !match {
		return []string{}, false
	}

	for i := 0; i < len(prefixes); i++ {
		if strings.HasPrefix(message, prefixes[i]) {
			newPrefixes, ok := matchesDay19Rule8(strings.TrimPrefix(message, prefixes[i]), rules)
			if ok {
				for _, newPrefix := range newPrefixes {
					prefixes = append(prefixes, prefixes[i]+newPrefix)
				}
			}
		}
	}

	prefixes = cleanupPrefixes(prefixes)
	return prefixes, true
}

func matchesDay19Rule11(message string, rules map[string]Day19Rule) ([]string, bool) {
	match, prefixes := rules["42"].matches(message)
	if !match {
		return []string{}, false
	}

	countsRules42 := make(map[string]int)
	for _, prefix := range prefixes {
		countsRules42[prefix] = 1
	}

	for i := 0; i < len(prefixes); i++ {
		if strings.HasPrefix(message, prefixes[i]) {
			matches, newPrefixes := rules["42"].matches(strings.TrimPrefix(message, prefixes[i]))
			if matches {
				for _, newPrefix := range newPrefixes {
					prefixes = append(prefixes, prefixes[i]+newPrefix)
					countsRules42[prefixes[i] + newPrefix] = countsRules42[prefixes[i]] + 1
				}
			}
		}
	}

	ret := make([]string, 0)

	for _, prefix := range prefixes {
		countRule42 := countsRules42[prefix]

		currStepPrefixes := []string {prefix}
		nextStepPrefixes := make([]string, 0)
		valid := true
		for i := 0; i < countRule42; i++ {
			for _, currStepPrefix := range currStepPrefixes {
				matches, newPrefixes := rules["31"].matches(strings.TrimPrefix(message, currStepPrefix))
				if matches {
					for _, newPrefix := range newPrefixes {
						nextStepPrefixes = append(nextStepPrefixes, currStepPrefix + newPrefix)
					}
				}
			}
			if len(nextStepPrefixes) == 0 {
				valid = false
				break
			}
			currStepPrefixes = nextStepPrefixes
			nextStepPrefixes = make([]string, 0)
		}
		if valid{
			ret = append(ret, currStepPrefixes...)
		}
	}

	if len(ret) > 0 {
		return ret, true
	}
	return []string {}, false
}
