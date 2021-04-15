package valrAccountApi

import (
	goVALRapi "github.com/hannessi/goVALR/api"
	"time"
)

type Requester interface {
	SetCredentials(apiKey, apiSecret string)
	GetBalances(GetBalancesRequest) (*GetBalancesResponse, error)
	GetTransactionHistory(GetTransactionHistoryRequest) (*GetTransactionHistoryResponse, error)
	GetTradeHistory(GetTradeHistoryRequest) (*GetTradeHistoryResponse, error)
}

type GetBalancesRequest struct {
}

type GetBalancesResponse struct {
	AccountBalances []goVALRapi.AccountBalance
}

type GetTransactionHistoryRequest struct {
	Skip             int
	Limit            int
	TransactionTypes []goVALRapi.TransactionType
	Currency         string
	StartTime        time.Time
	EndTime          time.Time
	BeforeId         string
}

type GetTransactionHistoryResponse struct {
	Transactions []goVALRapi.Transaction
}

type GetTradeHistoryRequest struct {
	CurrencyPair string
	Limit        int
}

type GetTradeHistoryResponse struct {
	Trades []goVALRapi.Trade
}
