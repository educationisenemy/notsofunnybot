package kaktam

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strconv"
)

type Link struct {
	Text string
	Href string
}

func GrabLinks(lastNewsId int) []Link {
	url := "http://kaktam.ru/data-news-list.php?last_news_id=" + strconv.Itoa(lastNewsId);
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	links := []Link{}
	doc.Find(".newscard").Each(func(i int, s *goquery.Selection) {
		text := s.Find("div.newscard__title").Text()
		href, exists := s.Find("a").Attr("href")

		if(exists) {
			links = append(links, Link{text, href});
		}
	})

	return links
}