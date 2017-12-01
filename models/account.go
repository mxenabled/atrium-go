package models

// TODO: Verify this junk
type Account struct {
	Apr               float64 `json:"apr,omitempty"`
	Apy               float64 `json:"apy,omitempty"`
	AvailableBalance  float64 `json:"available_balance,omitempty"`
	AvailableCredit   float64 `json:"available_credit,omitempty"`
	Balance           float64 `json:"balance,omitempty"`
	CreatedAt         float64 `json:"created_at,omitempty"`
	CreditLimit       float64 `json:"credit_limit,omitempty"`
	DayPaymentIsDue   string `json:"day_payment_is_due,omitempty"`
	Guid              string `json:"guid,omitempty"`
	InstitutionCode   string `json:"institution_code,omitempty"`
	InterestRate      float64 `json:"interest_rate,omitempty"`
	IsClosed          bool `json:"is_closed,omitempty"`
	LastPayment       float64 `json:"last_payment,omitempty"`
	LastPaymentAt     string `json:"last_payment_at,omitempty"`
	MaturesOn         string `json:"matures_on,omitempty"`
	MemberGuid        string `json:"member_guid,omitempty"`
	MinimumBalance    float64 `json:"minimum_balance,omitempty"`
	MinimumPayment    float64 `json:"minimum_payment,omitempty"`
	Name              string `json:"name,omitempty"`
	OriginalBalance   float64 `json:"original_balance,omitempty"`
	PaymentDueAt      string `json:"payment_due_at,omitempty"`
	PayoffBalance     float64 `json:"payoff_balance,omitempty"`
	StartedOn         string `json:"started_on,omitempty"`
	Subtype           string `json:"subtype,omitempty"`
	TotalAccountValue float64 `json:"total_account_value,omitempty"`
	Type              string `json:"type,omitempty"`
	UpdatedAt         string `json:"updated_at,omitempty"`
	UserGuid          string `json:"user_guid,omitempty"`
}

type AccountResponse struct {
	Account *Account `json:"account"`
}

type AccountsResponse struct {
	Accounts []*Account `json:"accounts"`
}
