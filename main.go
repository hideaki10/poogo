package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/hideaki10/poogo/nets/download"
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

	videoInfo, err := viemo.GetVideoInfo()
	if err != nil {
		log.Fatal(err)
	}
	header := http.Header{}
	header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")

	fmt.Println(videoInfo.Name, videoInfo.Price, videoInfo.Url)

	err = download.Download(videoInfo.Name, videoInfo.Url, header)
	if err != nil {
		log.Fatal("download failed")
	}
	fmt.Println("download success")
}
