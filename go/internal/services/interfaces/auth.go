package interfaces

import "emregami/internal/services/dto"

type AuthService interface {
	Register(request *dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(request *dto.LoginRequest) (*dto.LoginResponse, error)
}
