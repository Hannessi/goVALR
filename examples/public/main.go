package main

import (
	"fmt"
	"github.com/hannessi/goVALR"
)

func main() {
	client := goVALR.NewClient()

	// credentials not needed for public endpoints

	getPublicOrderBookResponse, err := client.GetPublicOrderBook(goVALR.GetPublicOrderBookRequest{
		CurrencyPair: "BTCZAR",
	})
	if err != nil {
		panic("Could not get public order book: " + err.Error())
	}

	fmt.Println("GetPublicOrderBook")
	fmt.Println("Response: ", getPublicOrderBookResponse)

	getPublicOrderBookNonAggregateResponse, err := client.GetPublicOrderBookNonAggregate(goVALR.GetPublicOrderBookNonAggregateRequest{
		CurrencyPair: "BTCZAR",
	})
	if err != nil {
		panic("Could not get no aggregate public order book: " + err.Error())
	}

	fmt.Println("GetPublicOrderBookNonAggregate")
	fmt.Println("Response: ", getPublicOrderBookNonAggregateResponse)

	getCurrenciesResponse, err := client.GetCurrencies(goVALR.GetCurrenciesRequest{})
	if err != nil{
		panic("Could not get currencies: "+err.Error())
	}

	fmt.Println("GetCurrencies")
	fmt.Println("Response: ", getCurrenciesResponse)

}
