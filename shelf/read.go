package shelf

import (
    "encoding/json"
    "fmt"
    s "github.com/marklauyq/go_learning/scraper"
    "io/ioutil"
)

func GetBook(title string) s.Book {
    barcode := titleToFile(title)
    //try to find the local file
    buf, err := ioutil.ReadFile(barcode)

    if err != nil {
        fmt.Println("No book was found ")
        fmt.Println(err)
        return s.Book{Title:title}
    }

    book := s.Book{}
    err = json.Unmarshal(buf, &book)

    if err != nil {
        fmt.Println("There was an error decoding Json")
        fmt.Println(err)
        return s.Book{Title:title}
    }

    return book
}