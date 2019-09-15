package shelf

import (
    "encoding/json"
    s "github.com/marklauyq/go_learning/scraper"
    "io/ioutil"
)

func WriteBook(book s.Book) error{
    barcode := titleToFile(book.Title)

    b, err := json.Marshal(book)

    if err != nil {
        return err
    }

    err = ioutil.WriteFile(barcode, b, 0777)

    if err != nil {
        return err
    }

    return nil
}
