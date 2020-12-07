package main

import (
	"fmt"
	"strconv"
	"strings"
)

func (day *Day) Day7a() {
	fmt.Printf("Part 1: %d\n", ComputeDay7a(day.input))
}

func (day *Day) Day7b() {
	fmt.Printf("Part 2: %d\n", ComputeDay7b(day.input))
}

type Bag struct {
	qualifier string
	color string
}

type BagNode struct {
	bag Bag
	contains map[Bag]int
	contained []Bag
}

var graph map[Bag]*BagNode


func ComputeDay7a(input string) int {
	graph = make(map[Bag]*BagNode)

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parseLine(line)
	}

	startBag := Bag { "shiny", "gold" }
	parents := make(map[Bag]bool)
	nodes := make([]Bag, 0)
	nodes = append(nodes, startBag)

	for len(nodes) != 0 {
		bag := nodes[0]
		nodes = nodes[1:]

		for _, v := range graph[bag].contained {
			_, ok := parents[v] ; if !ok {
				nodes = append(nodes, v)
				parents[v] = true
			}
		}
	}

	return len(parents)
}

func parseLine(line string) {
	toks := strings.Split(line, " contain ")

	bagtoks := strings.Split(toks[0], " ")
	bag1 := Bag { bagtoks[0], bagtoks[1] }

	var bag1Node * BagNode
	var ok bool
	if bag1Node, ok = graph[bag1]; !ok {
		bag1Node = &BagNode { bag1, make(map[Bag]int), make([]Bag, 0)}
		graph[bag1] = bag1Node
	}

	if toks[1] != "no other bags." {
		containedBagsList := strings.Split(toks[1], ", ")
		for _, containedBagElem := range containedBagsList {
			containedBagToks := strings.Split(containedBagElem, " ")
			containedBag := Bag{containedBagToks[1], containedBagToks[2]}

			var containedBagNode *BagNode
			if containedBagNode, ok = graph[containedBag]; !ok {
				containedBagNode = &BagNode{containedBag, make(map[Bag]int), make([]Bag, 0)}
				graph[containedBag] = containedBagNode
			}
			num, err := strconv.Atoi(containedBagToks[0])
			if err != nil {
				panic("Could not parse bag " + line)
			}

			bag1Node.contains[containedBag] = num
			containedBagNode.contained = append(containedBagNode.contained, bag1)
		}
	}
}


func ComputeDay7b(input string) int {
	graph = make(map[Bag]*BagNode)

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parseLine(line)
	}

	startBag := Bag { "shiny", "gold" }
	return numChildBags(startBag)
}

func numChildBags(bag Bag) int {
	nodeBag := graph[bag]
	res := 0
	for k, v := range nodeBag.contains {
		res = res +  v * (1 + numChildBags(k))
	}
	return res
}
