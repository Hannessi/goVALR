package goVALR

import (
	goVALRapi "github.com/hannessi/goVALR/api"
	valrPublicApi "github.com/hannessi/goVALR/api/public"
	"time"
)

func NewClient() *Client {
	return &Client{
		publicApiRequester: valrPublicApi.NewHTTPRequester(),
	}
}

type Client struct {
	apiKey             string
	apiSecret          string
	publicApiRequester valrPublicApi.Requester
}

// Setting authentication
func (c *Client) SetCredentials(ApiKey, ApiSecret string) {
	c.apiKey = ApiKey
	c.apiSecret = ApiSecret
}

// Public API
type GetPublicOrderBookRequest struct {
	CurrencyPair string
}

type GetPublicOrderBookResponse struct {
	Asks []goVALRapi.OrderBookEntry
	Bids []goVALRapi.OrderBookEntry
}

func (c *Client) GetPublicOrderBook(request GetPublicOrderBookRequest) (*GetPublicOrderBookResponse, error) {
	response, err := c.publicApiRequester.GetOrderBook(valrPublicApi.GetOrderBookRequest{
		CurrencyPair: request.CurrencyPair,
	})
	if err != nil {
		return nil, err
	}
	return &GetPublicOrderBookResponse{
		Asks: response.Asks,
		Bids: response.Bids,
	}, nil
}

type GetPublicOrderBookNonAggregateRequest struct {
	CurrencyPair string
}

type GetPublicOrderBookNonAggregateResponse struct {
	Asks       []goVALRapi.NonAggregateOrderBookEntry `json:"asks"`
	Bids       []goVALRapi.NonAggregateOrderBookEntry `json:"bids"`
	LastChange string                                 `json:"LastChange"`
}

func (c *Client) GetPublicOrderBookNonAggregate(request GetPublicOrderBookNonAggregateRequest) (*GetPublicOrderBookNonAggregateResponse, error) {
	response, err := c.publicApiRequester.GetOrderBookNonAggregate(valrPublicApi.GetOrderBookNonAggregateRequest{
		CurrencyPair: request.CurrencyPair,
	})
	if err != nil {
		return nil, err
	}
	return &GetPublicOrderBookNonAggregateResponse{
		Asks:       response.Asks,
		Bids:       response.Bids,
		LastChange: response.LastChange,
	}, nil

}

type GetCurrenciesRequest struct{}

type GetCurrenciesResponse struct {
	Currencies []goVALRapi.Currency
}

func (c *Client) GetCurrencies(request GetCurrenciesRequest) (*GetCurrenciesResponse, error) {
	response, err := c.publicApiRequester.GetCurrencies(valrPublicApi.GetCurrenciesRequest{})
	if err != nil {
		return nil, err
	}
	return &GetCurrenciesResponse{
		Currencies: response.Currencies,
	}, nil
}

type GetCurrencyPairsRequest struct{}

type GetCurrencyPairsResponse struct {
	CurrencyPairs []goVALRapi.CurrencyPair
}

func (c *Client) GetCurrencyPairs(request GetCurrencyPairsRequest) (*GetCurrencyPairsResponse, error) {
	response, err := c.publicApiRequester.GetCurrencyPairs(valrPublicApi.GetCurrencyPairsRequest{})
	if err != nil {
		return nil, err
	}
	return &GetCurrencyPairsResponse{
		CurrencyPairs: response.CurrencyPairs,
	}, nil
}

type GetOrderTypesRequest struct{}

type GetOrderTypesResponse struct {
	OrderTypes []goVALRapi.OrderTypePerCurrencyPair
}

func (c *Client) GetOrderTypes(request GetOrderTypesRequest) (*GetOrderTypesResponse, error) {
	response, err := c.publicApiRequester.GetOrderTypes(valrPublicApi.GetOrderTypesRequest{})
	if err != nil {
		return nil, err
	}
	return &GetOrderTypesResponse{
		OrderTypes: response.OrderTypes,
	}, nil
}

type GetOrderTypesForCurrencyPairRequest struct {
	CurrencyPair string
}

type GetOrderTypesForCurrencyPairResponse struct {
	OrderTypes []string
}

func (c *Client) GetOrderTypesForCurrencyPair(request GetOrderTypesForCurrencyPairRequest) (*GetOrderTypesForCurrencyPairResponse, error) {
	response, err := c.publicApiRequester.GetOrderTypesForCurrencyPair(valrPublicApi.GetOrderTypesForCurrencyPairRequest{
		CurrencyPair: request.CurrencyPair,
	})
	if err != nil {
		return nil, err
	}
	return &GetOrderTypesForCurrencyPairResponse{
		OrderTypes: response.OrderTypes,
	}, nil
}

type GetMarketSummaryRequest struct{}

type GetMarketSummaryResponse struct {
	// todo segregate goVALR types from VALR api types
	MarketSummaries []goVALRapi.MarketSummary
}

func (c *Client) GetMarketSummary(request GetMarketSummaryRequest) (*GetMarketSummaryResponse, error) {
	response, err := c.publicApiRequester.GetMarketSummary(valrPublicApi.GetMarketSummaryRequest{})
	if err != nil {
		return nil, err
	}
	return &GetMarketSummaryResponse{
		MarketSummaries: response.MarketSummaries,
	}, nil
}

type GetMarketSummaryForCurrencyPairRequest struct {
	CurrencyPair string
}

type GetMarketSummaryForCurrencyPairResponse struct {
	MarketSummary goVALRapi.MarketSummary
}

func (c *Client) GetMarketSummaryForCurrencyPair(request GetMarketSummaryForCurrencyPairRequest) (*GetMarketSummaryForCurrencyPairResponse, error) {
	response, err := c.publicApiRequester.GetMarketSummaryForCurrencyPair(valrPublicApi.GetMarketSummaryForCurrencyPairRequest{
		CurrencyPair: request.CurrencyPair,
	})
	if err != nil {
		return nil, err
	}
	return &GetMarketSummaryForCurrencyPairResponse{
		MarketSummary: response.MarketSummary,
	}, nil
}

type GetPublicTradeHistoryRequest struct {
	CurrencyPair string
	Skip         int
	Limit        int
	StartTime    time.Time
	EndTime      time.Time
	BeforeId     string
}

type GetPublicTradeHistoryResponse struct {
	Trades []goVALRapi.Trade
}

func (c *Client) GetPublicTradeHistory(request GetPublicTradeHistoryRequest) (*GetPublicTradeHistoryResponse, error) {
	response, err := c.publicApiRequester.GetTradeHistory(valrPublicApi.GetTradeHistoryRequest{
		CurrencyPair: request.CurrencyPair,
		Skip:         request.Skip,
		Limit:        request.Limit,
		StartTime:    request.StartTime,
		EndTime:      request.EndTime,
		BeforeId:     request.BeforeId,
	})
	if err != nil {
		return nil, err
	}
	return &GetPublicTradeHistoryResponse{
		Trades: response.Trades,
	}, nil
}

type GetServerTimeRequest struct{}

type GetServerTimeResponse struct {
	EpochTime int64
	Time      string // todo make time.Time{}
}

func (c *Client) GetServerTime(request GetServerTimeRequest) (*GetServerTimeResponse, error) {
	response, err := c.publicApiRequester.GetServerTime(valrPublicApi.GetServerTimeRequest{})
	if err != nil {
		return nil, err
	}
	return &GetServerTimeResponse{
		EpochTime: response.EpochTime,
		Time:      response.Time,
	}, nil
}

type GetValrStatusRequest struct{}

type GetValrStatusResponse struct {
	Status string
}

func (c *Client) GetValrStatus(request GetValrStatusRequest) (*GetValrStatusResponse, error) {
	response, err := c.publicApiRequester.GetValrStatus(valrPublicApi.GetValrStatusRequest{})
	if err != nil {
		return nil, err
	}
	return &GetValrStatusResponse{
		Status: response.Status,
	}, nil
}
