package shelf

import (
    "fmt"
    "testing"
)


func TestIfFileAlreadyExistIShouldAppend(t *testing.T) {
    book := GetBook("Test Book")

    err := AddChapter(&book, "Third chapter here")

    if  err != nil {
        t.Error("There was a problem adding a chapter")
    }

    if len(book.Pages) != 3 {
        t.Error("The number of pages is wrong")
    }
}

func TestIfFileDoesNotExistIShouldCreateNewFile(t *testing.T) {
    book := GetBook("New Book")

    err := AddChapter(&book, "First Chapter")

    if  err != nil {
        t.Error("There was a problem adding a chapter")
    }

    if len(book.Pages) != 1 {
        t.Error("The number of pages is wrong")
    }
}


