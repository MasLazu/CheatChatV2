package model

type LoginUserRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}

type CuerrentResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}
