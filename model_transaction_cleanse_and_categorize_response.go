/*
 * MX API
 *
 * The MX Atrium API supports over 48,000 data connections to thousands of financial institutions. It provides secure access to your users' accounts and transactions with industry-leading cleansing, categorization, and classification.  Atrium is designed according to resource-oriented REST architecture and responds with JSON bodies and HTTP response codes.  Use Atrium's development environment, vestibule.mx.com, to quickly get up and running. The development environment limits are 100 users, 25 members per user, and access to the top 15 institutions. Contact MX to purchase production access.
 *
 * API version: 0.1
 */

package atrium

type TransactionCleanseAndCategorizeResponse struct {
	Amount               float32 `json:"amount,omitempty"`
	Category             string  `json:"category,omitempty"`
	Description          string  `json:"description,omitempty"`
	Identifier           string  `json:"identifier,omitempty"`
	Type_                string  `json:"type,omitempty"`
	IsBillPay            bool    `json:"is_bill_pay,omitempty"`
	IsDirectDeposit      bool    `json:"is_direct_deposit,omitempty"`
	IsExpense            bool    `json:"is_expense,omitempty"`
	IsFee                bool    `json:"is_fee,omitempty"`
	IsIncome             bool    `json:"is_income,omitempty"`
	IsInternational      bool    `json:"is_international,omitempty"`
	IsOverdraftFee       bool    `json:"is_overdraft_fee,omitempty"`
	IsPayrollAdvance     bool    `json:"is_payroll_advance,omitempty"`
	MerchantCategoryCode float32 `json:"merchant_category_code,omitempty"`
	MerchantGUID         string  `json:"merchant_guid,omitempty"`
	OriginalDescription  string  `json:"original_description,omitempty"`
}
