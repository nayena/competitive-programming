package main

import "testing"

type Test struct {
	Input    string
	Expected int
}

func Test01(t *testing.T) {
	allTests := []Test{
		{
			Input:    "2x3x4",
			Expected: 58,
		},
		{
			Input:    "1x1x10",
			Expected: 43,
		},
	}

	RunTests(allTests, a, t)
}

func Test02(t *testing.T) {
	allTests := []Test{
		{
			Input:    "2x3x4",
			Expected: 34,
		},
		{
			Input:    "1x1x10",
			Expected: 14,
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
