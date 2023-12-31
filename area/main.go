package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

 func main() {
	s := square{10}
	t := triangle{10, 5}

	fmt.Println("area of square: ", s.getArea())
	fmt.Println("area of triangle: ", t.getArea())

	printArea(s)
	printArea(t)
	
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height

}
func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
