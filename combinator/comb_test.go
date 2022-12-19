package combinator

import (
	// "fmt" // fmt.Printf("%v %v\n", want, got)
	"testing"
	"github.com/google/go-cmp/cmp"
)

func TestCombinator(t *testing.T) {
	t.Run("test_combinator1", func(t *testing.T) {
		state := NewState("abc", 0)
		got := []string{}
		for i:= 0; i < 25; i++ {
			comb := state.Next()
			got = append(got, comb)
		}
		want := []string{
			"a", "b", "c",
			"aa", "ba", "ca",
			"ab", "bb", "cb",
			"ac", "bc", "cc",
			"aaa", "baa", "caa",
			"aba", "bba", "cba",
			"aca", "bca", "cca",
			"aab", "bab", "cab",
			"abb",
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("got vs want mismatch (-want +got):\n%s", diff)
		}
	})
}