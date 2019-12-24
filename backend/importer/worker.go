package importer

import (
	"github.com/cmckee-dev/go-alpha-vantage"
	"github.com/shindakun/envy"
	"log"
)

/**
This function imports historical stock data via the alphavantage api. The data retrived
only contains end of day values, so no intraday.
@param Symbol The Symbol of the equity for which the histrical data should be retrived
@param output The channel to which the result should be written.
**/
func GetStockData(symbol string, output chan av.TimeSeriesValue) {
	apiKey, err := envy.Get("API_KEY")

	client := av.NewClient(apiKey)
	timeSeries, err := client.StockTimeSeries(av.TimeSeriesDaily, symbol)
	if err != nil {
		log.Fatalln("Could not retrieve stock data from alpha vantage")
	}

	log.Println("Got daily stock data from alphavantage")
	for _, tp := range timeSeries {
		output <- *tp
	}
}

func GetIntraDayStockData(symbol string, output chan av.TimeSeriesValue) {
	apiKey, err := envy.Get("API_KEY")
	if err != nil {
		log.Fatalln("Could not retrieve api key from environment")
	}

	client := av.NewClient(apiKey)
	timeSeries, err := client.StockTimeSeriesIntraday(av.TimeIntervalFifteenMinute, apiKey)
	if err != nil {
		log.Fatalln("Could not retrieve time series from alphavantage")
	}

	for _, tp := range timeSeries {
		output <- *tp
	}

}
