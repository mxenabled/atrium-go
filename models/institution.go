package models

type Institution struct {
	Code          string `json"code,omitempty"`
	MediumLogoURL string `json"code,omitempty"`
	Name          string `json"name,omitempty"`
	URL           string `json"url,omitempty"`
	SmallLogoURL  string `json"url,omitempty"`
}

type InstitutionsResponse struct {
	Institutions []*Institution `json"institutions"`
}

type InstitutionResponse struct {
	Institution *Institution `json:"institution"`
}
