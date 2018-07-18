package models

type Connect struct {
	URL  string `json:"connect_widget_url,omitempty"`
	Guid string `json:"guid,omitempty"`
}

type ConnectResponse struct {
	Connect *Connect `json:"user,omitempty"`
}

type ConnectParams struct {
	IsMobileWebView        bool    `json:"is_mobile_webview"`
	UpdateCredentials      bool    `json:"update_credentials"`
	CurrentInstitutionCode *string `json:"current_institution_code"`
	CurrentMemberGuid      *string `json:"current_member_guid"`
}
