package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	document, err := loadSite("")
	if err != nil {
		panic(err)
	}
	fmt.Println(document)
}

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
