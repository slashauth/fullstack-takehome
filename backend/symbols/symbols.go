package symbols

import (
	"time"

	"github.com/getdebrief/fullstack-takehome/graph/model"
)

var availableSymbolNames = []string{
	"GOOG",
	"AAPL",
	"MSFT",
	"IBM",
	"F",
	"GME",
	"BB",
	"BBBY",
	"TSLA",
	"SPY",
}

func GetAvailableSymbols() []string {
	return availableSymbolNames
}

// TODO: Get the price history for the provided ID
func GetPriceHistory(id string) ([]*model.TradingSession, error) {
	return []*model.TradingSession{
		{
			Time:  time.Now().Add(time.Minute * -1),
			Open:  10000,
			High:  10050,
			Low:   9990,
			Close: 10010,
		},
		{
			Time:  time.Now().Add(time.Minute * -2),
			Open:  9991,
			High:  10050,
			Low:   9990,
			Close: 10000,
		},
	}, nil
}

// TODO: Updates price history across all symbols
func UpdatePriceHistory(id string) (*model.TradingSession, error) {
	return &model.TradingSession{
		Time:  time.Now(),
		Open:  10000,
		High:  10100,
		Low:   9000,
		Close: 10000,
	}, nil
}
