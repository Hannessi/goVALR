package limitBuy

import (
	"encoding/json"
	"errors"
	goVALRapi "github.com/hannessi/goVALR/api"
)

type Transaction struct {
	TransactionType goVALRapi.TransactionTypeWithDescription `json:"transactionType"`
	DebitCurrency   string                                   `json:"debitCurrency"`
	DebitValue      string                                   `json:"debitValue"`
	CreditCurrency  string                                   `json:"creditCurrency"`
	CreditValue     string                                   `json:"creditValue"`
	EventAt         string                                   `json:"eventAt"`
	Id              string                                   `json:"id"`
	AdditionalInfo  struct {
		CostPerCoin        int64  `json:"costPerCoin"`
		CostPerCoinSymbol  string `json:"costPerCoinSymbol"`
		CurrencyPairSymbol string `json:"currencyPairSymbol"`
		OrderId            string `json:"orderId"`
	} `json:"additionalInfo"`
}

func (m *Transaction) GetType() goVALRapi.TransactionType {
	return goVALRapi.LIMIT_BUY
}

func Unmarshal(rawJson []byte) (goVALRapi.Transaction, error) {

	limitBuy := Transaction{}

	err := json.Unmarshal(rawJson, &limitBuy)
	if err != nil {
		return nil, errors.New("could not unmarshal the raw JSON: " + err.Error())
	}
	return goVALRapi.Transaction(&limitBuy), nil
}
