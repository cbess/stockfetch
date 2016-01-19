package main

import (
	"fmt"
	"github.com/cbess/stockfetch/network"
	flag "github.com/spf13/pflag"
	"strings"
)

var (
	fSymbol    = flag.StringP("symbol", "s", "", "The stock symbol. Ex: GOOG, AAPL, NFLX")
	fStartDate = flag.StringP("start", "b", "", "The start/beginning date. Ex: 08-16-2015")
	fEndDate   = flag.StringP("end", "e", "", "The end date. Ex: 01-01-2016")
)

func main() {
	flag.Parse()

	if *fSymbol == "" {
		fmt.Println("Symbol is required: GOOG, AAPL")
		return
	}

	params := network.Params{
		Symbol:    strings.ToUpper(*fSymbol),
		StartDate: network.DateComponentsFromString(*fStartDate),
		EndDate:   network.DateComponentsFromString(*fEndDate),
	}

	fmt.Printf("Fetching: %s %s-%s\n", params.Symbol, params.StartDate, params.EndDate)

	stock, err := network.FetchStockData(params)
	if err != nil {
		fmt.Println("Unable to fetch stock:", err)
		fmt.Println("args:", fSymbol, fStartDate, *fEndDate)
		return
	}

	fmt.Println(stock)
}
