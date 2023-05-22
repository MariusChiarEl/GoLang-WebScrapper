package main

import (
	"fmt"

	"strconv"
	"strings"

	// pentru functiile care prelucreaza sirurile de caractere

	"github.com/gocolly/colly"    // pentru colectarea datelor din pagini HTML
	"github.com/xuri/excelize/v2" // pentru crearea si manipularea de documente Excel
)

type item struct {
	Name  string
	Price string
}

func extrageDateProdus(produs string, out *excelize.File, index int, rand int) {
	lines := strings.Split(produs, "\n")
	for i := 0; i < len(lines); i++ {
		if strings.Contains(lines[i], "Telefon mobil") {
			nume := strings.Split(lines[i], "Telefon mobil ")
			out.SetCellValue(out.GetSheetName(index), "A"+strconv.Itoa(rand), nume[1])
		}
		if strings.Contains(lines[i], "Lei") && !strings.Contains(lines[i], "PRP") {
			pret := strings.Split(lines[i], " ")
			out.SetCellValue(out.GetSheetName(index), "B"+strconv.Itoa(rand), pret[0])
		}
	}
	fmt.Println("Product " + strconv.Itoa(rand-1) + " added.")
}

func main() {
	c := colly.NewCollector()    // c va extrage codul HTML din paginile date
	output := excelize.NewFile() // pregatim fisierul excel
	rand := 2
	var i int
	c.OnHTML("div[class=card-v2]", func(h *colly.HTMLElement) {
		// secventa aceasta va rula pentru fiecare produs gasit
		if h.Text != "" {
			extrageDateProdus(h.Text, output, i, rand)
			// prelucram "div-ul" fiecarui produs
		} else {
			fmt.Println("No more content to scrape.")
			return
		}
		rand++
	})
	for i = 1; i <= 5; i++ {
		// secventa aceasta va rula pentru fiecare pagina
		rand = 2
		index, err := output.NewSheet("Pagina " + strconv.Itoa(i))
		// pentru fiecare pagina WEB se creaza o pagina in fisierul excel
		if err != nil {
			fmt.Println(err)
			return
		}
		output.SetActiveSheet(index)
		output.SetCellValue(output.GetSheetName(index), "A1", "Nume Produs")
		output.SetCellValue(output.GetSheetName(index), "B1", "Pret")
		fmt.Println("=====================" + "PAGINA" + strconv.Itoa(i) + "=====================")
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
