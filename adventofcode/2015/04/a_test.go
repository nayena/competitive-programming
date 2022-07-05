package main

import "testing"

type Test struct {
	Input    string
	Expected int
}

func Test01(t *testing.T) {
	allTests := []Test{
		{
			Input:    "abcdef",
			Expected: 609043,
		},
		{
			Input:    "pqrstuv",
			Expected: 1048970,
		},
	}
	RunTests(allTests, a, t)
}

func RunTests(tests []Test, f func(string) int, t *testing.T) {
	for _, test := range tests {
		output := f(test.Input)

		if output != test.Expected {
			t.Errorf("expected '%d' but got '%d'", test.Expected, output)
		}
	}
}
