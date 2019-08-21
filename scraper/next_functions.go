package scraper

type nextSelector func(*Chapter) string

// WuxiaWorldNextSelector is the function that can be used to search for the next chapter in wuxia world
func WuxiaWorldNextSelector(c *Chapter) string {
    path, status := c.Find(".next a").Attr("href")
    if !status {
        return ""
    }
    return `https://www.wuxiaworld.com` + path
}

// Boxnovel is the function for searching the URL for the next chapter in  Boxnovel.com.
func BoxnovelNextSelector(c *Chapter)  string {
    path, status := c.Find(".nav-next a").Attr("href")
    if !status {
        return ""
    }
    return path
}
