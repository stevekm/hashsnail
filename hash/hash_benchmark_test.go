package hash

import (
	"fmt"
	// "log"
	"testing"
	"hashsnail/combinator"
)


// $ go test -v -bench=. hash/*
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
				combinator.CharSetDefault,                          // charSet
				"900150983cd24fb0d6963f7d28e17f72", // wantedHash
				false,                              // print
				4, // numThreads
			),
			want: "abc",
		},
		"hash_abcd": {
			finder: NewHashFinder(
				10000000*10000000,                  // numCombs
				4,                                  // maxSize
				0,                                  // minSize
				combinator.CharSetDefault,                          // charSet
				"e2fc714c4727ee9395f324cd2e7f331f", // wantedHash
				false,                              // print
				4, // numThreads
			),
			want: "abcd",
		},
	}

	for name, tc := range tests {
		b.Run(name, func(b *testing.B) {
			tc.finder.FindParallel() // result, _ :=
			fmt.Printf("Result found: %v\n", tc.finder.DescribeResults())
		})
	}

	// numCombs := 10000000*10000000
	// maxSize := -1
	// minSize := 0
	// charSet := combinator.CharSetDefault
	// wanted := "2510c39011c5be704182423e3a695e91" // h
	// print := false
	// numThreads := 2
	// finder := NewHashFinder(numCombs, maxSize, minSize, charSet, wanted, print, numThreads)
	// log.Printf("Starting finder:%v\n", finder.DescribeStart())
	// result, err := finder.FindParallel() // result, err :=
	// if err != nil {
	// 	log.Fatalf("%v", err)
	// 	// return err
	// }
	// log.Printf("Result found: %v %v\n", result, finder.DescribeResults())
}
