package main

import "testing"

type Test struct {
	Input    string
	Expected int
}

func Test01(t *testing.T) {
	allTests := []Test{
		{
			Input:    ">",
			Expected: 2,
		},
		{
			Input:    "^>v<",
			Expected: 4,
		},
		{
			Input:    "^v^v^v^v^v",
			Expected: 2,
		},
	}
	RunTests(allTests, a, t)
}

func Test02(t *testing.T) {
	allTests := []Test{
		{
			Input:    "^v",
			Expected: 3,
		},
		{
			Input:    "^>v<",
			Expected: 3,
		},
		{
			Input:    "^v^v^v^v^v",
			Expected: 11,
		},
	}
	RunTests(allTests, b, t)
}

func RunTests(tests []Test, f func(string) int, t *testing.T) {
	for _, test := range tests {
		output := f(test.Input)

		if output != test.Expected {
			t.Errorf("expected '%d' but got '%d'", test.Expected, output)
		}
	}
}
