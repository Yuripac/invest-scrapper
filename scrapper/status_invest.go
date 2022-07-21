package scrapper

import (
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

func (s *StatusInvest) FetchValue(name string) (float64, error) {
	collector := colly.NewCollector()

	var value float64
	var err error
	collector.OnHTML(
		"div[title='Valor atual do ativo'] .value",
		func(e *colly.HTMLElement) {
			textValue := strings.Replace(e.Text, ",", ".", -1)

			value, err = strconv.ParseFloat(textValue, 32)
		},
	)

	collector.Visit(fullURL(name))
	return value, err
}

func fullURL(name string) string {
	path := pathACOES
	if ok, _ := regexp.MatchString("34", name); ok {
		path = pathBDR
	}
	return url + path + name
}
