package scrapper

import (
	g "github.com/PuerkitoBio/goquery"
	"net/http"
)

type Chapter g.Document

type nextSelector func(*Chapter) string

type Book []*Chapter

// Scrap makes a request to a website and returns the body as a string
func getChapter(url string) (*Chapter, error) {
	resp, err := http.Get(url)

	if err != nil{
		return &Chapter{} , err
	}

	doc, err := g.NewDocumentFromReader(resp.Body)

	if err != nil {
		return &Chapter{} , err
	}

	c := Chapter(*doc)

	return &c, nil
}

func Start(url , bodySelector string, next nextSelector) Book{

	//start the loop
	{
		//first we need to retrieve the first chapter

		//add the chapter to  the book

		//than we need to find the url for the next chapter
	}

	return Book{}
}







