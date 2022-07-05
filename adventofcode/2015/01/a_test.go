package main

import "testing"

type Test struct {
	Input    string
	Expected int
}

func Test01(t *testing.T) {

	allTests := []Test{
		{
			Input:    "(())",
			Expected: 0,
		},
		{
			Input:    "(())",
			Expected: 0,
		},
		{
			Input:    "(((",
			Expected: 3,
		},
		{
			Input:    "(()(()(",
			Expected: 3,
		},
		{
			Input:    "))(((((",
			Expected: 3,
		},
		{
			Input:    "())",
			Expected: -1,
		},
		{
			Input:    "))(",
			Expected: -1,
		},
		{
			Input:    ")))",
			Expected: -3,
		},
		{
			Input:    ")())())",
			Expected: -3,
		},
	}

	for _, test := range allTests {
		output := a(test.Input)

		if output != test.Expected {
			t.Errorf("expected '%d' but got '%d'", test.Expected, output)
		}
	}
}

func Test02(t *testing.T) {

	allTests := []Test{
		{
			Input:    ")",
			Expected: 1,
		},
		{
			Input:    "()())",
			Expected: 5,
		},
	}

	for _, test := range allTests {
		output := b(test.Input)

		if output != test.Expected {
			t.Errorf("expected '%d' but got '%d'", test.Expected, output)
		}
	}
}
