package shelf

import (
    "encoding/json"
    s "github.com/marklauyq/go_learning/scraper"
    "io/ioutil"
    "log"
    "os"
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


func WriteToTxt(book s.Book) error {
    barcode := titleToFileWithExtensions(book.Title, ".txt")

    f, err := os.OpenFile(barcode,
        os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

    if err != nil {
        log.Println(err)
        return err
    }

    for chapter, _ := range book.Pages {
        if _, err := f.WriteString("----- END ------"); err != nil {
            log.Println(err)
        }
        if _, err := f.WriteString(chapter); err != nil {
            log.Println(err)
        }
    }

    return f.Close()
}
