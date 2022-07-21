package main

import (
	"log"
	"os"

	"github.com/yuripac/invest-scrapper/scrapper"
)

var (
	scr = scrapper.StatusInvest{}
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

	wallet, err := scrapper.InitWallet(config.FilePath(walletFilename))
	if err != nil {
		log.Fatal(err)
	}

	wallet.FetchStockValues(&scr)
	for _, stock := range wallet.Stocks {
		if n, _ := repo.UpdateMaxValue(*stock); n > 0 {
			stock.SysNotify()
		}
	}
}
