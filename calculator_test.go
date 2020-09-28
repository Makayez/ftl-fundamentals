package calculator_test

import (
	"calculator"
	"testing"
)

type testCase struct {
	a, b        float64
	want        float64
	name        string
	errExpected bool
}

func TestAdd(t *testing.T) {

	testCases := []testCase{
		{a: 2, b: 2, want: 4, name: "positive addition", errExpected: false},
		{a: -1, b: -1, want: -2, name: "negative addition", errExpected: false},
		{a: 5, b: 0, want: 5, name: "zero addition", errExpected: false},
	}

	t.Parallel()
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Test:%s | want %f, got %f", tc.name, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {

	testCases := []testCase{
		{a: 12, b: 2, want: 10, name: "positive subtraction", errExpected: false},
		{a: -16, b: 1, want: -17, name: "negative subtraction", errExpected: false},
		{a: 5, b: 0, want: 5, name: "zero subtraction", errExpected: false},
	}

	t.Parallel()
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Test:%s | want %f, got %f", tc.name, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {

	testCases := []testCase{
		{a: 12, b: 2, want: 24, name: "positive multiplication", errExpected: false},
		{a: 16, b: -1, want: -16, name: "negative multiplication", errExpected: false},
		{a: 5, b: 0, want: 0, name: "zero multiplication", errExpected: false},
	}

	t.Parallel()
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Test:%s | want %f, got %f", tc.name, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {

	testCases := []testCase{
		{a: 12, b: 2, want: 6, name: "positive division", errExpected: false},
		{a: 16, b: -1, want: -16, name: "negative division", errExpected: false},
		{a: 5, b: 0, want: -1, name: "invalid divide by zero division", errExpected: true},
	}

	t.Parallel()
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := err != nil
		if tc.errExpected != (err != nil) {
			t.Fatalf("Test:%s | unexpected error status: %v", tc.name, errReceived)
		}
		if !tc.errExpected && tc.want != got {
			t.Errorf("Test:%s | want %f, got %f", tc.name, tc.want, got)
		}
	}
}
