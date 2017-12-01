package models

type Transaction struct {
	AccountGuid          string  `json:"account_guid,omitempty"`
	Amount               float64 `json:"amount,omitempty"`
	Category             string  `json:"category,omitempty"`
	CheckNumber          int64   `json:"check_number,omitempty"`
	CreatedAt            string  `json:"created_at,omitempty"`
	Description          string  `json:"description,omitempty"`
	Guid                 string  `json:"guid,omitempty"`
	IsBillPay            bool    `json:"is_bill_pay,omitempty"`
	IsDirectDeposit      bool    `json:"is_direct_deposit,omitempty"`
	IsExpense            bool    `json:"is_expense,omitempty"`
	IsFee                bool    `json:"is_fee,omitempty"`
	IsIncome             bool    `json:"is_income,omitempty"`
	IsOverdraftFee       bool    `json:"is_overdraft_fee,omitempty"`
	IsPayrollAdvance     bool    `json:"is_payroll_advance,omitempty"`
	Latitude             float64 `json:"latitude,omitempty"`
	Longitude            float64 `json:"longitude,omitempty"`
	MemberGuid           string  `json:"member_guid,omitempty"`
	Memo                 string  `json:"memo,omitempty"`
	MerchantCategoryCode int64   `json:"merchant_category_code,omitempty"`
	OriginalDescription  string  `json:"original_description,omitempty"`
	PostedAt             string  `json:"posted_at,omitempty"`
	Status               string  `json:"status,omitempty"`
	TopLevelCategory     string  `json:"top_level_category,omitempty"`
	TransactedAt         string  `json:"transacted_at,omitempty"`
	Type                 string  `json:"type,omitempty"`
	UpdatedAt            string  `json:"updated_at,omitempty"`
	UserGuid             string  `json:"user_guid,omitempty"`
}

type TransactionResponse struct {
	Transaction *Transaction `json:"transaction"`
}

type TransactionsResponse struct {
	Transactions []*Transaction `json:"transactions"`
}
