package main

import (
	"fmt"
	"github.com/hannessi/goVALR"
	"time"
)

const sleep_seconds = 7

func main() {
	client := goVALR.NewClient()

	//credentials not needed for public endpoints

	getPublicOrderBookResponse, err := client.GetPublicOrderBook(goVALR.GetPublicOrderBookRequest{
		CurrencyPair: "BTCZAR",
	})
	if err != nil {
		panic("Could not get public order book: " + err.Error())
	}

	fmt.Println("GetPublicOrderBook")
	fmt.Println("Response: ", getPublicOrderBookResponse)
	time.Sleep(sleep_seconds * time.Second)

	getPublicOrderBookNonAggregateResponse, err := client.GetPublicOrderBookNonAggregate(goVALR.GetPublicOrderBookNonAggregateRequest{
		CurrencyPair: "BTCZAR",
	})
	if err != nil {
		panic("Could not get no aggregate public order book: " + err.Error())
	}

	fmt.Println("GetPublicOrderBookNonAggregate")
	fmt.Println("Response: ", getPublicOrderBookNonAggregateResponse)
	time.Sleep(sleep_seconds * time.Second)

	getCurrenciesResponse, err := client.GetCurrencies(goVALR.GetCurrenciesRequest{})
	if err != nil{
		panic("Could not get currencies: "+err.Error())
	}

	fmt.Println("GetCurrencies")
	fmt.Println("Response: ", getCurrenciesResponse)
	time.Sleep(sleep_seconds * time.Second)

	getCurrencyPairsResponse, err := client.GetCurrencyPairs(goVALR.GetCurrencyPairsRequest{})
	if err != nil{
		panic("Could not get currency pairs: "+err.Error())
	}

	fmt.Println("GetCurrencyPairs")
	fmt.Println("Response: ", getCurrencyPairsResponse)
	time.Sleep(sleep_seconds * time.Second)

	getOrderTypesResponse, err := client.GetOrderTypes(goVALR.GetOrderTypesRequest{})
	if err != nil{
		panic("Could not get order types: "+err.Error())
	}

	fmt.Println("GetOrderTypes")
	fmt.Println("Response: ", getOrderTypesResponse)
	time.Sleep(sleep_seconds * time.Second)

	getOrderTypesForCurrencyPairResponse, err := client.GetOrderTypesForCurrencyPair(goVALR.GetOrderTypesForCurrencyPairRequest{
		CurrencyPair: "BTCZAR",
	})
	if err != nil{
		panic("Could not get order types: "+err.Error())
	}

	fmt.Println("GetOrderTypesForCurrencyPair")
	fmt.Println("Response: ", getOrderTypesForCurrencyPairResponse)
	time.Sleep(sleep_seconds * time.Second)

	getMarketSummaryResponse, err := client.GetMarketSummary(goVALR.GetMarketSummaryRequest{})
	if err != nil{
		panic("Could not get market summary: "+err.Error())
	}

	fmt.Println("GetMarketSummary")
	fmt.Println("Response: ", getMarketSummaryResponse)
	time.Sleep(sleep_seconds * time.Second)

	getMarketForCurrencyPairResponse, err := client.GetMarketSummaryForCurrencyPair(goVALR.GetMarketSummaryForCurrencyPairRequest{
		CurrencyPair: "BTCZAR",
	})
	if err != nil{
		panic("Could not get market summary for currency pair: "+err.Error())
	}

	fmt.Println("GetMarketForCurrencyPairResponse")
	fmt.Println("Response: ", getMarketForCurrencyPairResponse)
	time.Sleep(sleep_seconds * time.Second)

	getPublicTradeHistoryResponse, err := client.GetPublicTradeHistory(goVALR.GetPublicTradeHistoryRequest{
		CurrencyPair: "BTCZAR",
		Skip: 4,
		Limit: 2,
	})
	if err != nil{
		panic("Could not get public trade history: "+err.Error())
	}

	fmt.Println("GetPublicTradeHistoryResponse")
	fmt.Println("Response: ", getPublicTradeHistoryResponse)
	time.Sleep(sleep_seconds * time.Second)

}
