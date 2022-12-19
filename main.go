package main

import (
	"fmt"
	"strings"
)

func main() {
	state := NewState()

	for i:= 0; i < 5000; i++ {
		comb := state.Next()
		fmt.Printf("%v\t%v\n", strings.Join(comb, ""), state.Indexes)
	}
}