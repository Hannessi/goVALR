package goVALRapi

type Requester interface {
	// public
	GetPublicOrderBook(GetPublicOrderBookRequest) (*GetPublicOrderBookResponse, error)
	GetPublicOrderBookNonAggregate(GetPublicOrderBookNonAggregateRequest) (*GetPublicOrderBookNonAggregateResponse, error)
	GetCurrencies(GetCurrenciesRequest) (*GetCurrenciesResponse, error)
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