package network

/*
refs:
	http://real-chart.finance.yahoo.com/table.csv?s=GOOG&a=07&b=16&c=2015&d=00&e=17&f=2016&g=d&ignore=.csv
*/

import (
	"fmt"
	"github.com/bmuller/arrow/lib"
	"github.com/cbess/stockfetch/parser"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	// yahooURLTemplate The URL template for the Yahoo CSV service
	yahooURLTemplate = "http://real-chart.finance.yahoo.com/table.csv?s=%s&a=%d&b=%d&c=%d&d=%d&e=%d&f=%d&g=%s&ignore=.csv"
)

// DateComponents represents the Date components
type DateComponents struct {
	Month time.Month
	Day   int
	Year  int
}

// Params represents the parameters for a CSV query
type Params struct {
	// Symbol the stock symbol
	Symbol string
	// StartDate the start date
	StartDate DateComponents
	// EndDate the end date
	EndDate DateComponents
	// Interval the interval of the data
	// d = daily, w = weekly, m = monthly
	Interval string
}

func (dc DateComponents) isValid() bool {
	if dc.Day < 1 || dc.Day > 31 {
		return false
	}

	if dc.Month < 1 || dc.Month > 12 {
		return false
	}

	// safe to assume no stock data before this year
	if dc.Year < 1900 {
		return false
	}

	return true
}

func (dc DateComponents) String() string {
	str := fmt.Sprintf("%02d/%02d/%04d", dc.Month, dc.Day, dc.Year)
	if dc.Year == 0 {
		return fmt.Sprintf("(invalid date: %s)", str)
	}
	return str
}

// DateComponentsFromString converts the specified string into DateComponents
//
// text - the date string, example: 8-16-2015
func DateComponentsFromString(text string) DateComponents {
	// assumes m-d-Y format
	if strings.Contains(text, "-") {
		date, err := arrow.CParse("%m-%d-%Y", text)
		if err != nil {
			return DateComponents{}
		}

		return DateComponents{
			Month: date.Month(),
			Day:   date.Day(),
			Year:  date.Year(),
		}
	}
	return DateComponents{}
}

// FetchContents fetches the contents from the specified url
func FetchContents(url string) (string, error) {
	// grab the remote contents
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	// get the body contents
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// check status code
	if res.StatusCode >= 300 {
		return "", fmt.Errorf("Bad status code (%d): %s", res.StatusCode, url)
	}

	return string(body), err
}

// FetchCSV fetches the CSV data from using the param values
func FetchCSV(p Params) ([][]string, error) {
	// add defaults
	if p.Interval == "" {
		p.Interval = "d"
	}

	if !p.StartDate.isValid() {
		return nil, fmt.Errorf("Invalid start date")
	}

	if !p.EndDate.isValid() {
		return nil, fmt.Errorf("Invalid end date")
	}

	// get the remote contents
	url := fmt.Sprintf(
		yahooURLTemplate,
		p.Symbol,
		p.StartDate.Month-1, // adjust month for query (zero based)
		p.StartDate.Day,
		p.StartDate.Year,
		p.EndDate.Month-1,
		p.EndDate.Day,
		p.EndDate.Year,
		p.Interval,
	)
	contents, err := FetchContents(url)
	if err != nil {
		return nil, err
	}

	csv := parser.TextToCSV(contents)
	return csv, nil
}

// FetchStockData fetches the stock data
func FetchStockData(p Params) ([]parser.StockData, error) {
	csvData, err := FetchCSV(p)
	if err != nil {
		return nil, err
	}

	data := parser.CSVToStockData(csvData)
	return data, nil
}
