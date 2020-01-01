package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	site := fillURL("")
	document, err := loadSite(site)
	if err != nil {
		panic(err)
	}

	document.Find("div.lister-item.mode-detail").Each(func(_ int, s *goquery.Selection) {
		title := getValueFromSiteSelection(s, "div.lister-item-content h3.lister-item-header a", "")
		fmt.Println(title)
	})
}
