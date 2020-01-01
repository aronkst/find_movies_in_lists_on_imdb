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

	var movies []string

	document.Find("div.lister-item.mode-detail").Each(func(_ int, s *goquery.Selection) {
		title := getTitle(s)
		score := getScore(s)
		metascore := getMetascore(s)
		votes := getVotes(s)

		average := getAverage(score, metascore)

		if average >= 7.5 && votes >= 100000 {
			movies = append(movies, title)
		}
	})

	for _, movie := range movies {
		fmt.Println(movie)
	}
}
