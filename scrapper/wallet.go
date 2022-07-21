package scrapper

import (
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type Wallet struct {
	Owner  string
	Stocks []*Stock
}

func InitWallet(path string) (*Wallet, error) {
	content, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var stockNames []string
	err = yaml.Unmarshal(content, &stockNames)
	if err != nil {
		return nil, err
	}

	var stocks []*Stock
	for _, name := range stockNames {
		stocks = append(stocks, &Stock{Name: name})
	}

	wallet := Wallet{}
	wallet.Stocks = stocks

	return &wallet, nil
}

func (w *Wallet) FetchStockValues(scrapper Scrapper) {
	wg := sync.WaitGroup{}
	wg.Add(len(w.Stocks))
	for _, stock := range w.Stocks {
		go func(stock *Stock) {
			defer wg.Done()

			value, err := scrapper.FetchValue(stock.Name)
			if err != nil {
				log.Fatal(err)
			}

			stock.Value = value
		}(stock)
	}
	wg.Wait()
}
