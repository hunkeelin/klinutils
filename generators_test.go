package klinutils

import (
	"fmt"
	"testing"
)

func TestGenToken(t *testing.T) {
	fmt.Println(Gentoken(5))
}
func ExampleGentoken() {
	fmt.Println(Gentoken(5))
	// Output: 73c31c7824
}
