package main

import (
	"fmt"
	"io/ioutil"
)

func a(input string) int {
	floor := 0

	for _, char := range input {
		if char == '(' {
			floor = floor + 1
		} else {
			floor = floor - 1
		}
	}

	return floor
}

func b(input string) int {
	floor := 0
	enteredBasementAt := 0

	for pos, char := range input {
		if char == '(' {
			floor = floor + 1
		} else {
			floor = floor - 1
		}

		if floor == -1 {
			enteredBasementAt = pos + 1
			break
		}
	}

	return enteredBasementAt
}

func main() {
	data, _ := ioutil.ReadFile("input_a.txt")

	fmt.Println("Assignment A: ")
	fmt.Println(a(string(data)))

	fmt.Println("Assignment B: ")
	fmt.Println(b(string(data)))
}
