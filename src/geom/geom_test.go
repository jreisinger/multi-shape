package geom_test

import (
	"fmt"
	"geom"
	"testing"
)

func TestCircleArea(t *testing.T) {
	c := &geom.Circle{0, 0, 5}
	v := fmt.Sprintf("%.2f", c.Area())
	if v != "78.54" {
		t.Error("Expected 78.54, got ", v)
	}

	c = &geom.Circle{0, 0, 0}
	v = fmt.Sprintf("%.2f", c.Area())
	if v != "0.00" {
		t.Error("Expected 0.00, got ", v)
	}
}
