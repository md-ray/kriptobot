package main

import (
	"net/http"
	"os"
	"time"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	msvc "github.com/saviourcat/kriptobot/market/service"
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
	svc.RefreshAllTicks(1)
	svc.RefreshAllTicks(2)

	/*
		if err != nil {
			fmt.Println(err)
		}
		tick, err := svc.GetCurrentTick(1, "BTC-NEO")
		fmt.Println(tick)
	*/
	// Logging Middleware

	// http handler

	getCurrentTickHandler := httptransport.NewServer(
		msvc.MakeGetCurrentTickEndpoint(svc),
		msvc.DecodeGetCurrentTickRequest,
		msvc.EncodeResponse,
	)
	getCurrentTick2Handler := httptransport.NewServer(
		msvc.MakeGetCurrentTick2Endpoint(svc),
		msvc.DecodeGetCurrentTick2Request,
		msvc.EncodeResponse,
	)

	// getMarketSummaryHandler := httptransport.NewServer(
	// 	dataservice.MakeGetMarketSummaryEndpoint(svc),
	// 	dataservice.DecodeGetMarketSummaryRequest,
	// 	dataservice.EncodeResponse,
	// )

	http.Handle("/getCurrentTick", getCurrentTickHandler)
	http.Handle("/getCurrentTick2", getCurrentTick2Handler)
	// http.Handle("/getMarketSummary", getMarketSummaryHandler)
	http.ListenAndServe(":9080", nil)

	go forever()
	select {}
}

func forever() {
	for {
		//fmt.Printf("%v+\n", time.Now())
		time.Sleep(time.Second)
	}
}
