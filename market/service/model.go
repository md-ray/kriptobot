package service

// data model
type MarketSummary struct {
	High    float64
	Low     float64
	Volume  float64
	Last    float64
	Bid     float64
	Ask     float64
	PrevDay float64
}
type Ticker struct {
	Bid        float64 `json:"Bid"`
	Ask        float64 `json:"Ask"`
	Last       float64 `json:"Last"`
	ServerTime int64
}

// service interface
type MarketService interface {
	RefreshAllTicks(int) error
	RefreshTick(int, string, Ticker) error
	GetCurrentTick(int, string) (Ticker, error)
	GetCurrentTick2(int, string, string) (Ticker, error)
	//RefreshCurrencies() error
}
