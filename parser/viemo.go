package parser

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
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
	header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")
	//header.Add("Referer", "https://lgtmoon.herokuapp.com/")

	req.Header = header
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	// doc.Find("div.buy_list > div.nayose > div.nayose_head > h2.link > p > a").Each(func(i int, s *goquery.Selection) {

	// 	log.Println(s.Text())
	// })
	doc.Find("div.buy_list > div.nayose").Each(func(i int, s *goquery.Selection) {
		log.Println(s.Attr("data-name"))
		log.Println(s.Attr("data-price"))
	})
	doc.Find("div.buy_list > div.nayose > div.nayose_head > div.itemBody > div.clearfix > div.itemImageContent > div.subImages > p.subImage > a.selected > img").Each(func(i int, s *goquery.Selection) {
		log.Println(s.Attr("data"))

	})
	//head := doc.Find("head")

	// head.Find("meta").Each(func(i int, s *goquery.Selection) {
	// 	if propetry, _ := s.Attr("property"); propetry == "og:title" {
	// 		title, _ := s.Attr("content")
	// 		fmt.Println(title)
	// 	}

	// 	fmt.Println("noting")
	// })

	// resultBytes, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, err
	// }
	//fmt.Println(resultBytes)
	//html := string(resultBytes)
	//pattern := `"title":".*","tag"`

	//compile := regexp.MustCompile(pattern)
	//matches := compile.FindStringSubmatch(html)

	// for _, match := range matches {
	// 	fmt.Println(match)
	// }

	//log.Println(string(resultBytes))

	return &VideoInfo{}, nil

	// return nil, nil
}
