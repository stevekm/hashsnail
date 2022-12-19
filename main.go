package main
// go run .

// comparison; <4s on M1 MacBook Air
// $ ./hashcat -m 0 -a 3 ab56b4d92b40713acc5af89985d4b786

import (
	"hashsnail/hash"
	"hashsnail/combinator"
)

func main() {
	numCombs := 10000000 * 10000000 // a big number
	maxSize := 5
	minSize := 0
	charSet := combinator.CharSetDefault
	print := false
	// wanted := "0cc175b9c0f1b6a831c399e269772661" // a 0.225s
	// wanted := "900150983cd24fb0d6963f7d28e17f72" // abc 2s
	wanted := "e2fc714c4727ee9395f324cd2e7f331f" // abcd 2:36
	// wanted := "ab56b4d92b40713acc5af89985d4b786" // abcde 1:40:50

	finder := hash.NewHashFinder(numCombs, maxSize, minSize, charSet, wanted, print)
	finder.Find()
}