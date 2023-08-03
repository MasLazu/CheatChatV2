package domain

type Contact struct {
	Name           string `json:"name,omitempty"`
	UserEmail      string `json:"user_email,omitempty"`
	SavedUserEmail string `json:"saved_user_email,omitempty"`
}
