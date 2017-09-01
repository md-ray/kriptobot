package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

// type VIPTicker
type LunoAPI struct {
	Timestamp int64  `json:"timestamp"`
	Bid       string `json:"bid"`
	Ask       string `json:"ask"`
	Last      string `json:"last_trade"`
	Rolling   string `json:"rolling_24_hour_volume"`
	Pair      string `json:"pair"`
}

func GetLunoTicker(c1 string, c2 string) (Ticker, error) {
	//coin1code := strings.ToLower(c1)
	//coin2code := strings.ToLower(c2)
	var ret Ticker
	url := "https://api.mybitx.com/api/1/ticker?pair=XBTIDR"

	spaceClient := http.Client{
		Timeout: time.Second * 7,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
		return ret, err
	}

	// req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
		return ret, err
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	var apiret LunoAPI
	fmt.Printf("body=?\n", string(body))
	jsonErr := json.Unmarshal(body, &apiret)
	if jsonErr != nil {
		log.Fatal(jsonErr)
		return ret, err
	}

	fmt.Printf("hasil api = ?\n", apiret)

	a, _ := strconv.ParseFloat(apiret.Bid, 64)
	b, _ := strconv.ParseFloat(apiret.Ask, 64)
	c, _ := strconv.ParseFloat(apiret.Last, 64)

	ret = Ticker{a, b, c, apiret.Timestamp / 1000}
	return ret, nil
}
