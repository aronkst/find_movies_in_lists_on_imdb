package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func loadSite(url string) (*goquery.Document, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		message := fmt.Sprintf("Status code error: %d %s", response.StatusCode, response.Status)
		err := errors.New(message)
		return nil, err
	}

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	return document, nil
}

func getValueFromSiteDocument(document *goquery.Document, selector string, attribute string) string {
	selection := document.Selection
	return getValueFromSiteSelection(selection, selector, attribute)
}

func getValueFromSiteSelection(selection *goquery.Selection, selector string, attribute string) string {
	var value string

	if attribute == "" {
		value = selection.Find(selector).Text()
	} else {
		value, _ = selection.Find(selector).Attr(attribute)
	}

	value = clearString(value)
	return value
}

func getTitle(selection *goquery.Selection) string {
	return getValueFromSiteSelection(selection, "div.lister-item-content h3.lister-item-header a", "")
}

func getScore(selection *goquery.Selection) float64 {
	score := getValueFromSiteSelection(selection, "div.ipl-rating-widget div.ipl-rating-star.small span.ipl-rating-star__rating", "")
	return stringToFloat(score)
}

func getMetascore(selection *goquery.Selection) int64 {
	metascore := getValueFromSiteSelection(selection, "div.inline-block.ratings-metascore span.metascore.favorable", "")
	if metascore == "" {
		return 0
	}
	return stringToInt(metascore)
}

func getVotes(selection *goquery.Selection) int64 {
	votes := getValueFromSiteSelection(selection, "p.text-muted.text-small span[name='nv']", "data-value")
	return stringToInt(votes)
}
