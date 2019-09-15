package scraper

type Page string

type Book struct {
    Title string
    Pages map[string]Page
    LastChapter string
}

