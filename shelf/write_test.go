package shelf

import (
    "fmt"
    "testing"
)


func TestIfFileAlreadyExistIShouldAppend(t *testing.T) {
    book := GetBook("Test Book")
    book.Pages = append(book.Pages, "Third Chapter");
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


