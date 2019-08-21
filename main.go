package main

import (
    "fmt"
    s "go_learning/scrapper"
)

func main() {
    fmt.Println(s.Start(
        "https://www.wuxiaworld.com/novel/dragon-prince-yuan/yz-chapter-1",
        "#chapter-content",
        s.WuxiaWorldNextSelector,
    ),
    )
}
