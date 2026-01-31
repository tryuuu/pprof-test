package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/tryuuu/pprof-test/pkg/bad"
	"github.com/tryuuu/pprof-test/pkg/good"
)

func main() {
	workType := flag.String("type", "good", "execution type: good or bad")
	flag.Parse()

	fmt.Printf("running heavywork (%s)\n", *workType)

	switch *workType {
	case "good":
		good.Do()
	case "bad":
		bad.Do()
	default:
		fmt.Printf("unknown type: %s\n", *workType)
		os.Exit(1)
	}

	fmt.Println("done")

	memFile, _ := os.Create("mem.prof")
	defer memFile.Close()
	runtime.GC()
	if err := pprof.WriteHeapProfile(memFile); err != nil {
		fmt.Fprintf(os.Stderr, "could not write memory profile: %v\n", err)
	}
}
