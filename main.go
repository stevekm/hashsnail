package main
// go run .

import (
	"fmt"
	"hashsnail/combinator"
	"crypto/md5"
	// "crypto/sha1"
	// "crypto/sha256"
	"encoding/hex"
)

type HashFinder struct {
	NumCombs int
	MaxSize int // size of string to hash
	combinator combinator.State
	Wanted string // the hash we want to match
	Print bool
}

func (f *HashFinder) Find() {
	for i:= 0; i < f.NumCombs; i++ {
		comb := f.combinator.Next()
		if len(comb) > f.MaxSize {
			break
		}
		hash := GetHash(comb)
		if f.Print {
			fmt.Printf("%v %v\n", comb, hash)
		}
		if hash == f.Wanted {
			fmt.Printf(">>> FOUND value '%v' for hash %v\n", comb, hash)
			break
		}
	}
}

func NewHashFinder(numCombs int, maxSize int, minSize int, charSet string, wanted string, print bool) HashFinder {
	comb := combinator.NewState(charSet, minSize)
	finder := HashFinder {
		NumCombs: numCombs,
		MaxSize: maxSize,
		Wanted: wanted,
		combinator: comb,
		Print: print,
	}
	return finder
}

func GetHash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func GetAllHash(texts []string) []string {
	hashes := []string{}
	for _, t := range texts {
		hashes = append(hashes, GetHash(t))
	}
	return hashes
}

func PrintHashCombs(numCombs int, maxSize int){
	state := combinator.NewState(combinator.CharSetDefault, 0)
	for i:= 0; i < numCombs; i++ {
		comb := state.Next()
		if len(comb) > maxSize {
			break
		}
		fmt.Printf("%v %v\n", comb, GetHash(comb))
	}

}

func main() {
	numCombs := 10000000 * 10000000
	maxSize := 5
	minSize := 0
	charSet := combinator.CharSetDefault
	print := false
	// a 0cc175b9c0f1b6a831c399e269772661
	// abc 900150983cd24fb0d6963f7d28e17f72 2s
	// abcd e2fc714c4727ee9395f324cd2e7f331f 2:36
	// abcde ab56b4d92b40713acc5af89985d4b786 1:40:50
	wanted := "ab56b4d92b40713acc5af89985d4b786"

	finder := NewHashFinder(numCombs, maxSize, minSize, charSet, wanted, print)
	finder.Find()
}