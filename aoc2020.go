package main

import (
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Day struct {
	input string
}

func main() {
	if len(os.Args) < 2 {
		panic("Need day number for execution.")
	}

	arg := os.Args[1]
	var day Day

	inputpath := "D:/Home/projects/adventcalendar/2020/inputs/day" + arg + ".dat"

	bytes, err := ioutil.ReadFile(inputpath)
	if err != nil {
		log.Println("No input file found, continuing assuming the input is provided in day functions.")
	} else {
		day.input = string(bytes)
	}

	reflect.ValueOf(&day).MethodByName("Day"+arg+"a").Call([]reflect.Value{})
	reflect.ValueOf(&day).MethodByName("Day"+arg+"b").Call([]reflect.Value{})
}

func createListAsInts(input string) []int {
	intsAsString := strings.Split(input, "\n")
	return convertStringListToInts(intsAsString)
}

func convertStringListToInts(input []string) []int {
	res := make([]int, len(input))
	var err error
	for i, v := range input {
		res[i], err = strconv.Atoi(v)
		if err != nil {
			panic("Could not convert input " + v + " to number")
		}
	}
	return res
}
