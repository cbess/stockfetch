package parser

import (
    "strings"
    "strconv"
    "github.com/bmuller/arrow/lib"
)

// StockData Represents the stock data for the interval
type StockData struct {
    Date arrow.Arrow
    Open float32
    Close float32
    High float32
    Low float32
    Volume int32
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
    // parse each line
    for idx, line := range lines {
        // ignore the first line
        if idx == 0 {
            continue
        }
        
        // get data from each line
        parts := strings.Split(line, ",")
        collection = append(collection, parts)
    }
    return collection
}

// CSVToStockData converts the string array to StockData
func CSVToStockData(data [][]string) []StockData {
    // build the stock data collection
    collection := []StockData{}
    for _, entry := range data {
        if len(entry) < 5 {
            continue
        }
        
        // parse stock date
        date, err := arrow.CParse("%Y-%m-%d", entry[0])
        if err != nil {
            return nil
        }
        
        item := StockData{
            Date: date,
            Open: parseFloat32(entry[1]),
            High: parseFloat32(entry[2]),
            Low: parseFloat32(entry[3]),
            Close: parseFloat32(entry[4]),
            Volume: parseInt32(entry[5]),
        }
        
        collection = append(collection, item)
    }
    return collection
}
