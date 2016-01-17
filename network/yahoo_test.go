package network

import (
    "testing"
)

func TestFetchContents(t *testing.T) {
    url := "https://api.ipify.org?format=json"
    t.Log("Fetching from URL:", url)
    
    contents, err := FetchContents(url)
    
    if err != nil {
        t.Fatal(err)
    }
    
    if contents == "" {
        t.Fatalf("No contents")
    }
}

func TestFetchYahooCSV(t *testing.T) {
    params := Params{
        Symbol: "goog",
        StartDate: DateComponents{
            Day: 1,
            Month: 1,
            Year: 2015,
        },
        EndDate: DateComponents{
            Month: 1,
            Day: 15,
            Year: 2015
        }
    }
    
    data, err := FetchCSV(params)
}