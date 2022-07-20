package scrapper

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

type StatusInvest struct{}

const (
	url       = "https://statusinvest.com.br"
	pathBDR   = "/bdrs/"
	pathACOES = "/acoes/"
)

func (s *StatusInvest) Fetch(stockName string) Stock {
	stock := Stock{Name: stockName}

	collector := colly.NewCollector()

	addValue(collector, &stock)

	collector.Visit(fullURL(stockName))

	return stock
}

func addValue(c *colly.Collector, stock *Stock) {
	c.OnHTML(
		"div[title='Valor atual do ativo'] .value",
		func(e *colly.HTMLElement) {
			textValue := strings.Replace(e.Text, ",", ".", -1)
			value, err := strconv.ParseFloat(textValue, 32)
			if err != nil {
				log.Fatal(err)
			}
			stock.Value = value
		},
	)
}

func fullURL(name string) string {
	path := pathACOES
	if ok, _ := regexp.MatchString("34", name); ok {
		path = pathBDR
	}
	return url + path + name
}
