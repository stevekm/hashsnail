package main

import (
	// "fmt"
	"github.com/alecthomas/kong"
	"hashsnail/combinator"
	_hash "hashsnail/hash"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

// https://pkg.go.dev/runtime/pprof
// https://github.com/google/pprof/blob/main/doc/README.md
// $ go tool pprof cpu.prof
// $ go tool pprof mem.prof
// (pprof) top
func StartProfiler() (*os.File, *os.File) {
	log.Println("Starting profiler")
	cpuFileName := "cpu.prof"
	cpuFile, err := os.Create(cpuFileName)
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	// defer cpuFile.Close() // error handling omitted for example
	if err := pprof.StartCPUProfile(cpuFile); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	// defer pprof.StopCPUProfile()
	log.Printf("cpuFile: %v\n", cpuFileName)

	memFileName := "mem.prof"
	memFile, err := os.Create(memFileName)
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	// defer WriteMemoryProfile(memFile)
	log.Printf("memFile: %v\n", memFileName)

	return cpuFile, memFile
}
func WriteMemoryProfile(memFile *os.File) {
	runtime.GC() // get up-to-date statistics
	if err := pprof.WriteHeapProfile(memFile); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
}

type CLI struct {
	Hash     string  `help:"hash string to crack" arg:""` // required positional arg
	MaxSize  int     `help:"max length of password to search for" default:-1`
	MinSize  int     `help:"min length of password to search for" default:0`
	Progress bool    `help:"print hasing progress to console"` // false by default
	Threads  *int    `help:"number of CPU threads to use, defaults all CPU cores"`
	CharSet  *string `help:"characters to use for search"`
	Combs    *int    `help:"max number of character combinations to test, defaults to unlimited"`
	Debug    bool    `help:"this option does nothing do not use it"` // false by default
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
		cli.Debug,
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
	debug bool,
) error {
	// $ go run . --debug zzz
	if debug == true {
		log.Println("debug activated")
		cpuFile, memFile := StartProfiler()
		defer cpuFile.Close()
		defer memFile.Close()
		defer pprof.StopCPUProfile()
		defer WriteMemoryProfile(memFile)

		// put some function calls here for when I am debugging things and need
		// an easier CLI entrypoint
		// combinator.DemoCombinator(100000000, "debug")
		// combinator.DemoIterate(10000000000, "debug")
		return nil
	}
	numThreads := runtime.NumCPU()
	if threads != nil {
		numThreads = *threads
	}
	numCombs := -1
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
	log.Printf("Starting finder:%v\n", finder.DescribeStart())
	_, err := finder.FindParallel2() // result, err
	if err != nil {
		log.Fatalf("%v", err)
		return err
	}
	log.Printf("Result found: %v", finder.DescribeResults())

	return nil
}

func main() {
	var cli CLI

	ctx := kong.Parse(&cli,
		kong.Name("hashsnail"),
		kong.Description("program for cracking hash values"))

	ctx.FatalIfErrorf(ctx.Run())

}
