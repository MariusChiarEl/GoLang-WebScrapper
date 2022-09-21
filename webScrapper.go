package main

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/gocolly/colly"
)

type item struct {
	Name     string
	Price    string
}

func arrangeText(t string){
	lines := strings.Split(t, "\n")
	//product := item
	//product.Name = lines[4]
	//product.Price = lines[16]
	for i:= 0; i<len(lines); i++{
		if strings.Contains(lines[i], "Televizor"){
			fmt.Println("Nume produs: " + lines[i])
		} else if strings.Contains(lines[i], "Lei") && !strings.Contains(lines[i], "PRP"){
			fmt.Println("Pret produs: " + lines[i])
		}
	}
	fmt.Println("-----")
}

func main() {
	c := colly.NewCollector()
	c.OnHTML("div[class=card-v2]", func(h *colly.HTMLElement) {
		if h.Text != "" {
			arrangeText(h.Text)
			//fmt.Println(h.Text)
		} else {
			fmt.Println("No more content to scrape.")
			return
		}
	})
	i := 1
	for i <= 5 {
		fmt.Println("=====================" + "PAGINA" + strconv.Itoa(i) + "=====================")
		link := "https://www.emag.ro/televizoare/p" + strconv.Itoa(i) + "/c"
		c.Visit(link)
		i++
	}
}
