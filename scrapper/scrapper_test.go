package scrapper

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
    b := Start("https://google.com", "Does not exist", WuxiaWorldNextSelector)

    if len(b.pages) > 0 {
        t.Error("Not breaking out of the loop")
    }
}

func Test_loop_stops_once_the_chapter_is_preview(t *testing.T) {

}
