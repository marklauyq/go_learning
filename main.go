package main

import (
	"fmt"
	s "simple_scrapper/scrapper"
)

func main() {
	fmt.Println(s.Start(
		"https://www.wuxiaworld.com/novel/dragon-prince-yuan/yz-chapter-1",
		"#chapter-content",
			s.WuxiaWorldNextSelector,
		),
	)
}

