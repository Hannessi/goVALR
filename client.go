package goVALR

import (
	goVALRapi "github.com/hannessi/goVALR/api"
)

func NewClient() *Client {
	return &Client{
		requester: goVALRapi.NewHTTPRequester(),
	}
}

type Client struct {
	apiKey    string
	apiSecret string
	requester goVALRapi.Requester
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
	response, err := c.requester.GetPublicOrderBook(goVALRapi.GetPublicOrderBookRequest{
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
	response, err := c.requester.GetPublicOrderBookNonAggregate(goVALRapi.GetPublicOrderBookNonAggregateRequest{
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
	response, err := c.requester.GetCurrencies(goVALRapi.GetCurrenciesRequest{})
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
	response, err := c.requester.GetCurrencyPairs(goVALRapi.GetCurrencyPairsRequest{})
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
	response, err := c.requester.GetOrderTypes(goVALRapi.GetOrderTypesRequest{})
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
	response, err := c.requester.GetOrderTypesForCurrencyPair(goVALRapi.GetOrderTypesForCurrencyPairRequest{
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
	response, err := c.requester.GetMarketSummary(goVALRapi.GetMarketSummaryRequest{})
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
	response, err := c.requester.GetMarketSummaryForCurrencyPair(goVALRapi.GetMarketSummaryForCurrencyPairRequest{
		CurrencyPair: request.CurrencyPair,
	})
	if err != nil {
		return nil, err
	}
	return &GetMarketSummaryForCurrencyPairResponse{
		MarketSummary: response.MarketSummary,
	}, nil
}
