package main

import (
	"fmt"
	"github.com/alecthomas/kong"
	"hashsnail/combinator"
	_hash "hashsnail/hash"
	"log"
	"runtime"
)

type CLI struct {
	Hash     string `help:"hash string to crack" arg:""` // required positional arg
	MaxSize  int    `help:"max length of password to search for" default:8`
	MinSize  int    `help:"min length of password to search for" default:0`
	Progress bool   `help:"print hasing progress to console"` // false by default
	Threads *int `help:"number of CPU threads to use, defaults all CPU cores"`
	CharSet *string `help:"characters to use for search"`
	Combs *int `help:"max number of character combinations to test, defaults to unlimited"`
}

func (cli *CLI) Run() error {
	err := run(
		cli.Hash,
		cli.MaxSize,
		cli.MinSize,
		cli.Progress,
		cli.Threads,
		cli.CharSet,
		cli.Combs,
	)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}

func run(
	hash string,
	maxSize int,
	minSize int,
	progress bool,
	threads *int,
	chars *string,
	combs *int,
) error {
	numThreads := runtime.NumCPU()
	if threads != nil {
		numThreads = *threads
	}
	numCombs := -1 // 10000000 * 10000000 // a big number
	if combs != nil {
		numCombs = *combs
	}
	charSet := combinator.CharSetDefault
	if chars != nil {
		charSet = *chars
	}
	print := progress
	wanted := hash

	finder := _hash.NewHashFinder(numCombs, maxSize, minSize, charSet, wanted, print, numThreads)
	// fmt.Printf("finder:%v\n", finder)
	result, err := finder.FindParallel()
	if err != nil {
		log.Fatalf("%v", err)
		return err
	}
	fmt.Printf(">>> FOUND value '%v' for hash %v\n", result, wanted)

	return nil
}

func main() {
	var cli CLI

	ctx := kong.Parse(&cli,
		kong.Name("hashsnail"),
		kong.Description("program for cracking hash values"))

	ctx.FatalIfErrorf(ctx.Run())

}
