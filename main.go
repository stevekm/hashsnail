package main
// go run .

import (
	"fmt"
	"hashsnail/combinator"
)

func main() {
	state := combinator.NewState()

	for i:= 0; i < 500 * 1000 * 1000; i++ {
		comb := state.Next()
		fmt.Printf("%v\n", comb)
	}
}