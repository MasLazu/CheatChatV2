package domain

type Contact struct {
	Name           string `json:"name,omitempty"`
	UserEmail      string `json:"user_email,omitempty"`
	SavedUserEmail string `json:"email,omitempty"`
}
