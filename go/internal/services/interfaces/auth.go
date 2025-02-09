package services

import (
	"context"
	"emregami/internal/services/dto"
)

type AuthService interface {
	Register(ctx context.Context, request *dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(ctx context.Context, request *dto.LoginRequest) (*dto.LoginResponse, error)
}
