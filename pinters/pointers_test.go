package main

import (
	"fmt"
	"testing"
)

func changeA(x *int) {
	(*x) = 10
}

// Test exec cmd = go test -run Test_changeA
func Test_changeA(t *testing.T) {
	a := 1
	fmt.Println(a)
	changeA(&a)
	fmt.Println(a)
}
