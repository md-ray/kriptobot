package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// type VIPTicker
type VIPAPI struct {
	VipTicker struct {
		High       string `json:"high"`
		Low        string `json:"low"`
		VOLBTC     string `json:"vol_btc"`
		VOLIDR     string `json:"vol_idr"`
		Last       string `json:"last"`
		Buy        string `json:"buy"`
		Sell       string `json:"sell"`
		ServerTime int64  `json:"server_time"`
	} `json:"ticker"`
}

func GetVIPTicker(c1 string, c2 string) Ticker {
	coin1code := strings.ToLower(c1)
	coin2code := strings.ToLower(c2)
	url := "http://vip.bitcoin.co.id/api/" + coin1code + "_" + coin2code + "/ticker"

	spaceClient := http.Client{
		Timeout: time.Second * 20, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	var apiret VIPAPI
	fmt.Printf("body=?\n", string(body))
	jsonErr := json.Unmarshal(body, &apiret)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Printf("hasil api = ?\n", apiret)

	a, _ := strconv.ParseFloat(apiret.VipTicker.Buy, 64)
	b, _ := strconv.ParseFloat(apiret.VipTicker.Sell, 64)
	c, _ := strconv.ParseFloat(apiret.VipTicker.Last, 64)
	ticker := Ticker{a, b, c, apiret.VipTicker.ServerTime}
	return ticker
}
