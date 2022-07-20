package main

import (
	"log"
	"os"
	"sync"

	"github.com/yuripac/invest-scrapper/scrapper"
)

var (
	scr = scrapper.StatusInvest{}
	wg  = sync.WaitGroup{}
)

const (
	walletFilename = "wallet.yml"
	stocksFilename = "stocks.yml"
)

func main() {
	config, err := scrapper.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	file, _ := os.OpenFile(config.FilePath(stocksFilename), os.O_CREATE|os.O_RDWR, os.ModePerm)
	defer file.Close()

	repo := scrapper.YMLRepo{File: file}

	stockNames, err := scrapper.GetWallet(config.FilePath(walletFilename))
	if err != nil {
		log.Fatal(err)
	}

	wg.Add(len(stockNames))
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
