package scraper

import (
    "fmt"
    g "github.com/PuerkitoBio/goquery"
    "github.com/marklauyq/go_learning/shelf"
    "net/http"
    "strings"
)

type Chapter g.Document

// Scrap makes a request to a website and returns the body as a string
func getChapter(url string) (*Chapter, error) {
    resp, err := http.Get(url)

    if err != nil {
        return &Chapter{}, err
    }

    doc, err := g.NewDocumentFromReader(resp.Body)

    if err != nil {
        return &Chapter{}, err
    }

    c := Chapter(*doc)

    return &c, nil
}

func Start(url, bodySelector string, next nextSelector, b shelf.Book) shelf.Book {
    pageUrl := url
    for {

        //get chapter
        fmt.Println("Retrieving Page : " , pageUrl)
        c, err := getChapter(pageUrl)
        if err != nil {
            fmt.Println("Error", err)
            return b
        }

        b.LastChapter = pageUrl
        body := c.Find(bodySelector).Text()

        if body == "" {
            break
		}

        fmt.Println(b.Urls[pageUrl])
        if b.Urls[pageUrl] != true{
            p := shelf.Page(strings.TrimSpace(body))

            //append the Page to the book
            b.Pages = append(b.Pages, p)

            b.Urls[pageUrl] = true
        }

        pageUrl = next(c)

        if pageUrl == "" {
            break
        }
    }

    return b
}
