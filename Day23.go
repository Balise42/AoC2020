package main

import (
	"fmt"
	"strconv"
)

func (day *Day) Day23a() {
	fmt.Printf("Part 1: %s\n", ComputeDay23a("193467258", 100))
}

func (day *Day) Day23b() {
	fmt.Printf("Part 2: %d\n", ComputeDay23b("193467258"))
}


func ComputeDay23a(s string, rounds int) string {
	cups := make([]int, len(s))
	for i, c := range s {
		cup, err := strconv.Atoi(string(c))
		if err != nil {
			panic("Cannot interpret cup " + string(c))
		}
		cups[i] = cup
	}
	currCupIndex := -1
	for i := 0; i < rounds; i++ {
		cups, currCupIndex = runRoundPart1(cups, (currCupIndex + 1) % len(cups))
	}
	cupstr := ""

	pos1 := -1
	for i, v := range cups {
		if v == 1 {
			pos1 = i
			break
		}
	}
	if pos1 == -1 {
		panic("Could not find cup with index 1")
	}

	for i := (pos1 + 1) % len(cups); i != pos1; i = (i+1) % len(cups) {
		cupstr += strconv.Itoa(cups[i])
	}
	return cupstr
}

func runRoundPart1(state []int, currCupIndex int) ([]int, int) {
	cups := state

	three := make([]int, 3)
	currCup := cups[currCupIndex]
	if currCupIndex + 4 < len(cups) {
		three[0], three[1], three[2] = cups[currCupIndex + 1], cups[currCupIndex + 2], cups[currCupIndex + 3]
		cups = append(cups[:currCupIndex + 1], cups[currCupIndex + 4:]...)
	} else if currCupIndex + 3 < len(cups) {
		three[0], three[1], three[2] = cups[currCupIndex + 1], cups[currCupIndex + 2], cups[currCupIndex + 3]
		cups = cups[:currCupIndex + 1]
	} else if currCupIndex + 2 < len(cups) {
		three[0], three[1], three[2] = cups[currCupIndex + 1], cups[currCupIndex + 2], cups[0]
		cups = cups[1:currCupIndex + 1]
	} else if currCupIndex + 1 < len(cups) {
		three[0], three[1], three[2] = cups[currCupIndex + 1], cups[0], cups[1]
		cups = cups[2:currCupIndex + 1]
	} else {
		copy(three, cups[:3])
		cups = cups[3:]
	}

	dest := currCup - 1
	if dest < 1 {
		dest = 9
	}
	for dest == three[0] || dest == three[1] || dest == three[2] {
		dest--
		if dest < 1 {
			dest = 9
		}
	}

	destIndex := -1
	for i, v := range cups {
		if v == dest {
			destIndex = i
			break
		}
	}
	if destIndex == -1 {
		panic("Could not find insertion index for " + strconv.Itoa(dest))
	}

	if destIndex < len(cups) - 1 {
		cups = append(append(cups[:destIndex + 1], append(three, cups[destIndex + 1:]...)...))
	} else {
		cups = append(cups, three...)
	}

	if cups[currCupIndex] != currCup {
		for i, v := range cups {
			if v == currCup {
				currCupIndex = i
				break
			}
		}
	}

	return cups, currCupIndex
}

type Day23Node struct {
	i int
	next *Day23Node
	prev *Day23Node
}

// fine, let's implement a ringlist then >_<
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

	currCup := first

	for i := 0; i < tenmil; i++ {
		three := currCup.next
		currCup.next = currCup.next.next.next.next
		dest := currCup.i - 1
		if dest == 0 {
			dest = onemil
		}
		for dest == three.i || dest == three.next.i || dest == three.next.next.i {
			dest--
			if dest == 0 {
				dest = onemil
			}
		}

		destNode := currCup
		for destNode.i != dest {
			destNode = destNode.prev
		}
		three.next.next.next = destNode.next
		destNode.next = three
	}

	node1 := first
	for node1.i != 1 {
		node1 = node1.next
	}

	return int64(node1.next.i) * int64(node1.next.next.i)
}