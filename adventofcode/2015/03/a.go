package main

import (
	"fmt"
	"io/ioutil"
)

func a(input string) int {
	x, y := 0, 0

	town := make(map[string]int)
	town["0x0"] = 1

	for _, instruction := range input {

		if instruction == '^' {
			y++
		}

		if instruction == 'v' {
			y--
		}

		if instruction == '>' {
			x++
		}

		if instruction == '<' {
			x--
		}

		house := fmt.Sprintf("%dx%d", x, y)

		town[house] = town[house] + 1
	}

	return len(town)
}

type Santa struct {
	X int
	Y int
}

func b(input string) int {
	var santas [2]Santa

	town := make(map[string]int)
	town["0x0"] = 1

	santas[0].X = 0
	santas[0].Y = 0

	santas[1].X = 0
	santas[1].Y = 0

	active := 0

	for _, instruction := range input {

		if active == 0 {
			active = 1
		} else {
			active = 0
		}

		if instruction == '^' {
			santas[active].Y++
		}

		if instruction == 'v' {
			santas[active].Y--
		}

		if instruction == '>' {
			santas[active].X++
		}

		if instruction == '<' {
			santas[active].X--
		}

		house := fmt.Sprintf("%dx%d", santas[active].X, santas[active].Y)

		town[house]++
	}

	return len(town)
}

func main() {
	data, _ := ioutil.ReadFile("input_a.txt")

	fmt.Println("Assignment A: ")
	fmt.Println(a(string(data)))

	fmt.Println("Assignment B: ")
	fmt.Println(b(string(data)))
}
