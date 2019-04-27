package main

import (
	"fmt"
	"math"
	"testing"
)

//https://gobyexample.com/interfaces

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// cmd = go test -run Test_measure -v
func Test_measure(t *testing.T) {
	r := rect{width: 3, height: 4}
	measure(r)

	c := circle{radius: 3}
	measure(c)
}
