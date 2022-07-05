package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
)

func mine(base string, req string) int {
	val := 0

	for {
		check := []byte(fmt.Sprintf("%s%d", base, val))
		hash := fmt.Sprintf("%x", md5.Sum(check))
		if hash[0:len(req)] == req {
			break
		}
		val++
	}
	return val
}

func a(input string) int {
	return mine(input, "00000")
}

func b(input string) int {
	return mine(input, "000000")
}

func main() {
	data, _ := ioutil.ReadFile("input_a.txt")

	fmt.Println("Assignment A: ")
	fmt.Println(a(string(data)))

	fmt.Println("Assignment B: ")
	fmt.Println(b(string(data)))
}
