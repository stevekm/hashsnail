package hash

import (
	"fmt"
	// "log"
	"hashsnail/combinator"
	"testing"
)

// always store the result to a package level variable
// so the compiler cannot eliminate the Benchmark itself.
// https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
var result string

// $ go test -v -bench=. hash/*
// to run only benchmarks and no unit tests;
// $ go test -v -bench=. -run=^# hash/*
// $ go test -trace trace.out -memprofile mem.out -cpuprofile cpu.out -v -bench=. -run=^# hash/*
// https://go.dev/blog/pprof
func BenchmarkHash(b *testing.B) {
	// numCombs := -1 // this one deadlocks !!
	tests := map[string]struct {
		finder HashFinder
		want   string
	}{
		"hash_abc": {
			finder: NewHashFinder(
				10000000*10000000,                  // numCombs
				4,                                  // maxSize
				0,                                  // minSize
				combinator.CharSetDefault,          // charSet
				"900150983cd24fb0d6963f7d28e17f72", // wantedHash
				false,                              // print
				4,                                  // numThreads
			),
			want: "abc",
		},
		"hash_abcd": {
			finder: NewHashFinder(
				10000000*10000000,                  // numCombs
				4,                                  // maxSize
				0,                                  // minSize
				combinator.CharSetDefault,          // charSet
				"e2fc714c4727ee9395f324cd2e7f331f", // wantedHash
				false,                              // print
				4,                                  // numThreads
			),
			want: "abcd",
		},
		"hash_h": {
			finder: NewHashFinder(
				10000000*10000000,                  // numCombs
				4,                                  // maxSize
				0,                                  // minSize
				combinator.CharSetDefault,          // charSet
				"2510c39011c5be704182423e3a695e91", // wantedHash
				false,                              // print
				4,                                  // numThreads
			),
			want: "h",
		},
		"hash_hu": {
			finder: NewHashFinder(
				10000000*10000000,                  // numCombs
				4,                                  // maxSize
				0,                                  // minSize
				combinator.CharSetDefault,          // charSet
				"18bd9197cb1d833bc352f47535c00320", // wantedHash
				false,                              // print
				4,                                  // numThreads
			),
			want: "hu",
		},
		"hash_hun": {
			finder: NewHashFinder(
				10000000*10000000,                  // numCombs
				4,                                  // maxSize
				0,                                  // minSize
				combinator.CharSetDefault,          // charSet
				"fe1b3b54fde5b24bb40f22cdd621f5d0", // wantedHash
				false,                              // print
				4,                                  // numThreads
			),
			want: "hun",
		},
		"hash_hunt": {
			finder: NewHashFinder(
				10000000*10000000,                  // numCombs
				4,                                  // maxSize
				0,                                  // minSize
				combinator.CharSetDefault,          // charSet
				"bc9bf7bb6c4ab8d0daf374963110f4a7", // wantedHash
				false,                              // print
				4,                                  // numThreads
			),
			want: "hunt",
		},
		// These are too long to test for right now
		// "hash_hunte": {
		// 	finder: NewHashFinder(
		// 		10000000*10000000,                  // numCombs
		// 		4,                                  // maxSize
		// 		0,                                  // minSize
		// 		combinator.CharSetDefault,          // charSet
		// 		"9e3ae1b513b828922d4f691254bda0c1", // wantedHash
		// 		false,                              // print
		// 		4,                                  // numThreads
		// 	),
		// 	want: "hunte",
		// },
		// "hash_hunter": {
		// 	finder: NewHashFinder(
		// 		10000000*10000000,                  // numCombs
		// 		4,                                  // maxSize
		// 		0,                                  // minSize
		// 		combinator.CharSetDefault,          // charSet
		// 		"6b1b36cbb04b41490bfc0ab2bfa26f86", // wantedHash
		// 		false,                              // print
		// 		4,                                  // numThreads
		// 	),
		// 	want: "hunter",
		// },
		// "hash_hunter2": {
		// 	finder: NewHashFinder(
		// 		10000000*10000000,                  // numCombs
		// 		4,                                  // maxSize
		// 		0,                                  // minSize
		// 		combinator.CharSetDefault,          // charSet
		// 		"2ab96390c7dbe3439de74d0c9b0b1767", // wantedHash
		// 		false,                              // print
		// 		4,                                  // numThreads
		// 	),
		// 	want: "hunter2",
		// },
	}

	for name, tc := range tests {
		b.Run(name, func(b *testing.B) {
			r, _ := tc.finder.FindParallel()
			result = r
			fmt.Printf("Result found: %v\n", tc.finder.DescribeResults())
		})
	}
}
