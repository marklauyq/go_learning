package shelf

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
)

func GetBook(title string) Book {
    barcode := titleToFile(title)
    //try to find the local file
    buf, err := ioutil.ReadFile(barcode)

    if err != nil {
        fmt.Println("No book was found ")
        fmt.Println(err)
        return Book{Title:title, Pages:[]Page{}, Urls: map[string]bool{}}
    }

    book := Book{}
    err = json.Unmarshal(buf, &book)

    if err != nil {
        fmt.Println("There was an error decoding Json")
        fmt.Println(err)
        return Book{Title:title, Pages:[]Page{}, Urls: map[string]bool{}}
    }

    return book
}