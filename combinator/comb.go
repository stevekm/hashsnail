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

// func GetCombs(charSet string, numCombs int, minSize int) []string {
// 	// return a list of all combinations up to a specific amount
// 	combs := []string{}
// 	state := NewState(charSet, minSize)
// 	for i := 0; i < numCombs; i++ {
// 		combs = append(combs, state.Next())
// 	}
// 	return combs
// }

// // methods for debugging and testing

// func CombTimer() {
// 	//
// 	// method for calculating how fast we can generate combinations
// 	// because I think that is bottlenecking the program
// 	//
// 	// 2022/12/27 19:47:07 numCombs: 100, elapsed 29.996µs, rate:3.3M/s
// 	// 2022/12/27 19:47:07 numCombs: 1000, elapsed 159.849µs, rate:6.3M/s
// 	// 2022/12/27 19:47:07 numCombs: 10000, elapsed 1.640843ms, rate:6.1M/s
// 	// 2022/12/27 19:47:07 numCombs: 100000, elapsed 27.661377ms, rate:3.6M/s
// 	// 2022/12/27 19:47:07 numCombs: 1000000, elapsed 448.970281ms, rate:2.2M/s
// 	// 2022/12/27 19:47:11 numCombs: 10000000, elapsed 3.478159358s, rate:2.9M/s
// 	// 2022/12/27 19:47:17 numCombs: 20000000, elapsed 6.037647609s, rate:3.3M/s
// 	// 2022/12/27 19:47:28 numCombs: 40000000, elapsed 11.67794788s, rate:3.4M/s
// 	// 2022/12/27 19:47:50 numCombs: 80000000, elapsed 21.39272071s, rate:3.7M/s
// 	// 2022/12/27 19:48:18 numCombs: 100000000, elapsed 27.834257914s, rate:3.6M/s
// 	// 2022/12/27 19:49:00 numCombs: 150000000, elapsed 42.817057752s, rate:3.5M/s
// 	// 2022/12/27 19:50:01 numCombs: 200000000, elapsed 1m0.495746696s, rate:3.3M/s

// 	nums := []int{
// 		100,
// 		1000,
// 		10000,
// 		100000,
// 		1000000,
// 		10000000,
// 		20000000,
// 		40000000,
// 		80000000,
// 		100000000,
// 		150000000,
// 		200000000,
// 	}
// 	log.Println(nums)
// 	for _, num := range nums {
// 		startTime := time.Now()
// 		combs := GetCombs(CharSetDefault, num, 0)
// 		numCombs := len(combs)
// 		elapsed := time.Now().Sub(startTime)
// 		rate := float64(numCombs) / elapsed.Seconds()

// 		log.Printf("numCombs: %v, elapsed %v, rate:%.1fM/s",
// 			numCombs,
// 			elapsed,
// 			rate/1000000,
// 		)
// 	}
// }

// func CombTimer2() {
// 	// 2022/12/27 20:14:36 numCombs2: 100, elapsed 34.444µs, rate:2.9M/s
// 	// 2022/12/27 20:14:36 numCombs2: 1000, elapsed 127.668µs, rate:7.8M/s
// 	// 2022/12/27 20:14:36 numCombs2: 10000, elapsed 1.243411ms, rate:8.0M/s
// 	// 2022/12/27 20:14:36 numCombs2: 100000, elapsed 17.733685ms, rate:5.6M/s
// 	// 2022/12/27 20:14:37 numCombs2: 1000000, elapsed 178.165684ms, rate:5.6M/s
// 	// 2022/12/27 20:14:40 numCombs2: 10000000, elapsed 3.747518823s, rate:2.7M/s
// 	// 2022/12/27 20:14:49 numCombs2: 20000000, elapsed 8.297200054s, rate:2.4M/s
// 	// 2022/12/27 20:15:05 numCombs2: 40000000, elapsed 16.489429985s, rate:2.4M/s
// 	// 2022/12/27 20:15:38 numCombs2: 80000000, elapsed 32.87068933s, rate:2.4M/s
// 	// 2022/12/27 20:16:22 numCombs2: 100000000, elapsed 43.965012641s, rate:2.3M/s
// 	// 2022/12/27 20:17:34 numCombs2: 150000000, elapsed 1m12.368473008s, rate:2.1M/s
// 	// 2022/12/27 20:19:14 numCombs2: 200000000, elapsed 1m40.225849975s, rate:2.0M/s

// 	nums := []int{
// 		100,
// 		1000,
// 		10000,
// 		100000,
// 		1000000,
// 		10000000,
// 		20000000,
// 		40000000,
// 		80000000,
// 		100000000,
// 		150000000,
// 		200000000,
// 	}
// 	log.Println(nums)
// 	for _, num := range nums {
// 		startTime := time.Now()

// 		state := NewState(CharSetDefault, 0)
// 		for i := 0; i < num; i++ {
// 			state.Next()
// 		}

// 		elapsed := time.Now().Sub(startTime)
// 		rate := float64(num) / elapsed.Seconds()

// 		log.Printf("numCombs2: %v, elapsed %v, rate:%.1fM/s",
// 			num,
// 			elapsed,
// 			rate/1000000,
// 		)
// 	}
// }

// func IterTimer() {
// 	nums := []int{
// 		100,
// 		1000,
// 		10000,
// 		100000,
// 		1000000,
// 		10000000,
// 		20000000,
// 		40000000,
// 		80000000,
// 		100000000,
// 		150000000,
// 		200000000,
// 	}
// 	for _, num := range nums {
// 		startTime := time.Now()
// 		for i := 0; i < num; i++ {
// 			continue
// 		}
// 		elapsed := time.Now().Sub(startTime)
// 		rate := float64(num) / elapsed.Seconds()

// 		log.Printf("num: %v, elapsed %v, rate:%.1fM/s",
// 			num,
// 			elapsed,
// 			rate/1000000,
// 		)

// 	}
// }
