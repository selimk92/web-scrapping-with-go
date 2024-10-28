package main

import (
	"github.com/gocolly/colly"
)

func main() {

	 c :0 colly.NewCollector(
		 colly.AllowedDomains("j2store.net"))

	 c.OnHTML("", func(e *colly.HTMLElement) {})
}

type item struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgURL string `json:"img_url"`
}
