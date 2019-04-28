package main

import (
	"fmt"
	"testing"
)

type animal interface {
	eat()
	sleep()
	drink()
}

type cat struct {
	name   string
	age    int
	weight int
}

// Try (c cat) and (c *cat)
// func (c cat) eat(food int) int {
func (c *cat) eat(food int) int {
	fmt.Println(fmt.Sprintf(c.name+" eats %d", food))
	c.weight += food // (*c).weight += food
	return c.weight  // return (*c).weight
}

func (c cat) sleep() {
	fmt.Println(c.name + " ZzzZZ")
}

func (c cat) drink() {
	fmt.Println(c.name + " Drinks")
}

func Test_cat(t *testing.T) {
	c := cat{name: "Oscar", age: 4, weight: 5}
	fmt.Println(c)
	c.eat(2) // (&c).eat(2)
	c.sleep()
	c.drink()
	fmt.Println(c.weight)

	p := &c
	p.eat(3)  //(*p).eat()
	p.sleep() // (*p).sleep()
	c.drink()

	z := p.weight // (*p).weight
	fmt.Println(z)

}
