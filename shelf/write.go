package shelf

import (
    "encoding/json"
    "fmt"
    s "github.com/marklauyq/go_learning/scraper"
    epub "github.com/bmaupin/go-epub"
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

    for url, chapter := range book.Pages {
        fmt.Println("Writing URL : " + url)
        if _, err := f.WriteString("\n\n ----- END ------ \n\n"); err != nil {
            log.Println(err)
        }
        if _, err := f.WriteString(string(chapter)); err != nil {
            log.Println(err)
        }
    }

    return f.Close()
}

func WriteToEpub(book s.Book) error {
    barcode := titleToFileWithExtensions(book.Title, ".epub")

    e := epub.NewEpub(book.Title)

    e.SetAuthor("Web Novels")

    chapterNum := 1
    for url, chapter := range book.Pages {
        fmt.Println("Writing URL : " + url)
        chapterStr := fmt.Sprintf(`Chapter %v` , chapterNum)
        sectionBody := fmt.Sprintf(`<h1> %s</h1> <p> %s </p>`, chapterStr , string(chapter))

        chapterNum += 1
        _, err := e.AddSection(sectionBody, chapterStr,url,"")

        if err != nil{
            fmt.Println("ERROR " , err )
        }
    }

    err := e.Write(barcode)

    if err != nil {
        return err
    }

    return nil
}