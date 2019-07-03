/*
 * MX API
 *
 * The MX Atrium API supports over 48,000 data connections to thousands of financial institutions. It provides secure access to your users' accounts and transactions with industry-leading cleansing, categorization, and classification.  Atrium is designed according to resource-oriented REST architecture and responds with JSON bodies and HTTP response codes.  Use Atrium's development environment, vestibule.mx.com, to quickly get up and running. The development environment limits are 100 users, 25 members per user, and access to the top 15 institutions. Contact MX to purchase production access.
 *
 * API version: 0.1
 */

package atrium

type Transaction struct {
	AccountGUID          string  `json:"account_guid,omitempty"`
	Amount               float32 `json:"amount,omitempty"`
	Category             string  `json:"category,omitempty"`
	CheckNumber          int32   `json:"check_number,omitempty"`
	CheckNumberString    string  `json:"check_number_string,omitempty"`
	CreatedAt            string  `json:"created_at,omitempty"`
	CurrencyCode         string  `json:"currency_code,omitempty"`
	Date                 string  `json:"date,omitempty"`
	Description          string  `json:"description,omitempty"`
	GUID                 string  `json:"guid,omitempty"`
	IsBillPay            bool    `json:"is_bill_pay,omitempty"`
	IsDirectDeposit      bool    `json:"is_direct_deposit,omitempty"`
	IsExpense            bool    `json:"is_expense,omitempty"`
	IsFee                bool    `json:"is_fee,omitempty"`
	IsIncome             bool    `json:"is_income,omitempty"`
	IsInternational      bool    `json:"is_international,omitempty"`
	IsOverdraftFee       bool    `json:"is_overdraft_fee,omitempty"`
	IsPayrollAdvance     bool    `json:"is_payroll_advance,omitempty"`
	Latitude             float32 `json:"latitude,omitempty"`
	Longitude            float32 `json:"longitude,omitempty"`
	MemberGUID           string  `json:"member_guid,omitempty"`
	Memo                 string  `json:"memo,omitempty"`
	MerchantCategoryCode int32   `json:"merchant_category_code,omitempty"`
	MerchantGUID         string  `json:"merchant_guid,omitempty"`
	OriginalDescription  string  `json:"original_description,omitempty"`
	PostedAt             string  `json:"posted_at,omitempty"`
	Status               string  `json:"status,omitempty"`
	TopLevelCategory     string  `json:"top_level_category,omitempty"`
	TransactedAt         string  `json:"transacted_at,omitempty"`
	Type_                string  `json:"type,omitempty"`
	UpdatedAt            string  `json:"updated_at,omitempty"`
	UserGUID             string  `json:"user_guid,omitempty"`
}
