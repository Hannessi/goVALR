package valrAccountApi

import (
	goVALRapi "github.com/hannessi/goVALR/api"
	"strconv"
	"strings"
	"time"
)

func NewUrlManager() *urlManager {
	return &urlManager{
		baseUrl: goVALRapi.VALR_LIVE_API_BASE_URL + "account/",
	}
}

type urlManager struct {
	baseUrl string
}

func (u *urlManager) GetBalances() string {
	return u.baseUrl + "/balances"
}

func (u *urlManager) GetTransactionHistory(
	skip, limit int,
	transactionTypes []goVALRapi.TransactionType,
	currency string,
	startTime, endTime time.Time,
	beforeId string) string {

	url := u.baseUrl + "account/transactionhistory"

	additionalParameters := make([]string, 0)
	if skip >= 0 {
		additionalParameters = append(additionalParameters, "skip="+strconv.Itoa(skip))
	}
	if limit > 0 {
		additionalParameters = append(additionalParameters, "limit="+strconv.Itoa(limit))
	}
	if len(transactionTypes) > 0 {
		additionalParameters = append(additionalParameters, "transactionTypes="+goVALRapi.JoinTransactionTypes(transactionTypes, ","))
	}
	if currency != "" {
		additionalParameters = append(additionalParameters, "currency="+currency)
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

func (u *urlManager) GetTradeHistory(currencyPair string, limit int) string {
	url := u.baseUrl + currencyPair + "/tradeHistory"

	additionalParameters := make([]string, 0)
	if limit > 0 {
		additionalParameters = append(additionalParameters, "limit="+strconv.Itoa(limit))
	}

	if len(additionalParameters) > 0 {
		url = url + "?" + strings.Join(additionalParameters, "&")
	}

	return url
}
