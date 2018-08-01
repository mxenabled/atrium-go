package models

type AccountNumber struct {
	AccountGuid   string `json:"account_guid,omitempty"`
	AccountNumber string `json:"account_number,omitempty"`
	Guid          string `json:"guid,omitempty"`
	MemberGuid    string `json:"member_guid,omitempty"`
	RoutingNumber string `json:"routing_number,omitempty"`
	UserGuid      string `json:"user_guid,omitempty"`
}

type AccountNumbersResponse struct {
	AccountNumbers []*AccountNumber `json:"account_numbers"`
}
