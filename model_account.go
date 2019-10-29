/*
 * MX API
 *
 * The MX Atrium API supports over 48,000 data connections to thousands of financial institutions. It provides secure access to your users' accounts and transactions with industry-leading cleansing, categorization, and classification.  Atrium is designed according to resource-oriented REST architecture and responds with JSON bodies and HTTP response codes.  Use Atrium's development environment, vestibule.mx.com, to quickly get up and running. The development environment limits are 100 users, 25 members per user, and access to the top 15 institutions. Contact MX to purchase production access.
 *
 * API version: 0.1
 */

package atrium

type Account struct {
	AccountNumber      string  `json:"account_number,omitempty"`
	Apr                float32 `json:"apr,omitempty"`
	Apy                float32 `json:"apy,omitempty"`
	AvailableBalance   float32 `json:"available_balance,omitempty"`
	AvailableCredit    float32 `json:"available_credit,omitempty"`
	Balance            float32 `json:"balance,omitempty"`
	CashBalance        float32 `json:"cash_balance,omitempty"`
	CashSurrenderValue float32 `json:"cash_surrender_value,omitempty"`
	CreatedAt          string  `json:"created_at,omitempty"`
	CreditLimit        float32 `json:"credit_limit,omitempty"`
	CurrencyCode       string  `json:"currency_code,omitempty"`
	DayPaymentIsDue    int32   `json:"day_payment_is_due,omitempty"`
	DeathBenefit       float32 `json:"death_benefit,omitempty"`
	GUID               string  `json:"guid,omitempty"`
	HoldingsValue      float32 `json:"holdings_value,omitempty"`
	InstitutionCode    string  `json:"institution_code,omitempty"`
	InterestRate       float32 `json:"interest_rate,omitempty"`
	IsClosed           bool    `json:"is_closed,omitempty"`
	LastPayment        float32 `json:"last_payment,omitempty"`
	LoanAmount         float32 `json:"loan_amount,omitempty"`
	MaturesOn          string  `json:"matures_on,omitempty"`
	MemberGUID         string  `json:"member_guid,omitempty"`
	MinimumBalance     float32 `json:"minimum_balance,omitempty"`
	MinimumPayment     float32 `json:"minimum_payment,omitempty"`
	Name               string  `json:"name,omitempty"`
	OriginalBalance    float32 `json:"original_balance,omitempty"`
	PaymentDueAt       string  `json:"payment_due_at,omitempty"`
	PayoffBalance      float32 `json:"payoff_balance,omitempty"`
	StartedOn          string  `json:"started_on,omitempty"`
	Subtype            string  `json:"subtype,omitempty"`
	TotalAccountValue  float32 `json:"total_account_value,omitempty"`
	Type_              string  `json:"type,omitempty"`
	UpdatedAt          string  `json:"updated_at,omitempty"`
	UserGUID           string  `json:"user_guid,omitempty"`
}
