package main

import (
	"fmt"
	"strconv"

	"github.com/gocolly/colly"
)

type item struct {
	Name     string
	Price    string
	ImageUrl string
}

func main() {
	c := colly.NewCollector()
	c.OnHTML("div[class=card-v2]", func(h *colly.HTMLElement) {
		if h.Text != "" {
			fmt.Println(h.Text)
		} else {
			fmt.Println("No more content to scrape.")
			return
		}
	})
	i := 1
	for i <= 50 {
		fmt.Println("=====================" + "PAGINA" + strconv.Itoa(i) + "=====================")
		link := "https://www.emag.ro/televizoare/p" + strconv.Itoa(i) + "/c"
		c.Visit(link)
		i++
	}
}
