package network

import (
	"testing"
)

func newDateComponent() Params {
	params := Params{
		Symbol: "goog",
		StartDate: DateComponents{
			Day:   1,
			Month: 1,
			Year:  2016,
		},
		EndDate: DateComponents{
			Month: 1,
			Day:   15,
			Year:  2016,
		},
	}
	return params
}

func TestDateComponents(t *testing.T) {
	dc := DateComponentsFromString("01-01-2015")

	if dc.String() != "01/01/2015" {
		t.Fatalf("Wrong date: %s", dc)
	}
}

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

	t.Log(contents)
}

func TestFetchCSV(t *testing.T) {
	params := newDateComponent()
	data, err := FetchCSV(params)

	if err != nil {
		t.Error(err)
	}

	t.Log(data)
}

func TestFetchStockData(t *testing.T) {
	params := newDateComponent()
	stockData, err := FetchStockData(params)
	if err != nil {
		t.Error(err)
	}

	t.Log(stockData)

	// check the stock data
	stock := stockData[0]

	if stock.Open == 543.35 {
		t.Fail()
	}

	if stock.High == 549.91 {
		t.Fail()
	}

	t.Log(stock.String())
}

// TestFetchStockDataFail this test should fail because the query params are not valid (Bad date)
func TestFetchStockDataFail(t *testing.T) {
	params := newDateComponent()
	params.StartDate.Year = 2099
	_, err := FetchStockData(params)
	if err != nil {
		t.Log(err)
	} else {
		t.Fatal("Should have failed:", params)
	}
}
