package marketservice

import (
	"errors"

	bittrex "github.com/toorop/go-bittrex"
)

const (
	BITTREX_API_KEY    = "644a03c302374523869ba9e87421b3ae"
	BITTREX_API_SECRET = "MU72EQ6OTBDVGY65"
)

var bittrexapi = bittrex.New(BITTREX_API_KEY, BITTREX_API_SECRET)

type MarketService interface {
	RefreshAllTicks(int) error
	RefreshTick(int, string) error
	RefreshCurrencies() error
}

type MarketServiceImpl struct{}

func (mrs MarketServiceImpl) RefreshAllTicks(eid int) error {
	// TODO query all ticks available from DataService
	marketlist := [2]string{"BTC-NEO", "BTC-ETH"}

	if eid == 1 {
		for i := range marketlist {
			btrxticker, err := bittrexapi.GetTicker(marketlist[i])
			if err != nil {
				return errors.New("error in fetching tick:" + marketlist[i] + ", in eid:" + string(eid))
			}
			var ticker Ticker
			ticker.Bid = btrxticker.Bid
			ticker.Ask = btrxticker.Ask
			ticker.Last = btrxticker.Last

			mrs.RefreshTick(eid, marketlist[i], ticker)
		}
		if err != nil {
			return err
		}

		return nil
	} else {
		return errors.New("undefined exchange_id")
	}
}

func (mrs MarketServiceImpl) RefreshTick(eid int, mCode string) error {

	return nil
}
