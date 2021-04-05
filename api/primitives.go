package goVALRapi

type OrderBookEntry struct {
	Side         string `json:"side"`
	Quantity     string `json:"quantity"`
	Price        string `json:"price"`
	CurrencyPair string `json:"currencyPair"`
	OrderCount   int64  `json:"orderCount"`
}

type NonAggregateOrderBookEntry struct {
	Side            string `json:"side"`
	Quantity        string `json:"quantity"`
	Price           string `json:"price"`
	CurrencyPair    string `json:"currencyPair"`
	Id              string `json:"id"`
	PositionAtPrice int64  `json:"positionAtPrice"`
}
