package combinator

import (
	"fmt" // fmt.Printf("%v %v\n", want, got)
	"github.com/google/go-cmp/cmp"
	"testing"
	"time"
	"strings"
)

func TestCombinator(t *testing.T) {
	t.Run("test_combinator1", func(t *testing.T) {
		state := NewState("abc", 0)
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







var benchTable = []struct {
	input int
}{
	// {input: 100}, // these are too fast
	// {input: 1000},
	// {input: 10000},
	// {input: 100000},
	// {input: 1000000},
	{input: 10000000},
	{input: 20000000},
	{input: 40000000},
	{input: 80000000},
	{input: 100000000},
	{input: 150000000},
	// {input: 200000000}, // takes too long
}


func printResult(numCombs uint, startTime time.Time, stopTime time.Time){
	elapsed := stopTime.Sub(startTime)
	rate := float64(numCombs) / elapsed.Seconds()
	fmt.Printf(">>> numCombs: %v, elapsed %v, rate:%.1fM/s\n",
	numCombs, elapsed, rate/1000000)
}

func iterate(n int){
	for i := 0; i < n; i++ {
		continue
	}
}

// $ go test -bench=. -v combinator/*
// https://pkg.go.dev/testing#hdr-Benchmarks
// https://blog.logrocket.com/benchmarking-golang-improve-function-performance/
func BenchmarkCombinator(b *testing.B) {

	charSet := strings.Split(CharSetDefault, "") // set := []string{"a", "b", "c"}
	numChars := len(charSet)
	for len := 1; len <= numChars; len++ {
		for i := 0; i < numChars-len+1; i++ {
			for j := i + 1; j <= numChars; j++ {
				fmt.Printf(">>> %v\n", strings.Join(charSet[i:j], ""))
			}
		}

	}


	for _, v := range benchTable {
		b.Run(fmt.Sprintf("combinator_%d", v.input), func(b *testing.B) {
			numCombs := v.input
			state := NewState(CharSetDefault, 0)
			startTime := time.Now()
			for i := 0; i < numCombs; i++ {
				state.Next()
			}
			stopTime := time.Now()
			printResult(state.NumGenerated, startTime, stopTime)
		})
	}

	for _, v := range benchTable {
		b.Run(fmt.Sprintf("raw_iterator_%d", v.input), func(b *testing.B) {
			numCombs := v.input
			startTime := time.Now()
			iterate(numCombs)
			stopTime := time.Now()
			printResult(uint(numCombs), startTime, stopTime)
		})
	}


}
