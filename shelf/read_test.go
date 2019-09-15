package shelf

import (
    "fmt"
    "io/ioutil"
    "os"
    "testing"
)

func setup(){
    fmt.Println ("Creating the test book")
    s := `["This is a test book", "There are only 2 pages on it"]`

    _ = ioutil.WriteFile("test_book.json", []byte(s) , 0666)
}
func shutdown(){
    fmt.Println("Removing book")

    err := os.Remove("test_book.json")
    if err != nil{
        fmt.Println("There was an issue removing the test file")
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
    book := GetBook("Test book")

    if book.Title != "Test book"{
        t.Error("There is a problem with the title")
    }

    if len(book.Pages) == 0 {
        t.Error("There should be pages in a book")
    }

    if len(book.Pages) != 2 {
        t.Error("The number of pages is wrong")
    }
}
