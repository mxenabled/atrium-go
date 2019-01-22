/*
 * MX API
 *
 * The MX Atrium API supports over 48,000 data connections to thousands of financial institutions. It provides secure access to your users' accounts and transactions with industry-leading cleansing, categorization, and classification.  Atrium is designed according to resource-oriented REST architecture and responds with JSON bodies and HTTP response codes.  Use Atrium's development environment, vestibule.mx.com, to quickly get up and running. The development environment limits are 100 users, 25 members per user, and access to the top 15 institutions. Contact MX to purchase production access.
 *
 * API version: 0.1
 */

package atrium

type Member struct {
	AggregatedAt             string `json:"aggregated_at,omitempty"`
	ConnectionStatus         string `json:"connection_status,omitempty"`
	GUID                     string `json:"guid,omitempty"`
	Identifier               string `json:"identifier,omitempty"`
	InstitutionCode          string `json:"institution_code,omitempty"`
	IsBeingAggregated        bool   `json:"is_being_aggregated,omitempty"`
	Metadata                 string `json:"metadata,omitempty"`
	Name                     string `json:"name,omitempty"`
	Status                   string `json:"status,omitempty"`
	SuccessfullyAggregatedAt string `json:"successfully_aggregated_at,omitempty"`
	UserGUID                 string `json:"user_guid,omitempty"`
}
