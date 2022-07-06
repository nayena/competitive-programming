package main

import "testing"

type Test struct {
	Input    string
	Expected int
}

func Test01(t *testing.T) {
	compare(t, removeDuplicates("abcd", 2), "abcd")
	compare(t, removeDuplicates("deeedbbcccbdaa", 3), "aa")
	compare(t, removeDuplicates("pbbcggttciiippooaais", 2), "ps")
}

func compare(t *testing.T, actual string, expected string) {
	if actual != expected {
		t.Errorf("expected '%s' but got '%s'", expected, actual)
	}
}
