package parser

/*
refs:
	https://github.com/bmuller/arrow
	http://man7.org/linux/man-pages/man3/strftime.3.html
*/

import (
	"fmt"
	"github.com/bmuller/arrow/lib"
	"strconv"
	"strings"
)

// minCSVRowLen the minimum size of the CSV row
const minCSVRowLen = 5

// StockData Represents the stock data for the interval
type StockData struct {
	Date   arrow.Arrow
	Open   float32
	Close  float32
	High   float32
	Low    float32
	Volume int32
}

func (sd *StockData) String() string {
	str := fmt.Sprintf(
		"Date: %s | Open: %.2f | Close: %.2f",
		sd.Date.CFormat("%b %d, %Y"),
		sd.Open,
		sd.Close,
	)
	return str
}

func parseFloat32(value string) float32 {
	if v, e := strconv.ParseFloat(value, 32); e == nil {
		return float32(v)
	}
	return 0
}

func parseInt32(value string) int32 {
	if v, e := strconv.ParseInt(value, 10, 32); e == nil {
		return int32(v)
	}
	return 0
}

// TextToCSV converts the specified text to an array of csv values
func TextToCSV(text string) [][]string {
	lines := strings.Split(text, "\n")
	collection := [][]string{}

	// should we skip the header
	hasHeader := false
	if len(lines) > 1 {
		// first row should be the header
		hasHeader = true
	}

	// parse each line
	for idx, line := range lines {
		// ignore the first row
		if hasHeader {
			if idx == 0 {
				continue
			}
		}

		// get data from each line
		parts := strings.Split(line, ",")
		if len(parts) < minCSVRowLen {
			// ignore lines that don't have enough columns
			continue
		}
		collection = append(collection, parts)
	}
	return collection
}

// CSVToStockData converts the string array to StockData
func CSVToStockData(data [][]string) []StockData {
	// build the stock data collection
	collection := []StockData{}
	for _, entry := range data {
		if len(entry) < minCSVRowLen {
			continue
		}

		// parse stock date
		date, err := arrow.CParse("%Y-%m-%d", entry[0])
		if err != nil {
			return nil
		}

		item := StockData{
			Date:   date,
			Open:   parseFloat32(entry[1]),
			High:   parseFloat32(entry[2]),
			Low:    parseFloat32(entry[3]),
			Close:  parseFloat32(entry[4]),
			Volume: parseInt32(entry[5]),
		}

		collection = append(collection, item)
	}
	return collection
}
