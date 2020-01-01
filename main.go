package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	list, average, votes := getParamsValues()

	site := fillURLWithList(list)
	document, err := loadSite(site)
	if err != nil {
		panic(err)
	}

	var movies []string

	document.Find("div.lister-item.mode-detail").Each(func(_ int, s *goquery.Selection) {
		title := getTitle(s)
		score := getScore(s)
		metascore := getMetascore(s)
		movieVotes := getVotes(s)

		movieAverage := getAverage(score, metascore)

		if movieAverage >= average && movieVotes >= votes {
			movies = append(movies, title)
		}
	})

	for _, movie := range movies {
		fmt.Println(movie)
	}
}
