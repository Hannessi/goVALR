package goVALRapi

type Transaction interface {
	GetType() TransactionType
}
