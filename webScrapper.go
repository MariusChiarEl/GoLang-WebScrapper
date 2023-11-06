/*
	Author: Marius Enache-Stratulat
	
	This is my bachelor thesis project.
	
	This application has the purpose of extracting specific data from a web page.
	
	In this instance, this application will scrape essential product data from the
	eMAG online store by providing the link to a product category (such as mobile phones)
	
	We will store the output in an excel file with one sheet for each store page. The first
	row will contain the name of each product, while the second row will contain its price.
*/
package main

import (
	"fmt"

	"strconv"
	"strings"
	// import functions that work with strings

	"github.com/gocolly/colly"    // for scrapping HTML files
	"github.com/xuri/excelize/v2" // for working with excel files
)

type item struct {
	Name  string
	Price string
}

func extract_product_data(product string, out *excelize.File, index int, row int) {
	lines := strings.Split(product, "\n")
	for i := 0; i < len(lines); i++ {
		if strings.Contains(lines[i], "Telefon mobil") { // telefon mobil is the romanian for mobile phone
			name := strings.Split(lines[i], "Telefon mobil ")
			out.SetCellValue(out.GetSheetName(index), "A"+strconv.Itoa(row), name[1])
		}
		if strings.Contains(lines[i], "Lei") && !strings.Contains(lines[i], "PRP") {
			pret := strings.Split(lines[i], " ")
			out.SetCellValue(out.GetSheetName(index), "B"+strconv.Itoa(row), pret[0])
		}
	}
	fmt.Println("Product " + strconv.Itoa(row-1) + " added.")
}

func main() {
	c := colly.NewCollector()    // will contain the HTML code from the link given
	output := excelize.NewFile() // will contain the output
	row := 2
	var i int
	c.OnHTML("div[class=card-v2]", func(h *colly.HTMLElement) {
		// extracts all the divs related to listed products
		if h.Text != "" {
			extract_product_data(h.Text, output, i, row)
			// extracts product data from the div
		} else {
			fmt.Println("No more content to scrape.")
			return
		}
		row++
	})
	for i = 1; i <= 5; i++ {
		// will iterate through the first 5 pages of that product category
		row = 2
		index, err := output.NewSheet("Page " + strconv.Itoa(i))
		// each store page will have its own excel sheet
		if err != nil {
			fmt.Println(err)
			return
		}
		output.SetActiveSheet(index)
		output.SetCellValue(output.GetSheetName(index), "A1", "product name")
		output.SetCellValue(output.GetSheetName(index), "B1", "price")
		fmt.Println("=====================" + "PAGE" + strconv.Itoa(i) + "=====================")
		link := "https://www.emag.ro/telefoane-mobile/p" + strconv.Itoa(i) + "/c"
		err = c.Visit(link)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	err := output.DeleteSheet(output.GetSheetName(0))
	if err != nil {
		fmt.Println(err)
		return
	}

	err = output.SaveAs("output.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
}
