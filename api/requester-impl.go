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

func (h *HTTPRequester) GetCurrencies(request GetCurrenciesRequest) (*GetCurrenciesResponse, error) {
	currencies := &[]Currency{}
	if err := HttpRequestWrapper(GET, h.urlManager.GetCurrencies(), nil, currencies, ""); err != nil {
		return nil, err
	}
	return &GetCurrenciesResponse{Currencies: *currencies}, nil
}

func (h *HTTPRequester) GetCurrencyPairs(request GetCurrencyPairsRequest) (*GetCurrencyPairsResponse, error) {
	currencyPairs := &[]CurrencyPair{}
	if err := HttpRequestWrapper(GET, h.urlManager.GetCurrencyPairs(), nil, currencyPairs, ""); err != nil {
		return nil, err
	}
	return &GetCurrencyPairsResponse{CurrencyPairs: *currencyPairs}, nil
}

func (h *HTTPRequester) GetOrderTypes(request GetOrderTypesRequest) (*GetOrderTypesResponse, error) {
	orderTypes := &[]OrderTypePerCurrencyPair{}
	if err := HttpRequestWrapper(GET, h.urlManager.GetOrderTypes(), nil, orderTypes, ""); err != nil {
		return nil, err
	}
	return &GetOrderTypesResponse{OrderTypes: *orderTypes}, nil
}

func (h *HTTPRequester) GetOrderTypesForCurrencyPair(request GetOrderTypesForCurrencyPairRequest) (*GetOrderTypesForCurrencyPairResponse, error) {
	orderTypes := &[]string{}
	if err := HttpRequestWrapper(GET, h.urlManager.GetOrderTypesForCurrencyPair(request.CurrencyPair), nil, orderTypes, ""); err != nil {
		return nil, err
	}
	return &GetOrderTypesForCurrencyPairResponse{OrderTypes: *orderTypes}, nil
}

func (h *HTTPRequester) GetMarketSummary(request GetMarketSummaryRequest) (*GetMarketSummaryResponse, error) {
	marketSummaries := &[]MarketSummary{}
	if err := HttpRequestWrapper(GET, h.urlManager.GetMarketSummary(), nil, marketSummaries, ""); err != nil {
		return nil, err
	}
	return &GetMarketSummaryResponse{MarketSummaries: *marketSummaries}, nil
}

func (h *HTTPRequester) GetMarketSummaryForCurrencyPair(request GetMarketSummaryForCurrencyPairRequest) (*GetMarketSummaryForCurrencyPairResponse, error) {
	marketSummary := &MarketSummary{}
	if err := HttpRequestWrapper(GET, h.urlManager.GetMarketSummaryForCurrencyPair(request.CurrencyPair), nil, marketSummary, ""); err != nil {
		return nil, err
	}
	return &GetMarketSummaryForCurrencyPairResponse{MarketSummary: *marketSummary}, nil
}

//func (h *HTTPRequester) GetOrderTypesForCurrencyPair(request GetOrderTypesForCurrencyPairRequest) (*GetOrderTypesForCurrencyPairResponse, error) {
//	orderTypes := &[]string{}
//	if err := HttpRequestWrapper(GET, h.urlManager.GetOrderTypesForCurrencyPair(request.CurrencyPair), nil, orderTypes, ""); err != nil {
//		return nil, err
//	}
//	return &GetOrderTypesForCurrencyPairResponse{OrderTypes: *orderTypes}, nil
//}
//
//func (h *HTTPRequester) GetOrderTypesForCurrencyPair(request GetOrderTypesForCurrencyPairRequest) (*GetOrderTypesForCurrencyPairResponse, error) {
//	orderTypes := &[]string{}
//	if err := HttpRequestWrapper(GET, h.urlManager.GetOrderTypesForCurrencyPair(request.CurrencyPair), nil, orderTypes, ""); err != nil {
//		return nil, err
//	}
//	return &GetOrderTypesForCurrencyPairResponse{OrderTypes: *orderTypes}, nil
//}
//
//func (h *HTTPRequester) GetOrderTypesForCurrencyPair(request GetOrderTypesForCurrencyPairRequest) (*GetOrderTypesForCurrencyPairResponse, error) {
//	orderTypes := &[]string{}
//	if err := HttpRequestWrapper(GET, h.urlManager.GetOrderTypesForCurrencyPair(request.CurrencyPair), nil, orderTypes, ""); err != nil {
//		return nil, err
//	}
//	return &GetOrderTypesForCurrencyPairResponse{OrderTypes: *orderTypes}, nil
//}
//
//
