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
}

func (f *HashFinder) Find() {
	for i:= 0; i < f.NumCombs; i++ {
		comb := f.combinator.Next()
		if len(comb) > f.MaxSize {
			break
		}
		hash := GetHash(comb)
		fmt.Printf("%v %v\n", comb, hash)
		if hash == f.Wanted {
			fmt.Printf(">>> FOUND value '%v' for hash %v\n", comb, hash)
			break
		}
	}
}

func NewHashFinder(numCombs int, maxSize int, minSize int, charSet string, wanted string) HashFinder {
	comb := combinator.NewState(charSet, minSize)
	finder := HashFinder {
		NumCombs: numCombs,
		MaxSize: maxSize,
		Wanted: wanted,
		combinator: comb,
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
	numCombs := 10000000
	maxSize := 3
	minSize := 0
	charSet := combinator.CharSetDefault
	// a 0cc175b9c0f1b6a831c399e269772661
	wanted := "0cc175b9c0f1b6a831c399e269772661"

	finder := NewHashFinder(numCombs, maxSize, minSize, charSet, wanted)
	finder.Find()
}