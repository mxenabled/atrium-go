/*
 * MX API
 *
 * The MX Atrium API supports over 48,000 data connections to thousands of financial institutions. It provides secure access to your users' accounts and transactions with industry-leading cleansing, categorization, and classification.  Atrium is designed according to resource-oriented REST architecture and responds with JSON bodies and HTTP response codes.  Use Atrium's development environment, vestibule.mx.com, to quickly get up and running. The development environment limits are 100 users, 25 members per user, and access to the top 15 institutions. Contact MX to purchase production access.
 *
 * API version: 0.1
 */

package atrium

type ConnectWidgetRequestBody struct {
	IsMobileWebview           bool    `json:"is_mobile_webview,omitempty"`
	ColorScheme               string  `json:"color_scheme,omitempty"`
	CurrentInstitutionCode    string  `json:"current_institution_code,omitempty"`
	CurrentMemberGUID         string  `json:"current_member_guid,omitempty"`
	DisableInstitutionSearch  bool    `json:"disable_institution_search,omitempty"`
	IncludeTransactions       bool    `json:"include_transactions,omitempty"`
	Mode                      string  `json:"mode,omitempty"`
	UiMessageVersion          float32 `json:"ui_message_version,omitempty"`
	UiMessageWebviewURLScheme string  `json:"ui_message_webview_url_scheme,omitempty"`
	UpdateCredentials         bool    `json:"update_credentials,omitempty"`
	WaitForFullAggregation    bool    `json:"wait_for_full_aggregation,omitempty"`
}
