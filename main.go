package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	site := fillURLWithList("")
	document, err := loadSite(site)
	if err != nil {
		panic(err)
	}

	document.Find("div.lister-item.mode-detail").Each(func(_ int, s *goquery.Selection) {
		title := getTitle(s)
		score := getScore(s)
		metascore := getMetascore(s)
		votes := getVotes(s)

		average := getAverage(score, metascore)

		fmt.Println(title, score, metascore, votes, average)
	})
}
