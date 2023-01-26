package combinator

import (
	// "log"
	"sort"
	"strings"
	// "time"
)

const CharSetDefault = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890 !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

type State struct {
	indexes      []int    // keeps track of which characters should be returned
	Chars        []string // character set to build combinations from
	NumGenerated uint     // number of combinations generated
}

func (s *State) Get() string {
	// get the next result from the combinator based on the saved index values
	result := []string{}

	// get all the characters for the result from the last active indexes
	for _, val := range s.indexes {
		result = append(result, s.Chars[val])
	}
	return strings.Join(result, "")
}

func (s *State) increment() {
	// increment the left-most index, then carry values through to the right-hand side
	s.indexes[0]++
	s.carry()
}

func (s *State) carry() {
	// reset index values from left to right to carry the value greater than the number of characters
	for i, val := range s.indexes {
		if val > len(s.Chars)-1 {
			// reset the left-hand value
			s.indexes[i] = 0
			// if we already have a value in the next right place, increment it
			if i < len(s.indexes)-1 {
				s.indexes[i+1]++
			} else {
				// otherwise add a new value in the next place
				s.indexes = append(s.indexes, 0)
			}
		}
	}
}

func (s *State) Next() string {
	// get the next combination then increment the indexes
	result := s.Get()
	s.increment()
	s.NumGenerated++
	return result
}

func NewState(charSet string, minSize int) State {
	// return a new blank state
	chars := strings.Split(charSet, "")
	chars = uniqueStrs(chars)
	sort.Strings(chars)
	state := State{
		Chars:   chars,
		indexes: []int{0},
	}
	for len(state.indexes) < minSize {
		state.indexes = append(state.indexes, 0)
	}
	return state
}

func uniqueStrs(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
