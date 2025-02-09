package dto

type LoginRequest struct {
	Username string `json:"username"` // we gonna get username or email
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Tokens
}
