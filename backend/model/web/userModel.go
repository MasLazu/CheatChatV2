package web

type RegisterUserRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}
