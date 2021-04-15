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

type CurrencyPair struct {
	Symbol            string `json:"symbol"`
	BaseCurrency      string `json:"baseCurrency"`
	QuoteCurrency     string `json:"quoteCurrency"`
	ShortName         string `json:"shortName"`
	Active            bool   `json:"active"`
	MinBaseAmount     string `json:"minBaseAmount"`
	MaxBaseAmount     string `json:"maxBaseAmount"`
	MinQuoteAmount    string `json:"minQuoteAmount"`
	MaxQuoteAmount    string `json:"maxQuoteAmount"`
	TickSize          string `json:"tickSize"`
	BaseDecimalPlaces string `json:"baseDecimalPlaces"`
}

type OrderTypePerCurrencyPair struct {
	CurrencyPair string   `json:"currencyPair"`
	OrderTypes   []string `json:"orderTypes"`
}

type MarketSummary struct {
	CurrencyPair       string `json:"currencyPair"`
	AskPrice           string `json:"askPrice"`
	BidPrice           string `json:"bidPrice"`
	LastTradedPrice    string `json:"lastTradedPrice"`
	PreviousClosePrice string `json:"previousClosePrice"`
	BaseVolume         string `json:"baseVolume"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Created            string `json:"created"`
	ChangeFromPrevious string `json:"changeFromPrevious"`
}

type Trade struct {
	Price        string `json:"price"`
	Quantity     string `json:"quantity"`
	CurrencyPair string `json:"currencyPair"`
	TradedAt     string `json:"tradedAt"`
	TakerSide    string `json:"takerSide"`
	SequenceId   int64  `json:"sequenceId"`
	Id           string `json:"id"`
	QuoteVolume  string `json:"quoteVolume"`
}

type AccountBalance struct {
	Currency  string `json:"currency"`
	Available string `json:"available"`
	Reserved  string `json:"reserved"`
	Total     string `json:"total"`
	UpdatedAt string `json:"updatedAt"`
}

type TransactionTypeWithDescription struct {
	Type        TransactionType `json:"type"`
	Description string          `json:"description"`
}