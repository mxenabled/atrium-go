package models

type Institution struct {
	Code string `json"code,omitempty"`
	Name string `json"name,omitempty"`
	URL string `json"url,omitempty"`
}

type InstitutionsResponse struct {
	Institutions []*Institution `json"institutions"`
}
