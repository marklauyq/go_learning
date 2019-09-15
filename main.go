package main

import (
    "fmt"
    s "github.com/marklauyq/go_learning/scraper"
    "go_learning/shelf"
)

func main() {
    b := s.Book{
        Title: "Death March",
    }

    s.Start(
        "https://boxnovel.com/novel/death-march-kara-hajimaru-isekai-kyusoukyoku-wn/chapter-1",
        s.BoxnovelBody,
        s.BoxnovelNextSelector,
        b,
    )

    err := shelf.WriteBook(b)

    if err != nil {
        fmt.Println(err)
    }
}
