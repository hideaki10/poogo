package parser

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type Viemo struct {
	Url string
}

func (vi *Viemo) GetVideoInfo() (*VideoInfo, error) {
	req, err := http.NewRequest(http.MethodGet, vi.Url, nil)
	if err != nil {
		return nil, err
	}

	header := http.Header{}
	header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")
	header.Add("referer", "https://www.ixigua.com/")
	req.Header = header
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// doc, err := goquery.NewDocumentFromReader(resp.Body)
	// if err != nil {
	// 	return nil, err
	// }

	// head := doc.Find("head")
	// fmt.Println(head.Text())
	// head.Find("meta").Each(func(i int, s *goquery.Selection) {
	// 	if propetry, _ := s.Attr("property"); propetry == "og:title" {
	// 		title, _ := s.Attr("content")
	// 		fmt.Println(title)
	// 	}

	// 	fmt.Println("noting")
	// })

	resultBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	html := string(resultBytes)
	pattern := `"title":".*","tag"`

	compile := regexp.MustCompile(pattern)
	matches := compile.FindStringSubmatch(html)

	for _, match := range matches {
		fmt.Println(match)
	}

	// log.Println(string(resultBytes))

	return &VideoInfo{}, nil

	// return nil, nil
}
