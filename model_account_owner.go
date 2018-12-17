/*
 * MX API
 *
 * The MX Atrium API supports over 48,000 data connections to thousands of financial institutions. It provides secure access to your users' accounts and transactions with industry-leading cleansing, categorization, and classification.  Atrium is designed according to resource-oriented REST architecture and responds with JSON bodies and HTTP response codes.  Use Atrium's development environment, vestibule.mx.com, to quickly get up and running. The development environment limits are 100 users, 25 members per user, and access to the top 15 institutions. Contact MX to purchase production access.
 *
 * API version: 0.1
 */

package atrium

type AccountOwner struct {
	AccountGUID string `json:"account_guid,omitempty"`
	Address     string `json:"address,omitempty"`
	City        string `json:"city,omitempty"`
	Country     string `json:"country,omitempty"`
	Email       string `json:"email,omitempty"`
	GUID        string `json:"guid,omitempty"`
	MemberGUID  string `json:"member_guid,omitempty"`
	OwnerName   string `json:"owner_name,omitempty"`
	PostalCode  string `json:"postal_code,omitempty"`
	State       string `json:"state,omitempty"`
	UserGUID    string `json:"user_guid,omitempty"`
}
