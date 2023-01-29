package combinator

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

//
//
// Benchmarks go here
// $ go test -bench=. -v combinator/*
// https://pkg.go.dev/testing#hdr-Benchmarks
// https://pkg.go.dev/cmd/go#hdr-Testing_flags
// https://blog.logrocket.com/benchmarking-golang-improve-function-performance/
// https://stackoverflow.com/questions/53322925/is-it-possible-to-limit-iterations-b-n-in-go-benchmarking
// $ go test -v -bench=BenchmarkCombinator2 combinator/* | grep '\['
// https://stackoverflow.com/questions/16161142/how-to-test-only-one-benchmark-function
//

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


func BenchmarkCombinator(b *testing.B) {
	// benchmark for the main character combination method
	for _, v := range benchTable {
		b.Run(fmt.Sprintf("combinator_%d", v.input), func(b *testing.B) {
			DemoCombinator(v.input, b.Name())
		})
	}
}

func Benchmark2Combinator(b *testing.B) {
	// benchmark experimental new methods for combinator
	charSet := strings.Split(CharSetDefault, "") // set := []string{"a", "b", "c"}
	numChars := len(charSet)
	numCombs := 0
	startTime := time.Now()
	for len := 1; len <= numChars; len++ {
		iMax := numChars - len + 1
		for i := 0; i < iMax; i++ {
			for j := i + 1; j <= numChars; j++ {
				// fmt.Printf(">>> len:%v iMax:%v i:%v j:%v\n", len, iMax, i, j)
				// fmt.Printf(">>> %v\n", strings.Join(charSet[i:j], ""))
				numCombs++
				continue
			}
		}
	}
	stopTime := time.Now()
	PrintResult(b.Name(), uint(numCombs), startTime, stopTime)
}

func iterate(n int) {
	// raw iterator to see how fast we can increment an iterator
	// with no other operations going on
	for i := 0; i < n; i++ {
		continue
	}
}
func BenchmarkRawIterator(b *testing.B) {
	for _, v := range benchTable {
		b.Run(fmt.Sprintf("raw_iterator_%d", v.input), func(b *testing.B) {
			DemoIterate(v.input, b.Name())
		})
	}

}

// Benchmark results
// M1 MacBook Air
// https://www.anandtech.com/show/16252/mac-mini-apple-m1-tested;
// 3.2GHz , 8x 16-bit memory channels LPDDR4X-4266 memory, peak of 68.25GB/s memory bandwidth
// need to look at more L1/L2 cache speed details too
// [BenchmarkCombinator/combinator_10000000] numCombs: 10000000, elapsed 1.265366083s, rate:7.9M/s
// [BenchmarkCombinator/combinator_20000000] numCombs: 20000000, elapsed 2.500602541s, rate:8.0M/s
// [BenchmarkCombinator/combinator_40000000] numCombs: 40000000, elapsed 5.019067959s, rate:8.0M/s
// [BenchmarkCombinator/combinator_80000000] numCombs: 80000000, elapsed 10.108567083s, rate:7.9M/s
// [BenchmarkCombinator/combinator_100000000] numCombs: 100000000, elapsed 13.558417s, rate:7.4M/s
// [BenchmarkCombinator/combinator_150000000] numCombs: 150000000, elapsed 22.210058625s, rate:6.8M/s
// [BenchmarkCombinator2] numCombs: 290320, elapsed 128.209µs, rate:2264.4M/s
// [BenchmarkCombinator2] numCombs: 290320, elapsed 133.75µs, rate:2170.6M/s
// [BenchmarkCombinator2] numCombs: 290320, elapsed 131.875µs, rate:2201.5M/s
// [BenchmarkCombinator2] numCombs: 290320, elapsed 129.833µs, rate:2236.1M/s
// [BenchmarkCombinator2] numCombs: 290320, elapsed 129.791µs, rate:2236.8M/s
// [BenchmarkCombinator2] numCombs: 290320, elapsed 133.333µs, rate:2177.4M/s
// [BenchmarkRawIterator/raw_iterator_10000000] numCombs: 10000000, elapsed 3.3645ms, rate:2972.2M/s
// [BenchmarkRawIterator/raw_iterator_10000000] numCombs: 10000000, elapsed 3.197708ms, rate:3127.2M/s
// [BenchmarkRawIterator/raw_iterator_10000000] numCombs: 10000000, elapsed 3.177209ms, rate:3147.4M/s
// [BenchmarkRawIterator/raw_iterator_10000000] numCombs: 10000000, elapsed 3.173084ms, rate:3151.5M/s
// [BenchmarkRawIterator/raw_iterator_10000000] numCombs: 10000000, elapsed 3.1965ms, rate:3128.4M/s
// [BenchmarkRawIterator/raw_iterator_10000000] numCombs: 10000000, elapsed 3.235458ms, rate:3090.8M/s
// [BenchmarkRawIterator/raw_iterator_20000000] numCombs: 20000000, elapsed 6.356209ms, rate:3146.5M/s
// [BenchmarkRawIterator/raw_iterator_20000000] numCombs: 20000000, elapsed 6.28575ms, rate:3181.8M/s

// Ryzen 3950X + 64GB DDR4 3200MHz
// [BenchmarkCombinator/combinator_10000000] numCombs: 10000000, elapsed 3.675008393s, rate:2.7M/s
// [BenchmarkCombinator/combinator_20000000] numCombs: 20000000, elapsed 7.834336221s, rate:2.6M/s
// [BenchmarkCombinator/combinator_40000000] numCombs: 40000000, elapsed 15.541499958s, rate:2.6M/s
// [BenchmarkCombinator/combinator_80000000] numCombs: 80000000, elapsed 31.235968711s, rate:2.6M/s
// [BenchmarkCombinator/combinator_100000000] numCombs: 100000000, elapsed 41.774112917s, rate:2.4M/s
// [BenchmarkCombinator/combinator_150000000] numCombs: 150000000, elapsed 1m8.917689113s, rate:2.2M/s
// [BenchmarkCombinator2] numCombs: 290320, elapsed 249.41µs, rate:1164.0M/s
// [BenchmarkCombinator2] numCombs: 290320, elapsed 213.003µs, rate:1363.0M/s
// [BenchmarkCombinator2] numCombs: 290320, elapsed 265.399µs, rate:1093.9M/s
// [BenchmarkCombinator2] numCombs: 290320, elapsed 214.696µs, rate:1352.2M/s
// [BenchmarkCombinator2] numCombs: 290320, elapsed 185.031µs, rate:1569.0M/s
// [BenchmarkCombinator2] numCombs: 290320, elapsed 191.634µs, rate:1515.0M/s
// [BenchmarkRawIterator/raw_iterator_10000000] numCombs: 10000000, elapsed 9.781939ms, rate:1022.3M/s
// [BenchmarkRawIterator/raw_iterator_10000000] numCombs: 10000000, elapsed 9.127493ms, rate:1095.6M/s
// [BenchmarkRawIterator/raw_iterator_10000000] numCombs: 10000000, elapsed 9.119298ms, rate:1096.6M/s
// [BenchmarkRawIterator/raw_iterator_10000000] numCombs: 10000000, elapsed 9.12055ms, rate:1096.4M/s
// [BenchmarkRawIterator/raw_iterator_10000000] numCombs: 10000000, elapsed 9.119299ms, rate:1096.6M/s
// [BenchmarkRawIterator/raw_iterator_10000000] numCombs: 10000000, elapsed 9.119026ms, rate:1096.6M/s
