package makerReward

import (
	"encoding/json"
	"errors"
	goVALRapi "github.com/hannessi/goVALR/api"
)

type Transaction struct {
	TransactionType goVALRapi.TransactionTypeWithDescription `json:"transactionType"`
	CreditCurrency  string                                   `json:"creditCurrency"`
	CreditValue     string                                   `json:"creditValue"`
	EventAt         string                                   `json:"eventAt"`
	Id              string                                   `json:"id"`
}

func (m *Transaction) GetType() goVALRapi.TransactionType {
	return goVALRapi.MAKER_REWARD
}

func Unmarshal(rawJson []byte) (goVALRapi.Transaction, error) {

	makerReward := Transaction{}

	err := json.Unmarshal(rawJson, &makerReward)
	if err != nil {
		return nil, errors.New("could not unmarshal the raw JSON: " + err.Error())
	}
	return goVALRapi.Transaction(&makerReward), nil
}
