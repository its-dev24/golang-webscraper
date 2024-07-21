package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {


	c := colly.NewCollector(colly.AllowedDomains("www.moneycontrol.com"))

	go c.OnHTML("td.tdred", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})
	go c.OnHTML("td.tbl_redtxt", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})
	c.Visit("https://www.moneycontrol.com/stocksmarketsindia/")
}
