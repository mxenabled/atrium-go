package models

type Transaction struct {
	TransactionTypeName string `json:"transaction_type_name,omitempty"`
	AccountGuid string `json:"account_guid,omitempty"`
	Amount float64 `json:"amount,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	Description string `json:"description,omitempty"`
	Guid string `json:"guid,omitempty"`
	HasBeenSplit bool `json:"has_been_split,omitempty"`
	HasBeenViewed bool `json:"has_been_viewed,omitempty"`
	IsBillPay bool `json:"is_bill_pay,omitempty"`
	IsDirectDeposit bool `json:"is_direct_deposit,omitempty"`
	IsExpense bool `json:"is_expense,omitempty"`
	IsFee bool `json:"is_fee,omitempty"`
	IsHidden bool `json:"is_hidden,omitempty"`
	IsIncome bool `json:"is_income,omitempty"`
	IsOverdraftFee bool `json:"is_overdraft_fee,omitempty"`
	IsPayrollAdvance bool `json:"is_payroll_advance,omitempty"`
	Latitude float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Memo string `json:"memo,omitempty"`
	MerchantCategoryCode int64 `json:"merchant_category_code,omitempty"`
	ParentGuid string `json:"parent_guid,omitempty"`
	PostedAt string `json:"posted_at,omitempty"`
	TransactionType int64 `json:"transaction_type,omitempty"`
	TransactedAt string `json:"transacted_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	UserGuid string `json:"user_guid,omitempty"`
	CheckNumber int64 `json:"check_number,omitempty"`
	FeedCheckNumber int64 `json:"feed_check_number,omitempty"`
}

type TransactionResponse struct {
	Transaction *Transaction `json"transaction"`
}

type TransactionsResponse struct {
	Transactions []*Transaction `json"transactions"`
}
