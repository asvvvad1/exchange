# exchange
[![Go Report Card](https://goreportcard.com/badge/github.com/asvvvad/exchange)](https://goreportcard.com/report/github.com/asvvvad/exchange)

Go library for current & historical exchange rates, forex & crypto currency conversion, fluctuation, and timeseries using the new [Free foreign exchange rates API](https://exchangerate.host/#/) by [arzzen](https://github.com/arzzen/) ([github](https://github.com/arzzen/exchangerate.host))

## Features:
- Currency conversion, historical & current exchange rates, timeseries and fluctuations
- Select any base currency
- 171 forex currency and 6000+ cryptocurrency!
- Caching (optional, default) using [go-cache](https://github.com/patrickmn/go-cache)
- Easy to use:

## Usage:

> #### `go get -u github.com/asvvvad/exchange` 

```go
package main

import (
	"fmt"

	"github.com/asvvvad/exchange"
)

func main() {
  // Set exchangerate.host ACCESS_KEY - either this way or via the EXCHANGERATE_ACCESS_KEY environment variable
  exchange.Client.AccessKey = "my-access-key"
  
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

	// Get the forex codes data, includes code and description.
	forexData, _ := ex.ForexData()

	// Get the crypto codes data, includes code and description.
	// Warning: +6000
	cryptoData, _ := ex.CryptoData()

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

```

### Results returned by each method:
- ConvertTo, ConvertAt, HistoricalRatesSingle, LatestRatesSingle
- - `big.Float`, error
- LatestRatesAll, LatestRatesMultiple, HistoricalRatesAll, HistoricalRatesMultiple:
- - `map[symbol(string)]rate(big.Float)`
- ForexCodes
- - `[]string{codes}`, error
- ForexData
- - `map[symbol]map[
    code
    description
]string`, error
- CryptoCodes
- - `[]string{codes}`, error
- CryptoData
- - `map[symbol]map[
    symbol
    name
]string`, error
- FluctuationAll, FluctuationMultiple,
- - `map[symbol]map[
    start_rate
    end_rate
    change
    change_pct
]*big.Float`, error
- FluctuationSingle
- - `map[
    start_rate
    end_rate
    change
    change_pct
]*big.Float`, error

- TimeseriesAll, TimeseriesMultiple
- - `map[date]map[symbols]*big.Float`, error
- TimeseriesSingle
- - `map[date]map[symbol]*big.Float`, error

> ## Notes:

- You can use All, Multiple, Single with all of LatestRates, HistoricalRates, Timeseries and Fluctuation.
- Oldest date for historical rates and conversion is 1999-01-04
- Maximum allowed timeframe for Timeseries is 365 days

#### Input validation with the appropriate errors for all methods is provided to help debug

#### Any help and contribution is welcome!
This is my first Go library and I had trouble with JSON parsing (and I still do, didn't use bitly/simplejson to reduce dependencies) Theres a lot of room for improvement
