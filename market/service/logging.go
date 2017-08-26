package service

import (
	"time"

	"github.com/go-kit/kit/log"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   MarketService
}

func (mw LoggingMiddleware) RefreshAllTicks(eid int) (err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "RefreshAllTicks",
			"eid", eid,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	err = mw.Next.RefreshAllTicks(eid)
	return err
}

func (mw LoggingMiddleware) RefreshTick(eid int, mCode string, ticker Ticker) (err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "RefreshTick",
			"eid", eid,
			"mCode", mCode,
			"ticker", ticker,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	err = mw.Next.RefreshTick(eid, mCode, ticker)
	return err
}

func (mw LoggingMiddleware) GetCurrentTick(eid int, mCode string) (tick Ticker, err error) {
	// fmt.Println("masuk log")
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "GetCurrentTick",
			"eid", eid,
			"mCode", mCode,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	ticker, err := mw.Next.GetCurrentTick(eid, mCode)
	return ticker, err
}

func (mw LoggingMiddleware) GetCurrentTick2(eid int, c1 string, c2 string) (tick Ticker, err error) {
	// fmt.Println("masuk log")
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "GetCurrentTick",
			"eid", eid,
			"c1", c1,
			"c2", c2,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	ticker, err := mw.Next.GetCurrentTick2(eid, c1, c2)
	return ticker, err
}

/*
func (mw LoggingMiddleware) RefreshCurrencies() (err error) {
	return nil
}
*/
