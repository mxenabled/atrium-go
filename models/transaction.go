package models

type Transaction struct {

}

type TransactionResponse struct {
	Transaction *Transaction `json"transaction"`
}

type TransactionsResponse struct {
	Transactions []*Transaction `json"transactions"`
}
