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

type Currency struct {
	Symbol                  string `json:"symbol"`
	IsActive                bool   `json:"isActive"`
	ShortName               string `json:"shortName"`
	LongName                string `json:"longName"`
	DecimalPlaces           string `json:"decimalPlaces"`
	WithdrawalDecimalPlaces string `json:"withdrawalDecimalPlaces"`
}
