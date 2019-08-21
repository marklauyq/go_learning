package scrapper

import (
    "fmt"
    g "github.com/PuerkitoBio/goquery"
    "net/http"
)

type Chapter g.Document

type nextSelector func(*Chapter) string

type Book struct {
    pages []string
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

func checkPreview(c *Chapter) bool {
    return false
}

func Start(url, bodySelector string, next nextSelector) Book {
    b := Book{}
    pageUrl := url
    for {
        //get chapter
        c, err := getChapter(pageUrl)
        if err != nil {
            fmt.Println("Error", err)
            return Book{}
        }

        //check if it is a preview
        if checkPreview(c) {
            break
        }

        body := c.Find(bodySelector).Text()

        if body == "" {
            break
		}

        b.pages = append(b.pages, body)

        pageUrl := next(c)

        if pageUrl == "" {
            break
        }
    }

    return b
}
