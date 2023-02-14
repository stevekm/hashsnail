package combinator

import (
	// "log"
	"fmt"
	"sort"
	"strings"
	"time"
)

const CharSetDefault = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890 !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

type State struct {
	indexes      []int    // keeps track of which characters should be returned
	Chars        []string // character set to build combinations from
	NumGenerated uint     // number of combinations generated
	IterStart    int      // starting value for iteration
	IterStep     int      // step size for iteration
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
	for i := 0; i < s.IterStep; i++ {
		s.increment()
	}
	s.NumGenerated++
	return result
}

func NewState(charSet string, minSize int, iterStart int, iterStep int) State {
	// return a new blank state
	// clean up the input character set
	chars := strings.Split(charSet, "")
	chars = uniqueStrs(chars)
	sort.Strings(chars)
	// initialize a new State object
	state := State{
		Chars:     chars,
		indexes:   []int{0},
		IterStart: iterStart,
		IterStep:  iterStep,
	}
	// update the indexes to the minimum requested combination length
	for len(state.indexes) < minSize {
		state.indexes = append(state.indexes, 0)
	}
	// increment the iterator to the requested starting position
	for i := 0; i < iterStart; i++ {
		state.increment()
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

func PrintResult(name string, numCombs uint, startTime time.Time, stopTime time.Time) {
	// print out how many combinations were generated and the
	// rate of millions of comb's per second
	elapsed := stopTime.Sub(startTime)
	rate := float64(numCombs) / elapsed.Seconds()
	fmt.Printf("[%v] numCombs: %v, elapsed %v, rate:%.1fM/s\n",
		name, numCombs, elapsed, rate/1000000)
}

func DemoCombinator(numCombs int, name string) {
	state := NewState(CharSetDefault, 0, 0, 1)
	startTime := time.Now()
	var result string
	for i := 0; i < numCombs; i++ {
		result = state.Next()
	}
	stopTime := time.Now()
	result = ""
	PrintResult(name+result, state.NumGenerated, startTime, stopTime)
}

func DemoIterate(numCombs int, name string) {
	// raw iterator to see how fast we can increment an iterator
	// with no other operations going on
	startTime := time.Now()
	var result int
	for i := 0; i < numCombs; i++ {
		result++
		// continue
	}
	stopTime := time.Now()
	PrintResult(name, uint(numCombs), startTime, stopTime)
}
