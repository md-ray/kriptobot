package main

import (
	"net/http"
	"os"
	"time"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/jasonlvhit/gocron"
	"github.com/kelseyhightower/envconfig"
	msvc "github.com/saviourcat/kriptobot/market/service"
)

var svc msvc.MarketService

func repeatRefresh() {
	svc.RefreshAllTicks(1)
	svc.RefreshAllTicks(2)
	svc.RefreshAllTicks(3)
}

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	// init config
	var config msvc.Config
	err := envconfig.Process("kmarket", &config)
	if err != nil {
		panic(err)
	}

	// init serviceimpl
	msvc.ServiceImplInit(config)

	//  Data Service
	svc = msvc.MarketServiceImpl{}
	svc = msvc.LoggingMiddleware{logger, svc}

	//lunotick := msvc.GetLunoTicker("1", "2")
	//fmt.Printf("hasil = ", lunotick.Ask)

	// repeatRefresh()
	gocron.Every(1).Minute().Do(repeatRefresh)
	gocron.Start()
	// svc.Init()
	/*
		var ticker msvc.Ticker
		ticker.Last = 1
		ticker.Ask = 2
		ticker.Bid = 3
		err := svc.RefreshTick(1, "BTC-NEO", ticker)
	*/

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
