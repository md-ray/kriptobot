package marketservice

import (
	"time"

	"github.com/go-kit/kit/log"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   DataService
}

func (mw LoggingMiddleware) RegisterUser(uid int, username string) (err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "RegisterUser",
			"uid", uid,
			"username", username,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	err = mw.Next.RegisterUser(uid, username)
	return
}

func (mw LoggingMiddleware) GetMarketSummary(market string) (MarketSummary, error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "GetMarketSummary",
			"code", market,
			"took", time.Since(begin),
		)
	}(time.Now())
	ms, err := mw.Next.GetMarketSummary(market)
	return ms, err
}

func (mw LoggingMiddleware) UserSubscribeTick(uid string, tick string, subFlag bool) error {
	return nil
}

func (mw LoggingMiddleware) UpdateCurrentTick(string, float32) error {
	return nil
}
