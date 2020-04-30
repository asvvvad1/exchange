# exchange
Go library for current and historical exchange rates and currency conversion using the new [Free foreign exchange rates API](https://exchangerate.host/#/) by [arzzen](https://github.com/arzzen/) ([github](https://github.com/arzzen/exchangerate.host))

## Features:
- Currency convertion, historical & current exchange rates, timeseries and fluctuations
- No authentication/token needed
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
	// Create a new Exchange instance and set USD as the base currency for the exchange rates and conversion
	ex := exchange.New("USD")
	// convert 10 USD to EUR
	fmt.Println(ex.ConvertTo("EUR", 10))
	// convert 10 USD to EUR at 2012-12-12 (date must be in the format YYYY-MM-DD)
	fmt.Println(ex.ConvertAt("2012-12-12", "EUR", 10))

	// Get the available symbols ([]string)
	symbols, _ := ex.Symbols()
	// Get the symbols data, includes code, description and symbol (such as $ for USD)
	symbolsData, _ := ex.SymbolsData()
	// loop through the symbols
	for _, symbol := range symbols {
		// print the symbols data in the format: USD US Dollar $100
		fmt.Println(symbol+":", symbolsData[symbol]["description"], symbolsData[symbol]["symbol"]+string(100))
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
- Symbols
- - `[]string{symbols}`, error
- SymbolsData
- - `map[symbol]map[
    code
    description
    symbol
    dec
    hex
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
