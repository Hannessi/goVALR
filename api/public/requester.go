package valrPublicApi

import (
	goVALRapi "github.com/hannessi/goVALR/api"
	"time"
)

type Requester interface {
	// public
	GetOrderBook(GetOrderBookRequest) (*GetOrderBookResponse, error)
	GetOrderBookNonAggregate(GetOrderBookNonAggregateRequest) (*GetOrderBookNonAggregateResponse, error)
	GetCurrencies(GetCurrenciesRequest) (*GetCurrenciesResponse, error)
	GetCurrencyPairs(GetCurrencyPairsRequest) (*GetCurrencyPairsResponse, error)
	GetOrderTypes(GetOrderTypesRequest) (*GetOrderTypesResponse, error)
	GetOrderTypesForCurrencyPair(GetOrderTypesForCurrencyPairRequest) (*GetOrderTypesForCurrencyPairResponse, error)
	GetMarketSummary(GetMarketSummaryRequest) (*GetMarketSummaryResponse, error)
	GetMarketSummaryForCurrencyPair(GetMarketSummaryForCurrencyPairRequest) (*GetMarketSummaryForCurrencyPairResponse, error)
	GetTradeHistory(GetTradeHistoryRequest) (*GetTradeHistoryResponse, error)
	GetServerTime(GetServerTimeRequest) (*GetServerTimeResponse, error)
	GetValrStatus(GetValrStatusRequest) (*GetValrStatusResponse, error)
}

type GetOrderBookRequest struct {
	CurrencyPair string
}

type GetOrderBookResponse struct {
	Asks []goVALRapi.OrderBookEntry `json:"asks"`
	Bids []goVALRapi.OrderBookEntry `json:"bids"`
}

type GetOrderBookNonAggregateRequest struct {
	CurrencyPair string
}

type GetOrderBookNonAggregateResponse struct {
	Asks       []goVALRapi.NonAggregateOrderBookEntry `json:"asks"`
	Bids       []goVALRapi.NonAggregateOrderBookEntry `json:"bids"`
	LastChange string                                 `json:"LastChange"`
}

type GetCurrenciesRequest struct {
}

type GetCurrenciesResponse struct {
	Currencies []goVALRapi.Currency `json:"currencies"`
}

type GetCurrencyPairsRequest struct {
}

type GetCurrencyPairsResponse struct {
	CurrencyPairs []goVALRapi.CurrencyPair `json:"currencyPairs"`
}

type GetOrderTypesRequest struct{}

type GetOrderTypesResponse struct {
	OrderTypes []goVALRapi.OrderTypePerCurrencyPair `json:"orderTypePerCurrencyPair"`
}

type GetOrderTypesForCurrencyPairRequest struct {
	CurrencyPair string
}

type GetOrderTypesForCurrencyPairResponse struct {
	OrderTypes []string
}

type GetMarketSummaryRequest struct{}

type GetMarketSummaryResponse struct {
	MarketSummaries []goVALRapi.MarketSummary
}

type GetMarketSummaryForCurrencyPairRequest struct {
	CurrencyPair string
}

type GetMarketSummaryForCurrencyPairResponse struct {
	MarketSummary goVALRapi.MarketSummary
}

type GetTradeHistoryRequest struct {
	CurrencyPair string
	Skip         int
	Limit        int
	StartTime    time.Time
	EndTime      time.Time
	BeforeId     string
}

type GetTradeHistoryResponse struct {
	Trades []goVALRapi.Trade
}

type GetServerTimeRequest struct{}

type GetServerTimeResponse struct {
	EpochTime int64  `json:"epochTime"`
	Time      string `json:"time"`
}

type GetValrStatusRequest struct {

}

type GetValrStatusResponse struct {
	Status string
}