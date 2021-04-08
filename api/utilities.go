package goVALRapi

import "strings"

func joinTransactionTypes(transactionTypes []TransactionType, separator string) string {
	convertedList := make([]string,0)
	for _, value := range transactionTypes {
		convertedList = append(convertedList, value.String())
	}
	return strings.Join(convertedList, separator)
}