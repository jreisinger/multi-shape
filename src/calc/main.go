package main

import (
	"fmt"
	"geom"
)

func main() {
	r := &geom.Rectangle{0, 0, 10, 10}
	c := &geom.Circle{0, 0, 5}
	multiShape := geom.MultiShape{Shapes: []geom.Shape{r, c}}

	fmt.Printf("Rectangle area:\t%.2f\n", r.Area())
	fmt.Printf("Circle area:\t%.2f\n", c.Area())
	fmt.Printf("Total area:\t%.2f\n", multiShape.Area())
}
