package main

import (
	"fmt"
	"strconv"
)

func (day *Day) Day23a() {
	fmt.Printf("Part 1: %s\n", ComputeDay23a("193467258", 100, 9))
}

func (day *Day) Day23b() {
	fmt.Printf("Part 2: %d\n", ComputeDay23b("193467258"))
}


func ComputeDay23a(s string, rounds int, max int) string {
	var first *Day23Node
	var prev *Day23Node
	var initFirst = false
	var latestNode *Day23Node

	for _, c := range s {
		cup, err := strconv.Atoi(string(c))
		if err != nil {
			panic("Cannot interpret cup " + string(c))
		}
		if initFirst == false {
			first = &Day23Node{cup, nil, nil}
			prev = first
			initFirst = true
		} else {
			latestNode = &Day23Node{cup, nil, prev}
			prev.next = latestNode
			prev = latestNode
		}
	}

	first.prev = latestNode
	latestNode.next = first


	computeRounds(first, rounds, max)


	node1 := first
	for node1.i != 1 {
		node1 = node1.next
	}

	res := ""
	node := node1.next
	for node.i != 1 {
		res = res + strconv.Itoa(node.i)
		node = node.next
	}

	return res
}

func computeRounds(currCup * Day23Node, rounds int, max int) {

	for i := 0; i < rounds; i++ {
		three := currCup.next
		currCup.next = currCup.next.next.next.next
		dest := currCup.i - 1
		if dest == 0 {
			dest = max
		}
		for dest == three.i || dest == three.next.i || dest == three.next.next.i {
			dest--
			if dest == 0 {
				dest = max
			}
		}

		destNode := currCup
		for destNode.i != dest {
			destNode = destNode.prev
		}
		three.next.next.next = destNode.next
		destNode.next = three
		currCup = currCup.next
	}
}

type Day23Node struct {
	i int
	next *Day23Node
	prev *Day23Node
}

func ComputeDay23b(s string) int64 {
	onemil := 1000000
	tenmil := 10000000

	var first *Day23Node
	var prev *Day23Node
	var initFirst = false

	for _, c := range s {
		cup, err := strconv.Atoi(string(c))
		if err != nil {
			panic("Cannot interpret cup " + string(c))
		}
		if initFirst == false {
			first = &Day23Node{ cup, nil , nil}
			prev = first
			initFirst = true
		} else {
			newNode := &Day23Node{cup, nil, prev}
			prev.next = newNode
			prev = newNode
		}
	}
	var latestNode *Day23Node

	for i := 10; i <= onemil; i++ {
		latestNode = &Day23Node{i, nil, prev }
		prev.next = latestNode
		prev = latestNode
	}

	first.prev = latestNode
	latestNode.next = first

	computeRounds(first, tenmil, onemil)

	node1 := first
	for node1.i != 1 {
		node1 = node1.next
	}

	return int64(node1.next.i) * int64(node1.next.next.i)
}