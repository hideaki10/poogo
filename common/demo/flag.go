package demo

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// 1
	var flagInt = flag.Int("fn", 100, "help message for flagname")
	var flagString = flag.String("fs", "name", "help message for flagstring")

	// 2
	var flagValue int
	flag.IntVar(&flagValue, "flagname", 30, "help message for flagname")

	// custom usage
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `poogo version: 1.0.0 Usage:poogo [options]`)

	}

	flag.Parse()
	fmt.Println(*flagInt, *flagString, flagValue)

	// flag.Arg
	fmt.Println(flag.Args())
	fmt.Println(flag.NArg())

}
