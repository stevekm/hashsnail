package main

import (
	"fmt"
	"testing"
	"github.com/google/go-cmp/cmp"
)


func TestFindAllFiles(t *testing.T) {
	t.Run("test_foo", func(t *testing.T) {
		fmt.Println("this is test")

		gen := NewPasswordGenerator()
		var got string = gen.foo()
		var want string = "foo"

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("got vs want mismatch (-want +got):\n%s", diff)
	}
	})

	t.Run("test_next", func(t *testing.T) {
		gen := NewPasswordGenerator()
		i := 0
		for i < 10 {
			fmt.Println(i)
			i += 1
			fmt.Println(gen.Next())
		}
	})
}