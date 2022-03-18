package parser

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var name, price, url string

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

	doc.Find("div.buy_list > div.nayose").Each(func(i int, s *goquery.Selection) {
		name, _ = s.Attr("data-name")
		price, _ = s.Attr("data-price")
	})
	doc.Find("div.buy_list > div.nayose > div.nayose_head > div.itemBody > div.clearfix > div.itemImageContent > div.subImages > p.subImage > a.selected > img").Each(func(i int, s *goquery.Selection) {
		url, _ = s.Attr("data")

	})

	videoInfo := VideoInfo{}
	videoInfo.Name = name
	videoInfo.Price = price
	videoInfo.Url = url

	return &videoInfo, nil

}
