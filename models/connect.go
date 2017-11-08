package models

type Connect struct {
	URL                  string  `json:"connect_widget_url,omitempty"`
	Guid                 string  `json:"guid,omitempty"`
}

type ConnectResponse struct {
	Connect *Connect `json:"user,omitempty"`
}
