package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-kit/kit/log"
	msvc "github.com/saviourcat/kriptobot/marketservice"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	//  Data Service
	var svc msvc.MarketService
	svc = msvc.MarketServiceImpl{}
	svc = msvc.LoggingMiddleware{logger, svc}
	// svc.Init()
	/*
		var ticker msvc.Ticker
		ticker.Last = 1
		ticker.Ask = 2
		ticker.Bid = 3
		err := svc.RefreshTick(1, "BTC-NEO", ticker)
	*/
	err := svc.RefreshAllTicks(1)
	if err != nil {
		fmt.Println(err)
	}

	// Logging Middleware

	// http handler
	/*
		registerUserHandler := httptransport.NewServer(
			dataservice.MakeRegisterUserEndpoint(svc),
			dataservice.DecodeRegisterUserRequest,
			dataservice.EncodeResponse,
		)
		getMarketSummaryHandler := httptransport.NewServer(
			dataservice.MakeGetMarketSummaryEndpoint(svc),
			dataservice.DecodeGetMarketSummaryRequest,
			dataservice.EncodeResponse,
		)

		http.Handle("/registerUser", registerUserHandler)
		http.Handle("/getMarketSummary", getMarketSummaryHandler)
		http.ListenAndServe(":9080", nil)


	*/
	go forever()
	select {}
}

func forever() {
	for {
		fmt.Printf("%v+\n", time.Now())
		time.Sleep(time.Second)
	}
}
