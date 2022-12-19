package main

import (
	"fmt"
	"strings"
	"./combinator/comb"
)

func main() {
	state := combinator.NewState()

	for i:= 0; i < 5000; i++ {
		comb := state.Next()
		fmt.Printf("%v\t%v\n", strings.Join(comb, ""), state.Indexes)
	}
}