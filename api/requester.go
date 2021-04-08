package goVALRapi

type Requester interface {
	// public
	GetPublicOrderBook(GetPublicOrderBookRequest) (*GetPublicOrderBookResponse, error)
	GetPublicOrderBookNonAggregate(GetPublicOrderBookNonAggregateRequest) (*GetPublicOrderBookNonAggregateResponse, error)
	GetCurrencies(GetCurrenciesRequest) (*GetCurrenciesResponse, error)
	GetCurrencyPairs(GetCurrencyPairsRequest) (*GetCurrencyPairsResponse, error)
	GetOrderTypes(GetOrderTypesRequest) (*GetOrderTypesResponse, error)
	GetOrderTypesForCurrencyPair(GetOrderTypesForCurrencyPairRequest) (*GetOrderTypesForCurrencyPairResponse, error)
	GetMarketSummary(GetMarketSummaryRequest) (*GetMarketSummaryResponse, error)
	GetMarketSummaryForCurrencyPair(GetMarketSummaryForCurrencyPairRequest) (*GetMarketSummaryForCurrencyPairResponse, error)
}

type GetPublicOrderBookRequest struct {
	CurrencyPair string
}

type GetPublicOrderBookResponse struct {
	Asks []OrderBookEntry `json:"asks"`
	Bids []OrderBookEntry `json:"bids"`
}

type GetPublicOrderBookNonAggregateRequest struct {
	CurrencyPair string
}

type GetPublicOrderBookNonAggregateResponse struct {
	Asks       []NonAggregateOrderBookEntry `json:"asks"`
	Bids       []NonAggregateOrderBookEntry `json:"bids"`
	LastChange string                       `json:"LastChange"`
}

type GetCurrenciesRequest struct {
}

type GetCurrenciesResponse struct {
	Currencies []Currency `json:"currencies"`
}

type GetCurrencyPairsRequest struct {
}

type GetCurrencyPairsResponse struct {
	CurrencyPairs []CurrencyPair `json:"currencyPairs"`
}

type GetOrderTypesRequest struct{}

type GetOrderTypesResponse struct {
	OrderTypes []OrderTypePerCurrencyPair `json:"orderTypePerCurrencyPair"`
}

type GetOrderTypesForCurrencyPairRequest struct {
	CurrencyPair string
}

type GetOrderTypesForCurrencyPairResponse struct {
	OrderTypes []string
}

type GetMarketSummaryRequest struct {}

type GetMarketSummaryResponse struct {
	MarketSummaries []MarketSummary
}

type GetMarketSummaryForCurrencyPairRequest struct {
	CurrencyPair string
}

type GetMarketSummaryForCurrencyPairResponse struct {
	MarketSummary MarketSummary
}