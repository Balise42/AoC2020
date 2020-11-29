package main

import (
	"io/ioutil"
	"log"
	"os"
	"reflect"
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
	println(inputpath)

	bytes, err := ioutil.ReadFile(inputpath)
	if err != nil {
		log.Println("No input file found, continuing assuming the input is provided in day functions.")
	} else {
		day.input = string(bytes)
	}

	reflect.ValueOf(&day).MethodByName("Day"+arg+"a").Call([]reflect.Value{})
	reflect.ValueOf(&day).MethodByName("Day"+arg+"b").Call([]reflect.Value{})
}
