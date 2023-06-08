package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	fName := "data.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Could not create file, err: %q", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector()
	c.OnHTML("table#customers", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(i int, h *colly.HTMLElement) {
			writer.Write([]string{
				h.ChildText("td:nth-child(1)"),
				h.ChildText("td:nth-child(2)"),
				h.ChildText("td:nth-child(3)"),
			})
		})
		fmt.Println("Finished successfully")
	})
	c.Visit("https://www.w3schools.com/html/html_tables.asp")
}
