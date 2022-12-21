package hash

import (
	"fmt"
	"log"
	"strings"
	// "errors" //errors.New("No value found for hash")
	"crypto/md5"
	"hashsnail/combinator"
	// "crypto/sha1"
	// "crypto/sha256"
	"context"
	"encoding/hex"
	"runtime"
	"sync"
	"time"
)

type HashResult struct {
	Hash   string
	Result string
}

type HashFinder struct {
	NumCombs     int
	MaxSize      int // size of string to hash
	combinator   combinator.State
	Wanted       string // the hash we want to match
	Print        bool
	NumWorkers   int
	NumGenerated uint
	Result       HashResult
	Time         time.Duration // an int64 nanosecond count https://pkg.go.dev/time#Duration
	Rate float64
}

func (f *HashFinder) IsMaxSize(comb string) bool {
	if f.MaxSize < 0 {
		return false
	} else {
		return len(comb) > f.MaxSize
	}
}

func (f *HashFinder) Find() (string, error) {
	// find the string combination that matches the wanted hash
	for i := 0; i < f.NumCombs; i++ {
		comb := f.combinator.Next()
		if len(comb) > f.MaxSize {
			break
		}
		hash := f.GetHash(comb)
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

func (f *HashFinder) FindParallel() (string, error) {
	// find the string char combination that creates the desired hash
	// using concurrent parallel hash worker threads
	startTime := time.Now()

	numWorkers := f.NumWorkers
	runtime.GOMAXPROCS(numWorkers + 1) // add an extra for the combinator

	work := make(chan string)        // send comb's in here
	results := make(chan HashResult) // send hash result back out here

	// signal to stop sending work to the hash checkers
	ctx, cancel := context.WithCancel(context.Background())

	// create hash checker worker goroutines
	wg := sync.WaitGroup{}
	for i := 0; i < numWorkers; i++ {
		// add a worker
		wg.Add(1)
		// add a copy of the goroutine
		go func() {
			defer wg.Done()
			// get the next combination from the work channel
			for comb := range work {
				// make the hash
				hash := f.GetHash(comb)
				if f.Print {
					fmt.Printf("%v %v\n", comb, hash)
				}
				// check if its the correct hash
				if hash == f.Wanted {
					result := HashResult{
						Hash:   hash,
						Result: comb,
					}
					// send the correct result to the output channel
					results <- result
					// stop sending work
					cancel()
					// wait for the workers to finish
					// then close the results channel
					wg.Wait()
					close(results)
				}
			}
		}()
	}

	// send the work to the workers
	// this happens in a goroutine in order
	// to not block the main function,
	// once all workers are busy
	go func(ctx context.Context) {
		// send all string combinations to the hash workers

		// TODO: make this cleaner

		// unlimited number of combinations
		if f.NumCombs < 0 {
		combinatorUnlimited:
			for {
				select {
				case <-ctx.Done(): // if cancel() is executed, stop sending more work
					// exit the combinator loop
					break combinatorUnlimited
				default:
					// get the next combination
					comb := f.combinator.Next()
					// stop sending work if we have exceeded the max size
					if f.IsMaxSize(comb) {
						cancel()
						continue
					} else {
						// send the combination to the worker
						work <- comb
					}
				}
			}
		} else {
			// only run until we hit max NumCombs value
		combinatorNumCombs:
			for i := 0; i < f.NumCombs; i++ {
				select {
				case <-ctx.Done(): // if cancel() is executed, stop sending more work
					// exit the combinator loop
					break combinatorNumCombs
				default:
					// get the next combination
					comb := f.combinator.Next()
					// stop sending work if we have exceeded the max size
					if f.IsMaxSize(comb) {
						cancel()
						continue
					} else {
						// send the combination to the worker
						work <- comb
					}
				}
			}

			// close the work channel after all the work has been sent
			// wait for the workers to finish
			// then close the results channel
			// NOTE: this should only trigger if no hash results were found!
			close(work)
			wg.Wait()
			close(results)
		}

	}(ctx)

	// collect the results
	// the iteration stops if the results
	// channel is closed and the last value
	// has been received
	for result := range results {
		f.Result = result
		f.Time = time.Now().Sub(startTime)
		f.Rate = float64(f.NumGenerated) / f.Time.Seconds()
		if f.Print {
			log.Printf("RESULT:%v\n", result)
		}
		return result.Result, nil
	}

	return "", fmt.Errorf("No value found for hash %v", f.Wanted)
}

func (f *HashFinder) GetHash(comb string) string {
	hash := GetHash(comb)
	f.NumGenerated++
	return hash
}

func (f *HashFinder) DescribeStart() string {
	// make a condensed string output for printing the starting state of the finder
	type Desc struct {
		Wanted string
		NumCombs int
		MaxSize int
		NumWorkers int
		Print bool
		CharSet string
	}
	d := Desc{
		CharSet: strings.Join(f.combinator.Chars, ""),
		NumCombs: f.NumCombs,
		MaxSize: f.MaxSize,
		Wanted: f.Wanted,
		Print: f.Print,
		NumWorkers: f.NumWorkers,
	}

	return fmt.Sprintf("%#v", d)
}

func (f *HashFinder) DescribeResults() string {
	// return a string describing the final state of the finder
	return fmt.Sprintf("%v %v (%v hashes, %v, %.1fMH on %v workers)", 
		f.Result.Result, 
		f.Result.Hash, 
		f.NumGenerated, 
		f.Time, f.Rate / 1000000, // megahashes
		f.NumWorkers,
		) 
}

func NewHashFinder(numCombs int,
	maxSize int,
	minSize int,
	charSet string,
	wanted string,
	print bool,
	numWorkers int) HashFinder {
	comb := combinator.NewState(charSet, minSize)
	finder := HashFinder{
		NumCombs:   numCombs,
		MaxSize:    maxSize,
		Wanted:     wanted,
		combinator: comb,
		Print:      print,
		NumWorkers: numWorkers,
		Result:     HashResult{},
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

func PrintHashCombs(numCombs int, maxSize int) {
	state := combinator.NewState(combinator.CharSetDefault, 0)
	for i := 0; i < numCombs; i++ {
		comb := state.Next()
		if len(comb) > maxSize {
			break
		}
		fmt.Printf("%v %v\n", comb, GetHash(comb))
	}

}
