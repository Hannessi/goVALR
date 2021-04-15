package main

import (
	"fmt"
	"github.com/hannessi/goVALR"
)

func main() {
	client := goVALR.NewClient()

	client.SetCredentials(api_key, api_secret)

	getBalancesResponse, err := client.GetBalances(goVALR.GetBalancesRequest{})
	if err != nil {
		panic("Could not get balances: "+ err.Error())
	}

	fmt.Println("GetBalances")
	fmt.Println("Response: ", getBalancesResponse)

	//todo awaiting feedback from VALR w.r.t. empty response vs empty list
	getTransactionHistoryResponse, err := client.GetTransactionHistory(goVALR.GetTransactionHistoryRequest{
		Currency:         "BTCZAR",
	})
	if err != nil {
		panic("Could not get transaction history: "+err.Error())
	}

	fmt.Println("GetTransactionHistory")
	fmt.Println("Response: ", getTransactionHistoryResponse)

	getTradeHistoryResponse, err := client.GetTradeHistory(goVALR.GetTradeHistoryRequest{
		CurrencyPair: "BTCZAR",
		Limit:        0,
	})
	if err != nil {
		panic("Could not get trade history: "+err.Error())
	}

	fmt.Println("GetTradeHistory")
	fmt.Println("Response: ", getTradeHistoryResponse)




}