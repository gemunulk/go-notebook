package main

import (
	"fmt"
	"testing"
)

// 1. Filename ends with = filename_test.go
// Start method name with Test = Test_methodname()
// CMD = go test -run Test_Something -v
func Test_Something(t *testing.T) {
	// test stuff here...
	Something(t)
}

func Something(t *testing.T) {
	// test stuff here...
	fmt.Println("GRRR")
}
