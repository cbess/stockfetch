package main

import (
	"fmt"
	"github.com/cbess/stockfetch/network"
	// flag "github.com/spf13/pflag"
)

func main() {
	// http://real-chart.finance.yahoo.com/table.csv?s=GOOG&a=07&b=16&c=2015&d=00&e=17&f=2016&g=d&ignore=.csv
	params := network.Params{
		Symbol:    "",
		StartDate: network.DateComponents{},
		EndDate:   network.DateComponents{},
	}

	stock, err := network.FetchStockData(params)
	if err != nil {
		fmt.Printf("Unable to fetch stock: %s", err)
		return
	}

	fmt.Print(stock)
}
