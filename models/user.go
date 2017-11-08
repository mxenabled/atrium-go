package models

type UserAttributes struct {
	Metadata   string `json:"metadata,omitempty"`
	IsDisabled bool   `json:"is_disabled,omitempty"`
	Identifier string `json:"identifier,omitempty"`
}

type User struct {
	Guid       string `json:"guid,omitempty"`
	Metadata   string `json:"metadata,omitempty"`
	IsDisabled bool   `json:"is_disabled,omitempty"`
	Identifier string `json:"identifier,omitempty"`
}

type UserResponse struct {
	User *User `json:"user"`
}

type UserRequest struct {
	UserAttributes *UserAttributes `json:"user"`
}

func NewUserRequest(user *User) *UserRequest {
	return &UserRequest{
		UserAttributes: &UserAttributes{
			Metadata:   user.Metadata,
			Identifier: user.Identifier,
			IsDisabled: user.IsDisabled,
		},
	}
}

type UsersResponse struct {
	Users []*User `json"users"`
}
