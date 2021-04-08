package goVALRapi

type TransactionType string

func (t TransactionType) String() string {
	return string(t)
}

const LIMIT_BUY TransactionType = "LIMIT_BUY"
const LIMIT_SELL TransactionType = "LIMIT_SELL"
const MARKET_BUY TransactionType = "MARKET_BUY"
const MARKET_SELL TransactionType = "MARKET_SELL"
const SIMPLE_BUY TransactionType = "SIMPLE_BUY"
const SIMPLE_SELL TransactionType = "SIMPLE_SELL"
const MAKER_REWARD TransactionType = "MAKER_REWARD"
const BLOCKCHAIN_RECEIVE TransactionType = "BLOCKCHAIN_RECEIVE"
const BLOCKCHAIN_SEND TransactionType = "BLOCKCHAIN_SEND"
const FIAT_DEPOSIT TransactionType = "FIAT_DEPOSIT"
const FIAT_WITHDRAWAL TransactionType = "FIAT_WITHDRAWAL"
const REFERRAL_REBATE TransactionType = "REFERRAL_REBATE"
const REFERRAL_REWARD TransactionType = "REFERRAL_REWARD"
const PROMOTIONAL_REBATE TransactionType = "PROMOTIONAL_REBATE"
const INTERNAL_TRANSFER TransactionType = "INTERNAL_TRANSFER"
const FIAT_WITHDRAWAL_REVERSAL TransactionType = "FIAT_WITHDRAWAL_REVERSAL"
