package main

import (
    s "github.com/marklauyq/go_learning/scraper"
)

func main() {
    b := s.Book{
        Title: "Dragon Prince Yuan",
    }

    s.Start(
        "https://www.wuxiaworld.com/novel/dragon-prince-yuan/yz-chapter-1",
        "#chapter-content",
        s.WuxiaWorldNextSelector,
        b,
    )
}
