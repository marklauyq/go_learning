package shelf

import (
    "encoding/json"
    "fmt"
    "github.com/bmaupin/go-epub"
    "io/ioutil"
    "log"
    "os"
)

func WriteBook(book Book) error{
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


func WriteToTxt(book Book) error {
    barcode := titleToFileWithExtensions(book.Title, ".txt")

    f, err := os.OpenFile(barcode,
        os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

    if err != nil {
        log.Println(err)
        return err
    }

    for index, chapter := range book.Pages {
        fmt.Println("Writing URL : " + string(index))
        if _, err := f.WriteString("\n\n ----- END ------ \n\n"); err != nil {
            log.Println(err)
        }
        if _, err := f.WriteString(string(chapter)); err != nil {
            log.Println(err)
        }
    }

    return f.Close()
}

func WriteToEpub(book Book, path string) error {
    barcode := titleToFileWithExtensions(book.Title, ".epub")

    e := epub.NewEpub(book.Title)

    e.SetAuthor("Web Novels")

    chapterNum := 1
    for index, chapter := range book.Pages {
        fmt.Println("Writing URL : " + string(index))
        chapterStr := fmt.Sprintf(`Chapter %v` , chapterNum)
        content := cleanString(string(chapter))
        sectionBody := fmt.Sprintf(`<h1> %s</h1> <p> %s </p>`, chapterStr , content)

        chapterNum += 1
        _, err := e.AddSection(sectionBody, chapterStr,"","")

        if err != nil{
            fmt.Println("ERROR " , err )
        }
    }

    err := e.Write(path + "/" + barcode)

    if err != nil {
        return err
    }

    return nil
}