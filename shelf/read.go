package shelf

import (
    "encoding/json"
    "fmt"
    s "go_learning/scraper"
    "io/ioutil"
)

func GetBook(title string) s.Book {
    barcode := titleToFile(title)
    barcode = barcode + ".json"

    //try to find the local file
    buf, err := ioutil.ReadFile(barcode)

    if err != nil {
        fmt.Println("No book was found ")
        fmt.Println(err)
        return s.Book{Title:title}
    }

    book := s.Book{Title:title}
    p := &book.Pages

    err = json.Unmarshal(buf, p)

    if err != nil {
        fmt.Println("There was an error decoding Json")
        fmt.Println(err)
        return s.Book{Title:title}
    }

    return book
}