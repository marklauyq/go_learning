package scrapper

// WuxiaWorldNextSelector is the function that can be used to search for the next chapter.
func WuxiaWorldNextSelector(c *Chapter) string {
	path , status := c.Find(".next a").Attr("href")
	if !status {
		return ""
	}
	return `https://www.wuxiaworld.com` + path
}


