package geom_test

import (
	"fmt"
	"geom"
	"testing"
)

type testpair struct {
	input    []float64
	expected string
}

var tests = []testpair{
	{[]float64{0, 0, 0}, "0.00"},
	{[]float64{0, 0, 5}, "78.54"},
	{[]float64{1, 0, 5}, "78.54"},
	{[]float64{0, 1, 5}, "78.54"},
	{[]float64{1, 1, 5}, "78.54"},
}

func TestCircleArea(t *testing.T) {
	for _, pair := range tests {
		c := &geom.Circle{pair.input[0], pair.input[1], pair.input[2]}
		out := fmt.Sprintf("%.2f", c.Area())
		if out != pair.expected {
			t.Error("Expected 0.00, got ", out)
		}
	}
}
