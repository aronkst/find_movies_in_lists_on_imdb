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
		title := getValueFromSiteSelection(s, "div.lister-item-content h3.lister-item-header a", "")
		score := getValueFromSiteSelection(s, "div.ipl-rating-widget div.ipl-rating-star.small span.ipl-rating-star__rating", "")
		metascore := getValueFromSiteSelection(s, "div.inline-block.ratings-metascore span.metascore.favorable", "")
		votes := getValueFromSiteSelection(s, "p.text-muted.text-small span[name='nv']", "data-value")
		fmt.Println(title, score, metascore, votes)
	})
}
