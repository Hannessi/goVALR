package goVALRapi

func NewUrlManager() *urlManager {
	return &urlManager{
		baseUrl: VALR_LIVE_API_BASE_URL,
	}
}

type urlManager struct {
	baseUrl string
}

func (u *urlManager) GetPublicOrderBook(currencyPair string) string {
	// todo validate currency pair
	return u.baseUrl + "public/" + currencyPair + "/orderbook"
}

func (u *urlManager) GetPublicOrderBookNonAggregate(currencyPair string) string {
	// todo validate currency pair
	return u.baseUrl + "public/" + currencyPair + "/orderbook/full"
}
