package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/hideaki10/poogo/parser"
)

func init() {
	flag.Parse()
}

func main() {

	if flag.NArg() == 0 {
		log.Fatal("No arguments ")
	}

	url := flag.Arg(0)
	fmt.Println(url)

	viemo := parser.Viemo{Url: url}

	_, err := viemo.GetVideoInfo()
	if err != nil {
		log.Fatal(err)
	}
}
