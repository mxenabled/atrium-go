package models

// TODO: Verify this junk
type Account struct {
	Apr float64
	Apy float64
	AvailableBalance float64
	AvailableCredit float64
	Balance float64
	CreatedAt float64
	CreditLimit float64
	DayPaymentIsDue string
	Guid string
	HoldingsValue float64
	InstitutionCode string
	InterestRate float64
	IsClosed bool
	LastPayment float64
	LastPaymentAt string
	MaturesOn string
	MemberGuid string
	MinimumBalance float64
	MinimumPayment float64
	Name string
	OriginalBalance float64
	PaymentDueAt string
	PayoffBalance float64
	StartedOn string
	Subtype string
	TotalAccountValue float64
	Type string
	UpdatedAt string
	UserGuid string
}

type AccountResponse struct {
	Account *Account `json"account"`
}

type AccountsResponse struct {
	Accounts []*Account `json"accounts"`
}
