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
