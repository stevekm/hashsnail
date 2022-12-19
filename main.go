package main
// $ go run . 0cc175b9c0f1b6a831c399e269772661
// $ go run . e2fc714c4727ee9395f324cd2e7f331f --max-size 2 --progress

// comparison; <4s on M1 MacBook Air
// $ ./hashcat -m 0 -a 3 ab56b4d92b40713acc5af89985d4b786

import (
	"fmt"
	"log"
	_hash "hashsnail/hash"
	"hashsnail/combinator"
	"github.com/alecthomas/kong"
)

type CLI struct {
	Hash string `help:"hash string to crack" arg:""` // required positional arg
	MaxSize int `help:"max length of password to search for" default:3`
	MinSize int `help:"min length of password to search for" default:0`
	Progress bool `help:"print hasing progress to console"` // false by default
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
	result, err := finder.Find()
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