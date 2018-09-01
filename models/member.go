package models

type Option struct {
	Label     string `json:"label,omitempty"`
	ImageData string `json:"image_data,omitempty"`
	Value     string `json:"value,omitempty"`
}

type Challenge struct {
	FieldName string `json:"field_name,omitempty"`
	Guid      string `json:"guid,omitempty"`
	ImageData string `json:"image_data,omitempty"`
	Label     string `json:"label,omitempty"`
	//TODO: What is this? - Optional string `json:"optional,omitempty"`
	Options []*Option `json:"options,omitempty"`
	Type    string    `json:"type,omitempty"`

	Value string `json:"value,omitempty"`
}

type ChallengesResponse struct {
	Challenges []*Challenge `json:"challenges"`
}

type Member struct {
	AggregatedAt             string       `json:"aggregated_at,omitempty"`
	Challenges               []*Challenge `json:"challenges,omitempty"`
	Guid                     string       `json:"guid,omitempty"`
	HasProcessedAccounts     string       `json:"has_processed_accounts,omitempty"`
	HasProcessedTransactions string       `json:"has_processed_transactions,omitempty"`
	Identifier               string       `json:"identifier,omitempty"`
	InstitutionCode          string       `json:"institution_code,omitempty"`
	IsBeingAggregated        bool         `json:"is_being_aggregated,omitempty"`
	Metadata                 string       `json:"metadata,omitempty"`
	Name                     string       `json:"name,omitempty"`
	Status                   string       `json:"status,omitempty"`
	SuccessfullyAggregatedAt string       `json:"successfully_aggregated_at,omitempty"`
	UserGuid                 string       `json:"user_guid,omitempty"`
	ConnectionStatus         string       `json:"connection_status,omitempty"`
}

type MemberCreate struct {
	Credentials     []*Credential `json:"credentials"`
	Identifier      string        `json:"identifier,omitempty"`
	InstitutionCode string        `json:"institution_code,omitempty"`
	Metadata        string        `json:"metadata,omitempty"`
}

type MemberCreateRequest struct {
	Member *MemberCreate `json:"member"`
}

func NewMemberCreateRequest(member *Member, credentials []*Credential) *MemberCreateRequest {
	return &MemberCreateRequest{
		Member: &MemberCreate{
			Credentials:     credentials,
			Identifier:      member.Identifier,
			InstitutionCode: member.InstitutionCode,
			Metadata:        member.Metadata,
		},
	}
}

type MemberUpdate struct {
	Credentials []*Credential `json:"credentials,omitempty"`
	Identifier  string        `json:"identifier,omitempty"`
	Metadata    string        `json:"metadata,omitempty"`
}

type MemberUpdateRequest struct {
	Member *MemberUpdate `json:"member,omitempty"`
}

func NewMemberUpdateRequest(member *Member, credentials []*Credential) *MemberUpdateRequest {
	return &MemberUpdateRequest{
		Member: &MemberUpdate{
			Credentials: credentials,
			Identifier:  member.Identifier,
			Metadata:    member.Metadata,
		},
	}
}

type MemberResume struct {
	Challenges []*Challenge `json:"challenges"`
}

type MemberResumeRequest struct {
	Member *MemberResume `json:"member"`
}

func NewMemberResumeRequest(challenges []*Challenge) *MemberResumeRequest {
	return &MemberResumeRequest{
		&MemberResume{
			Challenges: challenges,
		},
	}
}

type MemberResponse struct {
	Member *Member `json:"member,omitempty"`
}

type MembersResponse struct {
	Members []*Member `json:"members,omitempty"`
}
