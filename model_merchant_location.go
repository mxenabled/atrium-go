/*
 * MX API
 *
 * The MX Atrium API supports over 48,000 data connections to thousands of financial institutions. It provides secure access to your users' accounts and transactions with industry-leading cleansing, categorization, and classification.  Atrium is designed according to resource-oriented REST architecture and responds with JSON bodies and HTTP response codes.  Use Atrium's development environment, vestibule.mx.com, to quickly get up and running. The development environment limits are 100 users, 25 members per user, and access to the top 15 institutions. Contact MX to purchase production access.
 *
 * API version: 0.1
 */

package atrium

type MerchantLocation struct {
	City          string  `json:"city,omitempty"`
	GUID          string  `json:"guid,omitempty"`
	Latitude      float32 `json:"latitude,omitempty"`
	Longitude     float32 `json:"longitude,omitempty"`
	MerchantGUID  string  `json:"merchant_guid,omitempty"`
	PhoneNumber   string  `json:"phone_number,omitempty"`
	PostalCode    string  `json:"postal_code,omitempty"`
	State         string  `json:"state,omitempty"`
	StoreNumber   string  `json:"store_number,omitempty"`
	StreetAddress string  `json:"street_address,omitempty"`
}
