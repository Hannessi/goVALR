package goVALRapi

func NewHTTPRequester() *HTTPRequester {

	return &HTTPRequester{
		urlManager: *NewUrlManager(),
	}
}

type HTTPRequester struct {
	urlManager urlManager
}

func (h *HTTPRequester) GetPublicOrderBook(request GetPublicOrderBookRequest) (*GetPublicOrderBookResponse, error) {
	// todo segregate requester response with interface response
	response := &GetPublicOrderBookResponse{}
	if err := HttpRequestWrapper(GET, h.urlManager.GetPublicOrderBook(request.CurrencyPair), nil, response, ""); err != nil {
		return nil, err
	}
	return response, nil
}

func (h *HTTPRequester) GetPublicOrderBookNonAggregate(request GetPublicOrderBookNonAggregateRequest) (*GetPublicOrderBookNonAggregateResponse, error) {
	response := &GetPublicOrderBookNonAggregateResponse{}
	if err := HttpRequestWrapper(GET, h.urlManager.GetPublicOrderBookNonAggregate(request.CurrencyPair), nil, response, ""); err != nil {
		return nil, err
	}
	return response, nil
}
