package scrapper

import (
	g "github.com/PuerkitoBio/goquery"
	"net/http"
)

type chapter g.Document

type book struct {
	chapters []*chapter
	title string
}

// Scrap makes a request to a website and returns the body as a string
func getChapter(url string) (*chapter, string, error) {
	resp, err := http.Get(url)

	if err != nil{
		return &chapter{} , "", err
	}

	doc, err := g.NewDocumentFromReader(resp.Body)

	if err != nil {
		return &chapter{} , "" , err
	}

	c := chapter(*doc)

	return &c, "" , nil
}

func Start(url , bodySelector, nextSelect string) {
	
}







