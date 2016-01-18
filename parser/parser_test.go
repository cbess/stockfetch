package parser

import (
    "testing"
)

func TestTextToCSV(t *testing.T) {
    csvText := "Date,Open,High,Low,Close,Volume"
    collection := TextToCSV(csvText)
    
    t.Log(collection)
    
    if collection == nil {
        t.Errorf("No collection.")
    }
    
    if len(collection) == 0 {
        t.Errorf("Empty collection.")
    }
}

func TestCSVToStockData(t *testing.T) {
    // use test data from Yahoo CSV stock data
    csvText := `Date,Open,High,Low,Close,Volume,Adj Close
2016-01-15,692.289978,706.73999,685.369995,694.450012,3592400,694.450012
2016-01-14,705.380005,721.924988,689.099976,714.719971,2211900,714.719971
`
    data := CSVToStockData(TextToCSV(csvText))
    
    if data == nil {
        t.Errorf("No stock data.")
    }
    
    t.Log(data)
}