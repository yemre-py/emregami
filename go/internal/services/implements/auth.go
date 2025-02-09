package implements

import (
	"context"
	"emregami/internal/domain"
	"emregami/internal/repository/repos"
	"emregami/internal/services/dto"
	services "emregami/internal/services/interfaces"
	"emregami/pkg/utils"
	"emregami/pkg/validations"
	"errors"

	"github.com/go-playground/validator"
)

type service struct {
	authRepository repos.AuthRepository
	validate       *validator.Validate
}

func NewAuthService(authRepository repos.AuthRepository) services.AuthService {
	return &service{
		authRepository: authRepository,
		validate:       validations.GetValidator(),
	}
}

func (s *service) Register(ctx context.Context, request *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	// input validation
	if err := s.validate.Struct(request); err != nil {
		return nil, errors.New("invalid request")
	}

	// check if user already exists
	user, err := s.authRepository.GetByUsername(ctx, request.Username)
	if err != nil {
		return nil, errors.New("failed to get user by username")
	}

	if user != nil {
		return nil, errors.New("user already exists")
	}

	email, err := s.authRepository.GetByEmail(ctx, request.Email)
	if err != nil {
		return nil, errors.New("failed to get user by email")
	}

	if email != nil {
		return nil, errors.New("email already exists")
	}

	// generate uuid
	uuid, err := utils.GenerateUUID()
	if err != nil {
		return nil, errors.New("failed to generate uuid")
	}

	// hash password
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// create auth
	auth := &domain.Auth{
		ID:       uuid,
		Username: request.Username,
		Email:    request.Email,
		Password: hashedPassword,
	}

	// save user
	if err := s.authRepository.Save(ctx, auth); err != nil {
		return nil, errors.New("failed to save user")
	}

	// generate tokens
	tokens, err := utils.GenerateTokens(auth.ID)
	if err != nil {
		return nil, errors.New("failed to generate tokens")
	}

	return &dto.RegisterResponse{
		Message: "User registered successfully",
		Tokens:  tokens,
	}, nil
}

func (s *service) Login(request *dto.LoginRequest) (*dto.LoginResponse, error) {
	return nil, nil
}
