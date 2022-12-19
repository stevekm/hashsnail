package hash
import (
	"fmt"
	// "errors" //errors.New("No value found for hash")
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

func (f *HashFinder) Find() (string, error) {
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
			return comb, nil
			break
		}
	}
	return "", fmt.Errorf("No value found for hash %v", f.Wanted)
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