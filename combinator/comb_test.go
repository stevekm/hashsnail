package combinator

import (
	// "fmt" // fmt.Printf("%v %v\n", want, got)
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestCombinator(t *testing.T) {
	t.Run("test_combinator1", func(t *testing.T) {
		minSize := 0
		iterStart := 0
		iterStep := 1
		state := NewState("abc", minSize, iterStart, iterStep)
		got := []string{}
		for i := 0; i < 25; i++ {
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

func TestCombinatorStep(t *testing.T) {
	// test cases for using combinator with different start and step sizes
	tests := map[string]struct {
		state         State
		numIterations int
		want          []string
		got           []string
	}{
		"comb_0start_2step": {
			state: NewState(
				"abc", // charSet
				0,     // minSize
				0,     // iterStart
				2),    // iterStep
			numIterations: 6,
			want: []string{
				"a", "c", "ba",
				"ab", "cb", "bc",
			},
		},
		"comb_1start_2step": {
			state: NewState(
				"abc", // charSet
				0,     // minSize
				1,     // iterStart
				2),    // iterStep
			numIterations: 6,
			want: []string{
				"b", "aa", "ca",
				"bb", "ac", "cc",
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := []string{}
			for i := 0; i < tc.numIterations; i++ {
				comb := tc.state.Next()
				got = append(got, comb)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("got vs want mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
