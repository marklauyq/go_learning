package main

import (
    s "go_learning/scraper"
)

func main() {
    s.Start(
        "https://www.wuxiaworld.com/novel/dragon-prince-yuan/yz-chapter-1",
        "#chapter-content",
        s.WuxiaWorldNextSelector,
    )
}
