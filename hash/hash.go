package hash

import (
	"log"
	"fmt"
	// "errors" //errors.New("No value found for hash")
	"crypto/md5"
	"hashsnail/combinator"
	// "crypto/sha1"
	// "crypto/sha256"
	"encoding/hex"
	"runtime"
	"sync"
	"context"
)

type HashResult struct {
	Hash string
	Result string
}

type HashFinder struct {
	NumCombs   int
	MaxSize    int // size of string to hash
	combinator combinator.State
	Wanted     string // the hash we want to match
	Print      bool
}

func (f *HashFinder) Find() (string, error) {
	// find the string combination that matches the wanted hash
	for i := 0; i < f.NumCombs; i++ {
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



func (f *HashFinder) FindParallel() (string, error) {
	numWorkers := 2
	runtime.GOMAXPROCS(numWorkers) // runtime.NumCPU()

	work := make(chan string) // send comb's in here
	results := make(chan HashResult) // send hash results back out here
	ctx, cancel := context.WithCancel(context.Background()) // signal to stop sending work

	// create worker goroutines
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
				hash := GetHash(comb)
				if f.Print {
					fmt.Printf("%v %v\n", comb, hash)
				}
				// check if its the correct hash
				if hash == f.Wanted {
					result := HashResult{
						Hash: hash,
						Result: comb,
					}
					// fmt.Printf("worker found a result: %v %v\n", result.Result, result.Hash)
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
		// fmt.Println("starting combinator")
		combinator:
		for i := 0; i < f.NumCombs; i++ {
			select {
			case <-ctx.Done(): // if cancel() execute, stop sending more work
				// fmt.Println("combinator got cancel() signal")
				// close the work channel and exit
				// close(work)
				// return
				break combinator
			default:
				// get the next combination
				comb := f.combinator.Next()
				// stop sending work if we have exceeded the max size
				if len(comb) > f.MaxSize {
					// fmt.Println("combinator max size exceeded")
					cancel()
					// close(work)
					// return
					// break
					continue
				} else {
					// send the combination to the worker
					// fmt.Printf("sending comb:%v\n", comb)
					work <- comb
				}
			}
		}
		// fmt.Println("done all combinations")

		// close the work channel after all the work has been send
		// wait for the workers to finish
		// then close the results channel
		// NOTE: this should only trigger if no hash results were found!
		close(work)
		wg.Wait()
		close(results)
	}(ctx)

	// collect the results
	// the iteration stops if the results
	// channel is closed and the last value
	// has been received
	for result := range results {
		log.Printf("RESULT:%v\n", result)
		return result.Result, nil
	}

	return "", fmt.Errorf("No value found for hash %v", f.Wanted)
}

func NewHashFinder(numCombs int, maxSize int, minSize int, charSet string, wanted string, print bool) HashFinder {
	comb := combinator.NewState(charSet, minSize)
	finder := HashFinder{
		NumCombs:   numCombs,
		MaxSize:    maxSize,
		Wanted:     wanted,
		combinator: comb,
		Print:      print,
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
