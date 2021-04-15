package valrAccountApi

import (
	"encoding/json"
	"errors"
	"fmt"
	goVALRapi "github.com/hannessi/goVALR/api"
	"github.com/hannessi/goVALR/api/transactionTypes/limitBuy"
	"github.com/hannessi/goVALR/api/transactionTypes/makerReward"
)

func NewHTTPRequester() *HTTPRequester {

	return &HTTPRequester{
		urlManager: *NewUrlManager(),
	}
}

type HTTPRequester struct {
	apiKey     string
	apiSecret  string
	urlManager urlManager
}

func (h *HTTPRequester) SetCredentials(apiKey, apiSecret string) {
	h.apiKey = apiKey
	h.apiSecret = apiSecret
}

func (h *HTTPRequester) GetBalances(request GetBalancesRequest) (*GetBalancesResponse, error) {
	response := &[]goVALRapi.AccountBalance{}
	if err := goVALRapi.HttpRequestWrapper(goVALRapi.GET, h.urlManager.GetBalances(), nil, response, h.apiKey, h.apiSecret); err != nil {
		return nil, err
	}
	return &GetBalancesResponse{
		AccountBalances: *response,
	}, nil
}

func (h *HTTPRequester) GetTransactionHistory(request GetTransactionHistoryRequest) (*GetTransactionHistoryResponse, error) {
	rawList := &[][]byte{}

	if err := goVALRapi.HttpRequestWrapper(goVALRapi.GET, h.urlManager.GetTransactionHistory(
		request.Skip,
		request.Limit,
		request.TransactionTypes,
		request.Currency,
		request.StartTime,
		request.EndTime,
		request.BeforeId,
	), nil, rawList, h.apiKey, h.apiSecret); err != nil {
		return nil, err
	}

	// new transaction list
	transactionList := make([]goVALRapi.Transaction, len(*rawList))

	for i, rawItem := range *rawList {
		transactionTypeAware := goVALRapi.TransactionTypeWithDescription{}
		err := json.Unmarshal(rawItem, &transactionTypeAware)
		if err != nil {
			fmt.Println("could not unmarshal: ", err.Error())
			return nil, err
		}

		t := goVALRapi.Transaction(nil)

		switch transactionTypeAware.Type {
		case goVALRapi.MAKER_REWARD:
			t, err = makerReward.Unmarshal(rawItem)
		case goVALRapi.LIMIT_BUY:
			t, err = limitBuy.Unmarshal(rawItem)
		default:
			return nil, errors.New("Could not unmarshal response: Transaction type not implemented: " + transactionTypeAware.Type.String())
		}
		if err != nil {
			return nil, err
		}
		transactionList[i] = t
	}

	return &GetTransactionHistoryResponse{
		Transactions: transactionList,
	}, nil
}

func (h *HTTPRequester) GetTradeHistory(request GetTradeHistoryRequest) (*GetTradeHistoryResponse, error) {
	response := &[]goVALRapi.Trade{}
	if err := goVALRapi.HttpRequestWrapper(goVALRapi.GET, h.urlManager.GetTradeHistory(request.CurrencyPair, request.Limit), nil, response, h.apiKey,  h.apiSecret); err != nil {
		return nil, err
	}
	return &GetTradeHistoryResponse{
		Trades: *response,
	}, nil
}
