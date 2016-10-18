package models

type Credential struct {
	FieldName string `json:"field_name,omitempty"`
	Guid      string `json:"guid,omitempty"`
	Label     string `json:"label,omitempty"`
	Type      string `json:"type,omitempty"`

	// Only for creating
	Value      string `json:"value,omitempty"`
}

type CredentialsResponse struct {
	Credentials []*Credential `json:"credentials"`
}