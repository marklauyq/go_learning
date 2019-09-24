package shelf

import (
    "fmt"
    "testing"
)


func TestIfFileAlreadyExistIShouldAppend(t *testing.T) {
    book := GetBook("Test Book")
    book.Pages["chapter-3"] = "Chapter 3 here"
    book.Title = "New Book"

    err := WriteBook(book)

    bookToTest := GetBook(book.Title)

    if  err != nil {
        fmt.Println(err)
        t.Error("There was a problem writing to book")
    }

    if len(bookToTest.Pages) != 3 {
        t.Error("The number of pages is wrong expected 3 got" , len(bookToTest.Pages))
    }
}


func TestWriteToEpub(t *testing.T) {
    book := GetBook("Test Book")
    book.Pages["chapter-3"] = "Chapter 3 here"
    book.Title = "New Book"

    err := WriteToEpub(book)

    if  err != nil {
        fmt.Println(err)
        t.Error("There was a problem writing to book")
    }
}

