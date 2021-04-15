package valrPublicApi

import goVALRapi "github.com/hannessi/goVALR/api"

func NewHTTPRequester() *HTTPRequester {

	return &HTTPRequester{
		urlManager: *NewUrlManager(),
	}
}

type HTTPRequester struct {
	urlManager urlManager
}

func (h *HTTPRequester) GetOrderBook(request GetOrderBookRequest) (*GetOrderBookResponse, error) {
	// todo segregate requester response with interface response
	response := &GetOrderBookResponse{}
	if err := goVALRapi.HttpRequestWrapper(goVALRapi.GET, h.urlManager.GetOrderBook(request.CurrencyPair), nil, response, "", ""); err != nil {
		return nil, err
	}
	return response, nil
}

func (h *HTTPRequester) GetOrderBookNonAggregate(request GetOrderBookNonAggregateRequest) (*GetOrderBookNonAggregateResponse, error) {
	response := &GetOrderBookNonAggregateResponse{}
	if err := goVALRapi.HttpRequestWrapper(goVALRapi.GET, h.urlManager.GetOrderBookNonAggregate(request.CurrencyPair), nil, response, "", ""); err != nil {
		return nil, err
	}
	return response, nil
}

func (h *HTTPRequester) GetCurrencies(request GetCurrenciesRequest) (*GetCurrenciesResponse, error) {
	currencies := &[]goVALRapi.Currency{}
	if err := goVALRapi.HttpRequestWrapper(goVALRapi.GET, h.urlManager.GetCurrencies(), nil, currencies, "", ""); err != nil {
		return nil, err
	}
	return &GetCurrenciesResponse{Currencies: *currencies}, nil
}

func (h *HTTPRequester) GetCurrencyPairs(request GetCurrencyPairsRequest) (*GetCurrencyPairsResponse, error) {
	currencyPairs := &[]goVALRapi.CurrencyPair{}
	if err := goVALRapi.HttpRequestWrapper(goVALRapi.GET, h.urlManager.GetCurrencyPairs(), nil, currencyPairs, "", ""); err != nil {
		return nil, err
	}
	return &GetCurrencyPairsResponse{CurrencyPairs: *currencyPairs}, nil
}

func (h *HTTPRequester) GetOrderTypes(request GetOrderTypesRequest) (*GetOrderTypesResponse, error) {
	orderTypes := &[]goVALRapi.OrderTypePerCurrencyPair{}
	if err := goVALRapi.HttpRequestWrapper(goVALRapi.GET, h.urlManager.GetOrderTypes(), nil, orderTypes, "", ""); err != nil {
		return nil, err
	}
	return &GetOrderTypesResponse{OrderTypes: *orderTypes}, nil
}

func (h *HTTPRequester) GetOrderTypesForCurrencyPair(request GetOrderTypesForCurrencyPairRequest) (*GetOrderTypesForCurrencyPairResponse, error) {
	orderTypes := &[]string{}
	if err := goVALRapi.HttpRequestWrapper(goVALRapi.GET, h.urlManager.GetOrderTypesForCurrencyPair(request.CurrencyPair), nil, orderTypes, "", ""); err != nil {
		return nil, err
	}
	return &GetOrderTypesForCurrencyPairResponse{OrderTypes: *orderTypes}, nil
}

func (h *HTTPRequester) GetMarketSummary(request GetMarketSummaryRequest) (*GetMarketSummaryResponse, error) {
	marketSummaries := &[]goVALRapi.MarketSummary{}
	if err := goVALRapi.HttpRequestWrapper(goVALRapi.GET, h.urlManager.GetMarketSummary(), nil, marketSummaries, "", ""); err != nil {
		return nil, err
	}
	return &GetMarketSummaryResponse{MarketSummaries: *marketSummaries}, nil
}

func (h *HTTPRequester) GetMarketSummaryForCurrencyPair(request GetMarketSummaryForCurrencyPairRequest) (*GetMarketSummaryForCurrencyPairResponse, error) {
	marketSummary := &goVALRapi.MarketSummary{}
	if err := goVALRapi.HttpRequestWrapper(goVALRapi.GET, h.urlManager.GetMarketSummaryForCurrencyPair(request.CurrencyPair), nil, marketSummary, "", ""); err != nil {
		return nil, err
	}
	return &GetMarketSummaryForCurrencyPairResponse{MarketSummary: *marketSummary}, nil
}

func (h *HTTPRequester) GetTradeHistory(request GetTradeHistoryRequest) (*GetTradeHistoryResponse, error) {
	trades := &[]goVALRapi.Trade{}
	if err := goVALRapi.HttpRequestWrapper(goVALRapi.GET, h.urlManager.GetTradeHistory(request.CurrencyPair, request.Skip, request.Limit, request.StartTime, request.EndTime, request.BeforeId), nil, trades, "", ""); err != nil {
		return nil, err
	}
	return &GetTradeHistoryResponse{Trades: *trades}, nil
}

func (h *HTTPRequester) GetServerTime(request GetServerTimeRequest) (*GetServerTimeResponse, error) {
	getServerTimeResponse := &GetServerTimeResponse{}
	if err := goVALRapi.HttpRequestWrapper(goVALRapi.GET, h.urlManager.GetServerTime(), nil, getServerTimeResponse, "", ""); err != nil {
		return nil, err
	}
	return getServerTimeResponse, nil
}

func (h *HTTPRequester) GetValrStatus(request GetValrStatusRequest) (*GetValrStatusResponse, error) {
	status := struct {
		Status string `json:"status"`
	}{}
	if err := goVALRapi.HttpRequestWrapper(goVALRapi.GET, h.urlManager.GetValrStatus(), nil, &status, "", ""); err != nil {
		return nil, err
	}
	return &GetValrStatusResponse{Status: status.Status}, nil
}
