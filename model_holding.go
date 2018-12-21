/*
 * MX API
 *
 * The MX Atrium API supports over 48,000 data connections to thousands of financial institutions. It provides secure access to your users' accounts and transactions with industry-leading cleansing, categorization, and classification.  Atrium is designed according to resource-oriented REST architecture and responds with JSON bodies and HTTP response codes.  Use Atrium's development environment, vestibule.mx.com, to quickly get up and running. The development environment limits are 100 users, 25 members per user, and access to the top 15 institutions. Contact MX to purchase production access.
 *
 * API version: 0.1
 */

package atrium

type Holding struct {
	AccountGUID   string  `json:"account_guid,omitempty"`
	CostBasis     float32 `json:"cost_basis,omitempty"`
	CreatedAt     string  `json:"created_at,omitempty"`
	CurrencyCode  string  `json:"currency_code,omitempty"`
	Cusip         string  `json:"cusip,omitempty"`
	DailyChange   float32 `json:"daily_change,omitempty"`
	Description   string  `json:"description,omitempty"`
	GUID          string  `json:"guid,omitempty"`
	HoldingType   string  `json:"holding_type,omitempty"`
	MarketValue   float32 `json:"market_value,omitempty"`
	MemberGUID    string  `json:"member_guid,omitempty"`
	PurchasePrice float32 `json:"purchase_price,omitempty"`
	Shares        float32 `json:"shares,omitempty"`
	Symbol        string  `json:"symbol,omitempty"`
	UpdatedAt     string  `json:"updated_at,omitempty"`
	UserGUID      string  `json:"user_guid,omitempty"`
}
