package main

import (
	"log"
	"os"
	"sync"

	"github.com/yuripac/invest-scrapper/scrapper"
)

var (
	scr = scrapper.StatusInvest{}
)

func main() {
	config, err := scrapper.InitConfig("wallet.yml", "stocks.yml")
	if err != nil {
		log.Fatal(err)
	}

	stockNames, err := scrapper.GetWallet(config.Dir + "/" + config.WalletFilename)
	if err != nil {
		log.Fatal(err)
	}

	stocksPath := config.Dir + "/" + config.StocksFilename
	file, _ := os.OpenFile(stocksPath, os.O_CREATE|os.O_RDWR, os.ModePerm)
	defer file.Close()

	wg := sync.WaitGroup{}
	wg.Add(len(stockNames))
	repo := scrapper.YMLRepo{File: file}
	for _, name := range stockNames {
		go func(name string) {
			defer wg.Done()

			stock := scr.Fetch(name)
			if n, _ := repo.UpdateMaxValue(stock); n > 0 {
				stock.SysNotify()
			}
		}(name)
	}

	wg.Wait()
}
