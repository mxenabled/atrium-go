package models

type AccountOwner struct {
	Address     string `json:"address,omitempty"`
	AccountGuid string `json:"account_guid,omitempty"`
	City        string `json:"city,omitempty"`
	Country     string `json:"country,omitempty"`
	Email       string `json:"email,omitempty"`
	Guid        string `json:"guid,omitempty"`
	MemberGuid  string `json:"member_guid,omitempty"`
	OwnerName   string `json:"owner_name,omitempty"`
	Phone       string `json:"phone,omitempty"`
	PostalCode  string `json:"postal_code,omitempty"`
	State       string `json:"state,omitempty"`
	UserGuid    string `json:"user_guid,omitempty"`
}

type AccountOwnersResponse struct {
	AccountOwners []*AccountOwner `json:"account_owners"`
}
