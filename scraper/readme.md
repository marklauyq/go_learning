#Scrapper
A simple scraper that will search for the content based on the selector
and extract it out. It will than search for the "next" button and repeat
the cycle until it has retrieved all the chapters.

to use the scraper, you just need to import the package

```go
package main

import "github.com/marklauyq/go_learning/scraper"

func main() {
    scraper.Start(
        "https://url-here.com/path/here",
        ".selector #here",
        scraper.BoxnovelNextSelector,
        scraper.Book{Title:"Book Name"}
    )
}
```

`Start(string, string, scraper.nextSelector) scraper.Book`

The `nextSelector` is just a function that takes in the 
Chapter and search for the URL for the next chapter. 
Check out the `next_function.go` as an example and feel free to create
you own and pass it into the Start function

---
Future updates,
- I will need to move the book from scrapper over to the shelf library instead
    - We are doing this because we do not want a cyclic dependency
    - The book is more related to the shelf class which reads and writes to the book while the 
    scrapper class main job is to scrap the websites.   
