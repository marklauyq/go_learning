package main

import (
    "fmt"
    s "github.com/marklauyq/go_learning/scraper"
    "github.com/marklauyq/go_learning/shelf"
)

func main() {
    b := shelf.GetBook("Death March")

    if len(b.Pages) == 0 {
        b.LastChapter = "https://boxnovel.com/novel/death-march-kara-hajimaru-isekai-kyusoukyoku-wn/chapter-1"
    }

    b = s.Start(
        b.LastChapter,
        s.BoxnovelBody,
        s.BoxnovelNextSelector,
        b,
    )

    err := shelf.WriteBook(b)

    if err != nil {
        fmt.Println(err)
    }
}
