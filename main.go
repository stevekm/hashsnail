package main

import (
	"fmt"
	"github.com/alecthomas/kong"
	"hashsnail/combinator"
	_hash "hashsnail/hash"
	"log"
)

type CLI struct {
	Hash     string `help:"hash string to crack" arg:""` // required positional arg
	MaxSize  int    `help:"max length of password to search for" default:3`
	MinSize  int    `help:"min length of password to search for" default:0`
	Progress bool   `help:"print hasing progress to console"` // false by default
}

func (cli *CLI) Run() error {
	err := run(
		cli.Hash,
		cli.MaxSize,
		cli.MinSize,
		cli.Progress,
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
) error {
	numCombs := 10000000 * 10000000 // a big number
	charSet := combinator.CharSetDefault
	print := progress
	wanted := hash
	finder := _hash.NewHashFinder(numCombs, maxSize, minSize, charSet, wanted, print)
	// result, err := finder.Find()
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
