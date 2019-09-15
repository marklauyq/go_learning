package scraper

import (
    "testing"
)

func Test_able_to_get_body_from_getChapter(t *testing.T) {
    c, err := getChapter("https://www.wuxiaworld.com/novel/dragon-prince-yuan/yz-chapter-1")

    if err != nil {
        t.Error("There was an error", err)
    }

    if c.Find("body").Text() == "" {
        t.Error("Find is not working")
    }
}

func Test_loop_stops_once_there_is_no_next_page(t *testing.T) {
    b := Start("https://google.com", "Does not exist", WuxiaWorldNextSelector, Book{})

    if len(b.Pages) > 0 {
        t.Error("There is a problem with the loop")
    }

    b = Start(
        "https://boxnovel.com/novel/release-that-witch/chapter-1496",
        ".reading-content",
        BoxnovelNextSelector,
        Book{},
    )

    if len(b.Pages) != 3 {
        t.Error("There should only be 2 pages" , len(b.Pages))
    }
}
