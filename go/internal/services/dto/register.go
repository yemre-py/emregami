package dto

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"` // we gonna get username or email
}

type RegisterResponse struct {
	Message string `json:"message"`
	Tokens
}
