/*
 * MX API
 *
 * The MX Atrium API supports over 48,000 data connections to thousands of financial institutions. It provides secure access to your users' accounts and transactions with industry-leading cleansing, categorization, and classification.  Atrium is designed according to resource-oriented REST architecture and responds with JSON bodies and HTTP response codes.  Use Atrium's development environment, vestibule.mx.com, to quickly get up and running. The development environment limits are 100 users, 25 members per user, and access to the top 15 institutions. Contact MX to purchase production access.
 *
 * API version: 0.1
 */

package atrium

type Statement struct {
	// The unique identifier for the `account` associated with the `statement`. Defined by MX.
	AccountGUID string `json:"account_guid,omitempty"`
	// SHA256 digest of the pdf payload
	ContentHash string `json:"content_hash,omitempty"`
	// The date and time the `statement` was created.
	CreatedAt string `json:"created_at,omitempty"`
	// The unique identifier for the `statement`. Defined by MX.
	GUID string `json:"guid,omitempty"`
	// The unique identifier for the `member` associated with the `statement`.  Defined by MX.
	MemberGUID string `json:"member_guid,omitempty"`
	// A URI for accessing the byte payload of the `statement`.
	Uri string `json:"uri,omitempty"`
	// The unique identifier for the `user` associated with the `statement`.  Defined by MX.
	UserGUID string `json:"user_guid,omitempty"`
	// The date and time at which the `statement` was last updated.
	UpdatedAt string `json:"updated_at,omitempty"`
}
