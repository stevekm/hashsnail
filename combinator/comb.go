package combinator

import (
	"strings"
)

type State struct {
	Indexes []int
	Chars []string
}

func (s *State) Get() []string {
	// get the next result from the combinator based on the saved index values
	result := []string{}

	// get all the characters for the result from the last active indexes
	for _, val := range s.Indexes {
		result = append(result, s.Chars[val])
	}
	return result
}

func (s *State) Increment() {
	// increment the left-most index, then carry values through to the right-hand side
	s.Indexes[0]++
	s.Carry()
}

func (s *State) Carry(){
	// reset index values from left to right to carry the value greater than the number of characters
	for i, val := range s.Indexes {
		if val > len(s.Chars) - 1 {
			// reset the left-hand value
			s.Indexes[i] = 0
			// if we already have a value in the next right place, increment it
			if i < len(s.Indexes) - 1 {
				s.Indexes[i + 1]++
			} else {
				// otherwise add a new value in the next place
				s.Indexes = append(s.Indexes, 0)
			}
		}
	}
}

func (s *State) Next() []string {
	// get the next combination then increment the indexes
	result := s.Get()
	s.Increment()
	return result
}

func NewState() State {
	// return a new blank state
	chars := strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRST1234567890", "")
	state := State{
		Chars: chars,
		Indexes: []int{0},
	}
	return state
}