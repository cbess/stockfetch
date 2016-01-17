package network

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "github.com/cbess/stockfetch/parser"
)

const (
    // yahooURLTemplate The URL template for the Yahoo CSV service 
    yahooURLTemplate = "http://real-chart.finance.yahoo.com/table.csv?s=%s&a=%d&b=%d&c=%d&d=%d&e=%d&f=%d&g=%s&ignore=.csv"
)

// DateComponents represents the Date components
type DateComponents struct {
    Month byte
    Day byte
    Year byte
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

// FetchContents fetches the contents from the specified url
func FetchContents(url string) (string, error) {
    // grab the remote contents
    res, err := http.Get(url)
    if err != nil {
        return "", err
    }
    
    // get the body contents
    defer res.Body.Close()
    contents, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return "", err
    }
    
    return string(contents), err
}

// FetchCSV fetches the CSV data from using the param values
func FetchCSV(p Params) [][]string, error {
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
    
    // get the contents
    url := fmt.Sprintf(
        yahooURLTemplate,
        p.Symbol,
        p.StartDate.Month,
        p.StartDate.Day,
        p.StartDate.Year,
        p.EndDate.Month,
        p.EndDate.Day,
        p.EndDate.Year,
        p.Interval
    )
    contents, err := FetchContents(url)
    
    return nil, nil
}