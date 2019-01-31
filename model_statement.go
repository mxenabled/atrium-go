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
	// The date and time the `statement` was created.
	CreatedAt string `json:"created_at,omitempty"`
	// An SHA-256 hash value of the statement's byte payload, used as a unique identifier.
	ContentHash string `json:"content_hash,omitempty"`
	// The date and time the `statement` was deleted. Statements are automatically deleted when an `account` is deleted.
	DeletedAt string `json:"deleted_at,omitempty"`
	// The unique identifier for the `statement`. Defined by MX.
	GUID string `json:"guid,omitempty"`
	// This indicates whether the `statement` has been deleted. Statements are automatically deleted when an `account` is deleted.
	IsDeleted bool `json:"is_deleted,omitempty"`
	// The date and time at which the `statement` was last updated.
	UpdatedAt string `json:"updated_at,omitempty"`
	// A URI for accessing the byte payload of the `statement`.
	Uri string `json:"uri,omitempty"`
	// The unique identifier for the `user` associated with the `statement`.  Defined by MX.
	UserGUID string `json:"user_guid,omitempty"`
}
