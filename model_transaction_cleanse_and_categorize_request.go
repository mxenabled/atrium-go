/*
 * MX API
 *
 * The MX Atrium API supports over 48,000 data connections to thousands of financial institutions. It provides secure access to your users' accounts and transactions with industry-leading cleansing, categorization, and classification.  Atrium is designed according to resource-oriented REST architecture and responds with JSON bodies and HTTP response codes.  Use Atrium's development environment, vestibule.mx.com, to quickly get up and running. The development environment limits are 100 users, 25 members per user, and access to the top 15 institutions. Contact MX to purchase production access.
 *
 * API version: 0.1
 */

package atrium

type TransactionCleanseAndCategorizeRequest struct {
	Amount               float32 `json:"amount,omitempty"`
	Description          string  `json:"description,omitempty"`
	Identifier           string  `json:"identifier,omitempty"`
	MerchantCategoryCode float32 `json:"merchant_category_code,omitempty"`
	Type_                string  `json:"type,omitempty"`
}
