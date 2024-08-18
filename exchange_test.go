package exchange_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/asvvvad/exchange"
)

func TestExchange(t *testing.T) {
	ex := exchange.New("USD")
	f, err := ex.ConvertTo("EUR", 10)
	if err != nil {
		t.Fatal(err)
	}
	if f == nil {
		t.Fatal("nil value")
	}
}

func ExampleExchange() {
	// Set exchangerate.host ACCESS_KEY - either this way or via the EXCHANGERATE_ACCESS_KEY environment variable
	exchange.Client.AccessKey = os.Getenv("EXCHANGERATE_ACCESS_KEY")

	// Create a new Exchange instance and set USD as the base currency for the exchange rates and conversion
	ex := exchange.New("USD")
	// convert 10 USD to EUR
	fmt.Println(ex.ConvertTo("EUR", 10))

	// You can convert between 171 fiat and +6000 cryptocurrency aswelL!
	// convert 10 USD to BTC
	fmt.Println(ex.ConvertTo("BTC", 10))
	// convert 10 USD to EUR at 2012-12-12 (date must be in the format YYYY-MM-DD)
	fmt.Println(ex.ConvertAt("2012-12-12", "EUR", 10))

	// Get the available forex/fiat codes ([]string)
	forexCodes, _ := ex.ForexCodes()

	// Get the available crypto codes ([]string)
	// Warning: +6000
	cryptoCodes, _ := ex.CryptoCodes()
	fmt.Println(cryptoCodes)

	// Get the forex codes data, includes code and description.
	forexData, _ := ex.ForexData()

	// Get the crypto codes data, includes code and description.
	// Warning: +6000
	cryptoData, _ := ex.CryptoData()
	fmt.Println(cryptoData["BTC"]["code"])

	// loop through the forex cpdes
	for _, code := range forexCodes {
		// print the forex codes data in the format: USD: US Dollar
		fmt.Println(code+":", forexData[code]["description"])
	}

	// Change the base currency to euro
	ex.SetBase("EUR")
	// Get the latest exchange rates with all currencies (Base is EUR)
	fmt.Println(ex.LatestRatesAll())

	// Get the latest rates again, this time it will be loaded from in-memory cache
	// Cache last till midnight GMT because it's the time exchangerate.host update the rates
	fmt.Println(ex.LatestRatesAll())
	// disable caching
	ex.SetCache(false)

	// Get the latest rates with multiple currencies, not all (USD and JPY only)
	fmt.Println(ex.LatestRatesMultiple([]string{"USD", "JPY"}))

	// Get the exchange rates at 2012-12-12 but only with USD
	fmt.Println(ex.HistoricalRatesSingle("2012-12-12", "USD"))

	// Get historical rates between 2012 12 10 and 2012 12 12 for JPY and GBP
	fmt.Println(ex.TimeseriesMultiple("2012-12-10", "2012-12-12", []string{"USD", "JPY"}))

	// Get the fluctuation between 2012 12 10 and 2012 12 12 with USD
	fluctuation, _ := ex.FluctuationSingle("2012-12-10", "2012-12-12", "USD")
	// Print the change
	fmt.Println(fluctuation["change"])
}
