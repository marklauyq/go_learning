package shelf

import (
    "fmt"
    "io/ioutil"
    "os"
    "testing"
)

func setup(){
    fmt.Println ("Creating the test book")
    s := `{"Title":"Test Book", "Pages":{"Chapter-1":"First Chapter", "Chapter-2":"Second Chapter"}, 
"LastChapter":"some-url-here"}`

    _ = ioutil.WriteFile("test_book.json", []byte(s) , 0666)
}
func shutdown(){
    fmt.Println("Removing book")

    err := os.Remove("test_book.json")
    if err != nil{
        fmt.Println("There was an issue removing the test file")
    }

    err = os.Remove("new_book.json")
    if err != nil{
        fmt.Println(err)
    }
}

func TestMain(m *testing.M) {
    setup()
    code := m.Run()
    shutdown()
    os.Exit(code)
}

func TestReadShouldReturnEmptyBookIfUnableToFind(t *testing.T) {
    book := GetBook("Book does not exist")

    if len(book.Pages) != 0 {
        t.Error("I should be getting an empty book")
    }
}

func TestReadShouldBeAbleToRetrieveTheBook(t *testing.T) {
    book := GetBook("Test Book")

    fmt.Println(book)
    if book.Title != "Test Book"{
        t.Error("There is a problem with the title")
    }

    if len(book.Pages) == 0 {
        t.Error("There should be pages in a book")
    }

    if len(book.Pages) != 2 {
        t.Error("The number of pages is wrong")
    }
}
