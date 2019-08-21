package scraper

import (
    "fmt"
    g "github.com/PuerkitoBio/goquery"
    "net/http"
    "strings"
)

type Chapter g.Document
type page string

type Book struct {
    pages []page
}

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

func Start(url, bodySelector string, next nextSelector) Book {
    b := Book{}
    pageUrl := url
    for {
        //get chapter
        fmt.Println("Retrieving page : " , pageUrl)
        c, err := getChapter(pageUrl)
        if err != nil {
            fmt.Println("Error", err)
            return Book{}
        }

        body := c.Find(bodySelector).Text()

        if body == "" {
            break
		}

        b.pages = append(b.pages, page(strings.TrimSpace(body)))

        pageUrl = next(c)

        if pageUrl == "" {
            break
        }
    }

    return b
}
