package shelf

type Page string

type Book struct {
    Title string
    Pages []Page
    Urls map[string]bool
    LastChapter string
}

