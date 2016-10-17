package models

type Option struct {
	Label string `json"label,omitempty"`
	Value string `json"value,omitempty"`
}

type Challenge struct {
	FieldName string `json"field_name,omitempty"`
	Guid string `json"guid,omitempty"`
	Label string `json"label,omitempty"`
	//TODO: What is this? - Optional string `json"optional,omitempty"`
	Options []*Option `json"options,omitempty"`
	Type string
}

type Member struct {
	AggregatedAt string `json"aggregated_at,omitempty"`
	Guid string `json"guid,omitempty"`
	Identifier string `json"identifier,omitempty"`
	InstitutionCode string `json"institution_code,omitempty"`
	Metadata string `json"metadata,omitempty"`
	Name string `json"name,omitempty"`
	Status string `json"status,omitempty"`
	SuccessfullyAggregatedAt string `json"successfully_aggregated_at,omitempty"`
	UserGuid string `json"user_guid,omitempty"`
}

type MemberCreateRequest struct {
	Credentials []*Credential
	Identifier string `json"identifier,omitempty"`
	InstitutionCode string `json"institution_code,omitempty"`
	Metadata string `json"metadata,omitempty"`
}

type MemberUpdateRequest struct {
	Credentials []*Credential
	Identifier string `json"identifier,omitempty"`
	Metadata string `json"metadata,omitempty"`
}

type MemberResumeRequest struct {
	Challenges []*Challenge
}

type MemberResponse struct {
	Member *Member `json"member,omitempty"`
}

type MembersResponse struct {
	Members []*Member `json"members,omitempty"`
}
