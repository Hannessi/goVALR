package valrPublicApi

import (
	goVALRapi "github.com/hannessi/goVALR/api"
	"strconv"
	"strings"
	"time"
)

func NewUrlManager() *urlManager {
	return &urlManager{
		baseUrl: goVALRapi.VALR_LIVE_API_BASE_URL+"public/",
	}
}

type urlManager struct {
	baseUrl string
}

func (u *urlManager) GetOrderBook(currencyPair string) string {
	// todo validate currency pair
	return u.baseUrl + currencyPair + "/orderbook"
}

func (u *urlManager) GetOrderBookNonAggregate(currencyPair string) string {
	// todo validate currency pair
	return u.baseUrl + currencyPair + "/orderbook/full"
}

func (u *urlManager) GetCurrencies() string {
	return u.baseUrl + "currencies"
}

func (u *urlManager) GetCurrencyPairs() string {
	return u.baseUrl + "pairs"
}

func (u *urlManager) GetOrderTypes() string {
	return u.baseUrl + "ordertypes"
}

func (u *urlManager) GetOrderTypesForCurrencyPair(currencyPair string) string {
	return u.baseUrl + currencyPair + "/ordertypes"
}

func (u *urlManager) GetMarketSummary() string {
	return u.baseUrl + "marketsummary"
}

func (u *urlManager) GetMarketSummaryForCurrencyPair(currencyPair string) string {
	return u.baseUrl + currencyPair + "/marketsummary"
}

func (u *urlManager) GetTradeHistory(currencyPair string, skip, limit int, startTime, endTime time.Time, beforeId string) string {
	url := u.baseUrl + currencyPair + "/trades"

	additionalParameters := make([]string, 0)
	if skip >= 0 {
		additionalParameters = append(additionalParameters, "skip="+strconv.Itoa(skip))
	}
	if limit > 0 {
		additionalParameters = append(additionalParameters, "limit="+strconv.Itoa(limit))
	}
	if !startTime.IsZero() {
		additionalParameters = append(additionalParameters, "startTime="+startTime.Format(time.RFC3339))
	}
	if !endTime.IsZero() {
		additionalParameters = append(additionalParameters, "endTime="+endTime.Format(time.RFC3339))
	}
	if beforeId != "" {
		additionalParameters = append(additionalParameters, "beforeId="+beforeId)
	}
	if len(additionalParameters) > 0 {
		url = url + "?" + strings.Join(additionalParameters, "&")
	}

	return url
}

func (u *urlManager) GetServerTime() string {
	return u.baseUrl + "time"
}

func (u *urlManager) GetValrStatus() string {
	return u.baseUrl + "status"
}
//
//func (u *urlManager) GetAccountBalances() string {
//	return u.baseUrl + "account/balances"
//}
//
//func (u *urlManager) GetAccountTransactionHistory(
//	skip, limit int,
//	transactionTypes []TransactionType,
//	currency string,
//	startTime, endTime time.Time,
//	beforeId string) string {
//
//	url := u.baseUrl + "account/transactionhistory"
//
//	additionalParameters := make([]string, 0)
//	if skip >= 0 {
//		additionalParameters = append(additionalParameters, "skip="+strconv.Itoa(skip))
//	}
//	if limit > 0 {
//		additionalParameters = append(additionalParameters, "limit="+strconv.Itoa(limit))
//	}
//	if len(transactionTypes) > 0 {
//		additionalParameters = append(additionalParameters, "transactionTypes="+joinTransactionTypes(transactionTypes,","))
//	}
//	if currency != "" {
//		additionalParameters = append(additionalParameters, "currency="+currency)
//	}
//	if !startTime.IsZero() {
//		additionalParameters = append(additionalParameters, "startTime="+startTime.Format(time.RFC3339))
//	}
//	if !endTime.IsZero() {
//		additionalParameters = append(additionalParameters, "endTime="+endTime.Format(time.RFC3339))
//	}
//	if beforeId != "" {
//		additionalParameters = append(additionalParameters, "beforeId="+beforeId)
//	}
//
//	if len(additionalParameters) > 0 {
//		url = url + "?" + strings.Join(additionalParameters, "&")
//	}
//
//	return url
//}