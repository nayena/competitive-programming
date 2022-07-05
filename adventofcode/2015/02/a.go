package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func a(input string) int {
	lines := strings.Split(input, "\n")
	total := 0

	for _, line := range lines {
		length, width, height := parseBoxDimensions(line)

		surfaceArea := (2 * length * width) + (2 * width * height) + (2 * height * length)

		sides := []int{length, width, height}
		sort.Ints(sides)

		extra := sides[0] * sides[1]
		total = total + surfaceArea + extra
	}

	return total
}

func b(input string) int {
	lines := strings.Split(input, "\n")
	total := 0

	for _, line := range lines {
		length, width, height := parseBoxDimensions(line)

		volume := length * width * height

		sides := []int{length, width, height}
		sort.Ints(sides)

		extra := 2*sides[0] + 2*sides[1]
		total = total + volume + extra
	}

	return total
}

func parseBoxDimensions(input string) (int, int, int) {
	dimensions := strings.Split(input, "x")

	length, _ := strconv.Atoi(dimensions[0])
	width, _ := strconv.Atoi(dimensions[1])
	height, _ := strconv.Atoi(dimensions[2])

	return length, width, height
}

func main() {
	data, _ := ioutil.ReadFile("input_a.txt")

	fmt.Println("Assignment A: ")
	fmt.Println(a(string(data)))

	fmt.Println("Assignment B: ")
	fmt.Println(b(string(data)))
}
